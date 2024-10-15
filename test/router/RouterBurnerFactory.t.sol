// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Test} from "forge-std/Test.sol";

import {VaultFactory} from "@symbioticfi/core/src/contracts/VaultFactory.sol";
import {DelegatorFactory} from "@symbioticfi/core/src/contracts/DelegatorFactory.sol";
import {SlasherFactory} from "@symbioticfi/core/src/contracts/SlasherFactory.sol";
import {NetworkRegistry} from "@symbioticfi/core/src/contracts/NetworkRegistry.sol";
import {OperatorRegistry} from "@symbioticfi/core/src/contracts/OperatorRegistry.sol";
import {MetadataService} from "@symbioticfi/core/src/contracts/service/MetadataService.sol";
import {NetworkMiddlewareService} from "@symbioticfi/core/src/contracts/service/NetworkMiddlewareService.sol";
import {OptInService} from "@symbioticfi/core/src/contracts/service/OptInService.sol";

import {Vault} from "@symbioticfi/core/src/contracts/vault/Vault.sol";
import {NetworkRestakeDelegator} from "@symbioticfi/core/src/contracts/delegator/NetworkRestakeDelegator.sol";
import {FullRestakeDelegator} from "@symbioticfi/core/src/contracts/delegator/FullRestakeDelegator.sol";
import {OperatorSpecificDelegator} from "@symbioticfi/core/src/contracts/delegator/OperatorSpecificDelegator.sol";
import {Slasher} from "@symbioticfi/core/src/contracts/slasher/Slasher.sol";
import {VetoSlasher} from "@symbioticfi/core/src/contracts/slasher/VetoSlasher.sol";

import {Token} from "@symbioticfi/core/test/mocks/Token.sol";
import {VaultConfigurator, IVaultConfigurator} from "@symbioticfi/core/src/contracts/VaultConfigurator.sol";
import {IVault} from "@symbioticfi/core/src/interfaces/vault/IVault.sol";
import {INetworkRestakeDelegator} from "@symbioticfi/core/src/interfaces/delegator/INetworkRestakeDelegator.sol";
import {IBaseDelegator} from "@symbioticfi/core/src/interfaces/delegator/IBaseDelegator.sol";
import {IBaseSlasher} from "@symbioticfi/core/src/interfaces/slasher/IBaseSlasher.sol";
import {ISlasher} from "@symbioticfi/core/src/interfaces/slasher/ISlasher.sol";

import {RouterBurnerFactory} from "../../src/contracts/router/RouterBurnerFactory.sol";
import {RouterBurner} from "../../src/contracts/router/RouterBurner.sol";
import {IRouterBurner} from "../../src/interfaces/router/IRouterBurner.sol";

contract RouterBurnerFactoryTest is Test {
    address owner;
    address alice;
    uint256 alicePrivateKey;
    address bob;
    uint256 bobPrivateKey;

    VaultFactory vaultFactory;
    DelegatorFactory delegatorFactory;
    SlasherFactory slasherFactory;
    NetworkRegistry networkRegistry;
    OperatorRegistry operatorRegistry;
    MetadataService operatorMetadataService;
    MetadataService networkMetadataService;
    NetworkMiddlewareService networkMiddlewareService;
    OptInService networkVaultOptInService;
    OptInService operatorVaultOptInService;
    OptInService operatorNetworkOptInService;

    Token collateral;
    VaultConfigurator vaultConfigurator;

    Vault vault;
    FullRestakeDelegator delegator;
    Slasher slasher;

    RouterBurnerFactory routerBurnerFactory;
    RouterBurner routerBurner;

    function setUp() public {
        owner = address(this);
        (alice, alicePrivateKey) = makeAddrAndKey("alice");
        (bob, bobPrivateKey) = makeAddrAndKey("bob");

        vaultFactory = new VaultFactory(owner);
        delegatorFactory = new DelegatorFactory(owner);
        slasherFactory = new SlasherFactory(owner);
        networkRegistry = new NetworkRegistry();
        operatorRegistry = new OperatorRegistry();
        operatorMetadataService = new MetadataService(address(operatorRegistry));
        networkMetadataService = new MetadataService(address(networkRegistry));
        networkMiddlewareService = new NetworkMiddlewareService(address(networkRegistry));
        operatorVaultOptInService =
            new OptInService(address(operatorRegistry), address(vaultFactory), "OperatorVaultOptInService");
        operatorNetworkOptInService =
            new OptInService(address(operatorRegistry), address(networkRegistry), "OperatorNetworkOptInService");

        address vaultImpl =
            address(new Vault(address(delegatorFactory), address(slasherFactory), address(vaultFactory)));
        vaultFactory.whitelist(vaultImpl);

        address networkRestakeDelegatorImpl = address(
            new NetworkRestakeDelegator(
                address(networkRegistry),
                address(vaultFactory),
                address(operatorVaultOptInService),
                address(operatorNetworkOptInService),
                address(delegatorFactory),
                delegatorFactory.totalTypes()
            )
        );
        delegatorFactory.whitelist(networkRestakeDelegatorImpl);

        address fullRestakeDelegatorImpl = address(
            new FullRestakeDelegator(
                address(networkRegistry),
                address(vaultFactory),
                address(operatorVaultOptInService),
                address(operatorNetworkOptInService),
                address(delegatorFactory),
                delegatorFactory.totalTypes()
            )
        );
        delegatorFactory.whitelist(fullRestakeDelegatorImpl);

        address operatorSpecificDelegatorImpl = address(
            new OperatorSpecificDelegator(
                address(operatorRegistry),
                address(networkRegistry),
                address(vaultFactory),
                address(operatorVaultOptInService),
                address(operatorNetworkOptInService),
                address(delegatorFactory),
                delegatorFactory.totalTypes()
            )
        );
        delegatorFactory.whitelist(operatorSpecificDelegatorImpl);

        address slasherImpl = address(
            new Slasher(
                address(vaultFactory),
                address(networkMiddlewareService),
                address(slasherFactory),
                slasherFactory.totalTypes()
            )
        );
        slasherFactory.whitelist(slasherImpl);

        address vetoSlasherImpl = address(
            new VetoSlasher(
                address(vaultFactory),
                address(networkMiddlewareService),
                address(networkRegistry),
                address(slasherFactory),
                slasherFactory.totalTypes()
            )
        );
        slasherFactory.whitelist(vetoSlasherImpl);

        collateral = new Token("Token");

        vaultConfigurator =
            new VaultConfigurator(address(vaultFactory), address(delegatorFactory), address(slasherFactory));
    }

    function test_Create(uint256 receiverSetEpochsDelay, address globalReceiver) public {
        receiverSetEpochsDelay = bound(receiverSetEpochsDelay, 3, type(uint256).max);

        address routerBurnerImplementation = address(new RouterBurner(address(vaultFactory)));
        routerBurnerFactory = new RouterBurnerFactory(routerBurnerImplementation);

        uint160 N1 = 10;
        IRouterBurner.NetworkReceiver[] memory networkReceivers = new IRouterBurner.NetworkReceiver[](N1);
        uint160 N2 = 20;
        IRouterBurner.OperatorNetworkReceiver[] memory operatorNetworkReceivers =
            new IRouterBurner.OperatorNetworkReceiver[](N2);

        for (uint160 i; i < N1; ++i) {
            networkReceivers[i] =
                IRouterBurner.NetworkReceiver({network: address(i * 2 + 3), receiver: address(i * 2 + 1)});
        }

        for (uint160 i; i < N2; ++i) {
            operatorNetworkReceivers[i] = IRouterBurner.OperatorNetworkReceiver({
                network: address(i * 2 + 3),
                operator: address(i * 2 + 1),
                receiver: address(i * 2 + 2)
            });
        }

        IRouterBurner.InitParams memory initParams = IRouterBurner.InitParams({
            owner: owner,
            receiverSetEpochsDelay: receiverSetEpochsDelay,
            globalReceiver: globalReceiver,
            networkReceivers: networkReceivers,
            operatorNetworkReceivers: operatorNetworkReceivers
        });

        address routerBurnerAddress = routerBurnerFactory.create(initParams);
        routerBurner = RouterBurner(routerBurnerAddress);

        assertTrue(routerBurnerFactory.isEntity(routerBurnerAddress));

        assertEq(routerBurner.VAULT_FACTORY(), address(vaultFactory));
        assertEq(routerBurner.owner(), owner);
        assertEq(routerBurner.vault(), address(0));
        assertEq(routerBurner.collateral(), address(0));
        assertEq(routerBurner.receiverSetEpochsDelay(), receiverSetEpochsDelay);
        assertEq(routerBurner.lastBalance(), 0);
        assertEq(routerBurner.globalReceiver(), globalReceiver);
        (address pendingAddress, uint48 pendingTimestamp) = routerBurner.pendingGlobalReceiver();
        assertEq(pendingAddress, address(0));
        assertEq(pendingTimestamp, 0);
        for (uint160 i; i < N1 + 1; ++i) {
            assertEq(routerBurner.networkReceiver(address(i * 2 + 3)), i < N1 ? address(i * 2 + 1) : address(0));
        }
        (pendingAddress, pendingTimestamp) = routerBurner.pendingNetworkReceiver(address(0));
        assertEq(pendingAddress, address(0));
        assertEq(pendingTimestamp, 0);
        for (uint160 i; i < N2 + 1; ++i) {
            assertEq(
                routerBurner.operatorNetworkReceiver(address(i * 2 + 3), address(i * 2 + 1)),
                i < N2 ? address(i * 2 + 2) : address(0)
            );
        }
        (pendingAddress, pendingTimestamp) = routerBurner.pendingOperatorNetworkReceiver(address(0), address(0));
        assertEq(pendingAddress, address(0));
        assertEq(pendingTimestamp, 0);
        assertEq(routerBurner.balanceOf(address(0)), 0);
        assertEq(routerBurner.isInitialized(), false);

        (vault, delegator, slasher) = _getVaultWithDelegatorWithSlasher(address(routerBurner));

        vm.startPrank(owner);
        routerBurner.setVault(address(vault));
        vm.stopPrank();

        assertEq(routerBurner.VAULT_FACTORY(), address(vaultFactory));
        assertEq(routerBurner.owner(), owner);
        assertEq(routerBurner.vault(), address(vault));
        assertEq(routerBurner.collateral(), vault.collateral());
        assertEq(routerBurner.receiverSetEpochsDelay(), receiverSetEpochsDelay);
        assertEq(routerBurner.lastBalance(), 0);
        assertEq(routerBurner.globalReceiver(), globalReceiver);
        (pendingAddress, pendingTimestamp) = routerBurner.pendingGlobalReceiver();
        assertEq(pendingAddress, address(0));
        assertEq(pendingTimestamp, 0);
        for (uint160 i; i < N1 + 1; ++i) {
            assertEq(routerBurner.networkReceiver(address(i * 2 + 3)), i < N1 ? address(i * 2 + 1) : address(0));
        }
        (pendingAddress, pendingTimestamp) = routerBurner.pendingNetworkReceiver(address(0));
        assertEq(pendingAddress, address(0));
        assertEq(pendingTimestamp, 0);
        for (uint160 i; i < N2 + 1; ++i) {
            assertEq(
                routerBurner.operatorNetworkReceiver(address(i * 2 + 3), address(i * 2 + 1)),
                i < N2 ? address(i * 2 + 2) : address(0)
            );
        }
        (pendingAddress, pendingTimestamp) = routerBurner.pendingOperatorNetworkReceiver(address(0), address(0));
        assertEq(pendingAddress, address(0));
        assertEq(pendingTimestamp, 0);
        assertEq(routerBurner.balanceOf(address(0)), 0);
        assertEq(routerBurner.isInitialized(), true);
    }

    function test_CreateRevertInvalidReceiverSetEpochsDelay(
        uint256 receiverSetEpochsDelay,
        address globalReceiver
    ) public {
        receiverSetEpochsDelay = bound(receiverSetEpochsDelay, 0, 2);

        address routerBurnerImplementation = address(new RouterBurner(address(vaultFactory)));
        routerBurnerFactory = new RouterBurnerFactory(routerBurnerImplementation);

        uint160 N1 = 10;
        IRouterBurner.NetworkReceiver[] memory networkReceivers = new IRouterBurner.NetworkReceiver[](N1);
        uint160 N2 = 20;
        IRouterBurner.OperatorNetworkReceiver[] memory operatorNetworkReceivers =
            new IRouterBurner.OperatorNetworkReceiver[](N2);

        for (uint160 i; i < N1; ++i) {
            networkReceivers[i] =
                IRouterBurner.NetworkReceiver({network: address(i * 2 + 3), receiver: address(i * 2 + 1)});
        }

        for (uint160 i; i < N2; ++i) {
            operatorNetworkReceivers[i] = IRouterBurner.OperatorNetworkReceiver({
                network: address(i * 2 + 3),
                operator: address(i * 2 + 1),
                receiver: address(i * 2 + 2)
            });
        }

        IRouterBurner.InitParams memory initParams = IRouterBurner.InitParams({
            owner: owner,
            receiverSetEpochsDelay: receiverSetEpochsDelay,
            globalReceiver: globalReceiver,
            networkReceivers: networkReceivers,
            operatorNetworkReceivers: operatorNetworkReceivers
        });

        vm.expectRevert(IRouterBurner.InvalidReceiverSetEpochsDelay.selector);
        address routerBurnerAddress = routerBurnerFactory.create(initParams);
    }

    function test_CreateRevertInvalidReceiver1(uint256 receiverSetEpochsDelay, address globalReceiver) public {
        receiverSetEpochsDelay = bound(receiverSetEpochsDelay, 3, type(uint256).max);

        address routerBurnerImplementation = address(new RouterBurner(address(vaultFactory)));
        routerBurnerFactory = new RouterBurnerFactory(routerBurnerImplementation);

        uint160 N1 = 10;
        IRouterBurner.NetworkReceiver[] memory networkReceivers = new IRouterBurner.NetworkReceiver[](N1);
        uint160 N2 = 20;
        IRouterBurner.OperatorNetworkReceiver[] memory operatorNetworkReceivers =
            new IRouterBurner.OperatorNetworkReceiver[](N2);

        for (uint160 i; i < N1; ++i) {
            networkReceivers[i] = IRouterBurner.NetworkReceiver({network: address(i * 2 + 3), receiver: address(0)});
        }

        for (uint160 i; i < N2; ++i) {
            operatorNetworkReceivers[i] = IRouterBurner.OperatorNetworkReceiver({
                network: address(i * 2 + 3),
                operator: address(i * 2 + 1),
                receiver: address(i * 2 + 2)
            });
        }

        IRouterBurner.InitParams memory initParams = IRouterBurner.InitParams({
            owner: owner,
            receiverSetEpochsDelay: receiverSetEpochsDelay,
            globalReceiver: globalReceiver,
            networkReceivers: networkReceivers,
            operatorNetworkReceivers: operatorNetworkReceivers
        });

        vm.expectRevert(IRouterBurner.InvalidReceiver.selector);
        address routerBurnerAddress = routerBurnerFactory.create(initParams);
    }

    function test_CreateRevertInvalidReceiver2(uint256 receiverSetEpochsDelay, address globalReceiver) public {
        receiverSetEpochsDelay = bound(receiverSetEpochsDelay, 3, type(uint256).max);

        address routerBurnerImplementation = address(new RouterBurner(address(vaultFactory)));
        routerBurnerFactory = new RouterBurnerFactory(routerBurnerImplementation);

        uint160 N1 = 10;
        IRouterBurner.NetworkReceiver[] memory networkReceivers = new IRouterBurner.NetworkReceiver[](N1);
        uint160 N2 = 20;
        IRouterBurner.OperatorNetworkReceiver[] memory operatorNetworkReceivers =
            new IRouterBurner.OperatorNetworkReceiver[](N2);

        for (uint160 i; i < N1; ++i) {
            networkReceivers[i] =
                IRouterBurner.NetworkReceiver({network: address(i * 2 + 3), receiver: address(i * 2 + 1)});
        }

        for (uint160 i; i < N2; ++i) {
            operatorNetworkReceivers[i] = IRouterBurner.OperatorNetworkReceiver({
                network: address(i * 2 + 3),
                operator: address(i * 2 + 1),
                receiver: address(0)
            });
        }

        IRouterBurner.InitParams memory initParams = IRouterBurner.InitParams({
            owner: owner,
            receiverSetEpochsDelay: receiverSetEpochsDelay,
            globalReceiver: globalReceiver,
            networkReceivers: networkReceivers,
            operatorNetworkReceivers: operatorNetworkReceivers
        });

        vm.expectRevert(IRouterBurner.InvalidReceiver.selector);
        address routerBurnerAddress = routerBurnerFactory.create(initParams);
    }

    function test_CreateRevertDuplicateNetworkReceiver(uint256 receiverSetEpochsDelay, address globalReceiver) public {
        receiverSetEpochsDelay = bound(receiverSetEpochsDelay, 3, type(uint256).max);

        address routerBurnerImplementation = address(new RouterBurner(address(vaultFactory)));
        routerBurnerFactory = new RouterBurnerFactory(routerBurnerImplementation);

        uint160 N1 = 10;
        IRouterBurner.NetworkReceiver[] memory networkReceivers = new IRouterBurner.NetworkReceiver[](N1);
        uint160 N2 = 20;
        IRouterBurner.OperatorNetworkReceiver[] memory operatorNetworkReceivers =
            new IRouterBurner.OperatorNetworkReceiver[](N2);

        for (uint160 i; i < N1; ++i) {
            networkReceivers[i] = IRouterBurner.NetworkReceiver({network: address(1), receiver: address(i * 2 + 1)});
        }

        for (uint160 i; i < N2; ++i) {
            operatorNetworkReceivers[i] = IRouterBurner.OperatorNetworkReceiver({
                network: address(i * 2 + 3),
                operator: address(i * 2 + 1),
                receiver: address(i * 2 + 2)
            });
        }

        IRouterBurner.InitParams memory initParams = IRouterBurner.InitParams({
            owner: owner,
            receiverSetEpochsDelay: receiverSetEpochsDelay,
            globalReceiver: globalReceiver,
            networkReceivers: networkReceivers,
            operatorNetworkReceivers: operatorNetworkReceivers
        });

        vm.expectRevert(IRouterBurner.DuplicateNetworkReceiver.selector);
        address routerBurnerAddress = routerBurnerFactory.create(initParams);
    }

    function test_CreateRevertDuplicateOperatorNetworkReceiver(
        uint256 receiverSetEpochsDelay,
        address globalReceiver
    ) public {
        receiverSetEpochsDelay = bound(receiverSetEpochsDelay, 3, type(uint256).max);

        address routerBurnerImplementation = address(new RouterBurner(address(vaultFactory)));
        routerBurnerFactory = new RouterBurnerFactory(routerBurnerImplementation);

        uint160 N1 = 10;
        IRouterBurner.NetworkReceiver[] memory networkReceivers = new IRouterBurner.NetworkReceiver[](N1);
        uint160 N2 = 20;
        IRouterBurner.OperatorNetworkReceiver[] memory operatorNetworkReceivers =
            new IRouterBurner.OperatorNetworkReceiver[](N2);

        for (uint160 i; i < N1; ++i) {
            networkReceivers[i] =
                IRouterBurner.NetworkReceiver({network: address(i * 2 + 3), receiver: address(i * 2 + 1)});
        }

        for (uint160 i; i < N2; ++i) {
            operatorNetworkReceivers[i] = IRouterBurner.OperatorNetworkReceiver({
                network: address(1),
                operator: address(1),
                receiver: address(i * 2 + 2)
            });
        }

        IRouterBurner.InitParams memory initParams = IRouterBurner.InitParams({
            owner: owner,
            receiverSetEpochsDelay: receiverSetEpochsDelay,
            globalReceiver: globalReceiver,
            networkReceivers: networkReceivers,
            operatorNetworkReceivers: operatorNetworkReceivers
        });

        vm.expectRevert(IRouterBurner.DuplicateOperatorNetworkReceiver.selector);
        address routerBurnerAddress = routerBurnerFactory.create(initParams);
    }

    function _getVaultWithDelegatorWithSlasher(
        address burner
    ) internal returns (Vault, FullRestakeDelegator, Slasher) {
        address[] memory networkLimitSetRoleHolders = new address[](1);
        networkLimitSetRoleHolders[0] = alice;
        address[] memory operatorNetworkSharesSetRoleHolders = new address[](1);
        operatorNetworkSharesSetRoleHolders[0] = alice;
        (address vault_, address delegator_, address slasher_) = vaultConfigurator.create(
            IVaultConfigurator.InitParams({
                version: vaultFactory.lastVersion(),
                owner: alice,
                vaultParams: abi.encode(
                    IVault.InitParams({
                        collateral: address(collateral),
                        burner: burner,
                        epochDuration: 7 days,
                        depositWhitelist: false,
                        isDepositLimit: false,
                        depositLimit: 0,
                        defaultAdminRoleHolder: alice,
                        depositWhitelistSetRoleHolder: alice,
                        depositorWhitelistRoleHolder: alice,
                        isDepositLimitSetRoleHolder: alice,
                        depositLimitSetRoleHolder: alice
                    })
                ),
                delegatorIndex: 0,
                delegatorParams: abi.encode(
                    INetworkRestakeDelegator.InitParams({
                        baseParams: IBaseDelegator.BaseParams({
                            defaultAdminRoleHolder: alice,
                            hook: address(0),
                            hookSetRoleHolder: alice
                        }),
                        networkLimitSetRoleHolders: networkLimitSetRoleHolders,
                        operatorNetworkSharesSetRoleHolders: operatorNetworkSharesSetRoleHolders
                    })
                ),
                withSlasher: true,
                slasherIndex: 0,
                slasherParams: abi.encode(ISlasher.InitParams({baseParams: IBaseSlasher.BaseParams({isBurnerHook: true})}))
            })
        );

        return (Vault(vault_), FullRestakeDelegator(delegator_), Slasher(slasher_));
    }
}