// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./SymbioticBurnersIntegration.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {console2} from "forge-std/Test.sol";

contract SymbioticBurnersIntegrationExample is SymbioticBurnersIntegration {
    using SymbioticSubnetwork for bytes32;
    using SymbioticSubnetwork for address;

    address[] public networkVaults;

    address[] public confirmedNetworkVaults;
    mapping(address vault => address[]) public confirmedNetworkOperators;
    mapping(address vault => bytes32[]) public neighborNetworks;

    uint256 public SELECT_OPERATOR_CHANCE = 1; // lower -> higher probability

    function setUp() public override {
        SYMBIOTIC_CORE_PROJECT_ROOT = "lib/core/";
        SYMBIOTIC_BURNERS_PROJECT_ROOT = "";
        // vm.selectFork(vm.createFork(vm.rpcUrl("holesky")));
        // SYMBIOTIC_INIT_BLOCK = 2_727_202;
        // SYMBIOTIC_CORE_USE_EXISTING_DEPLOYMENT = true;
        // SYMBIOTIC_BURNERS_USE_EXISTING_DEPLOYMENT = true;

        SYMBIOTIC_CORE_NUMBER_OF_STAKERS = 10;
        SYMBIOTIC_BURNERS_NEED_BURNERS = false;

        super.setUp();
    }

    function test_NetworkWithCustomBurner() public {
        address middleware = address(111);
        Vm.Wallet memory network = _getNetworkWithMiddleware_SymbioticCore(middleware);
        uint96 identifier = 0;
        address collateral = tokens_SymbioticCore[0];
        bytes32 subnetwork = network.addr.subnetwork(identifier);
        address networkTreasury = address(123_321);

        console2.log("Network:", network.addr);
        console2.log("Identifier:", identifier);
        console2.log("Collateral:", collateral);
        console2.log("Network Treasury:", networkTreasury);

        for (uint256 i; i < vaults_SymbioticCore.length; ++i) {
            address slasher = ISymbioticVault(vaults_SymbioticCore[i]).slasher();
            if (
                ISymbioticVault(vaults_SymbioticCore[i]).collateral() == collateral && slasher != address(0)
                    && ISymbioticBaseSlasher(slasher).TYPE() == 0
            ) {
                networkVaults.push(vaults_SymbioticCore[i]);
            }
        }

        console2.log("Network Vaults:", networkVaults.length);

        for (uint256 i; i < networkVaults.length; ++i) {
            _networkSetMaxNetworkLimitRandom_SymbioticCore(network.addr, networkVaults[i], identifier);
            if (_delegateToNetworkTry_SymbioticCore(networkVaults[i], subnetwork)) {
                confirmedNetworkVaults.push(networkVaults[i]);
            }
        }

        console2.log("Confirmed Network Vaults:", confirmedNetworkVaults.length);
        console2.log("Operators:", operators_SymbioticCore.length);

        for (uint256 i; i < confirmedNetworkVaults.length; ++i) {
            for (uint256 j; j < operators_SymbioticCore.length; ++j) {
                if (
                    ISymbioticOptInService(symbioticCore.operatorVaultOptInService).isOptedIn(
                        operators_SymbioticCore[j].addr, confirmedNetworkVaults[i]
                    ) && _randomChoice_Symbiotic(SELECT_OPERATOR_CHANCE)
                ) {
                    _operatorOptInWeak_SymbioticCore(operators_SymbioticCore[j].addr, network.addr);
                    if (
                        _delegateToOperatorTry_SymbioticCore(
                            confirmedNetworkVaults[i], subnetwork, operators_SymbioticCore[j].addr
                        )
                    ) {
                        confirmedNetworkOperators[confirmedNetworkVaults[i]].push(operators_SymbioticCore[j].addr);
                    }
                }
            }

            console2.log("Confirmed Network Operators:", confirmedNetworkOperators[confirmedNetworkVaults[i]].length);
        }

        for (uint256 i; i < confirmedNetworkVaults.length; ++i) {
            console2.log("Confirmed Network Vault:", confirmedNetworkVaults[i]);
            console2.log("Confirmed Network Operators:", confirmedNetworkOperators[confirmedNetworkVaults[i]].length);
            for (uint256 j; j < confirmedNetworkOperators[confirmedNetworkVaults[i]].length; ++j) {
                console2.log("Operator:", confirmedNetworkOperators[confirmedNetworkVaults[i]][j]);
                console2.log(
                    "Stake:",
                    ISymbioticBaseDelegator(ISymbioticVault(confirmedNetworkVaults[i]).delegator()).stake(
                        subnetwork, confirmedNetworkOperators[confirmedNetworkVaults[i]][j]
                    )
                );
            }
        }

        for (uint256 i; i < confirmedNetworkVaults.length; ++i) {
            address burner = ISymbioticVault(confirmedNetworkVaults[i]).burner();
            if (
                symbioticBurnerRouterFactory.isEntity(burner)
                    && ISymbioticBurnerRouter(burner).networkReceiver(network.addr) == address(0)
            ) {
                _curatorSetNetworkReceiver_SymbioticBurners(
                    Ownable(burner).owner(), burner, network.addr, networkTreasury
                );
            }
        }

        _skipBlocks_Symbiotic(SYMBIOTIC_BURNERS_MAX_DELAY); // for simplicity

        for (uint256 i; i < confirmedNetworkVaults.length; ++i) {
            address burner = ISymbioticVault(confirmedNetworkVaults[i]).burner();
            if (symbioticBurnerRouterFactory.isEntity(burner)) {
                (address pendingReceiver,) = ISymbioticBurnerRouter(burner).pendingNetworkReceiver(network.addr);
                if (pendingReceiver != address(0)) {
                    _anyoneAcceptNetworkReceiver_SymbioticBurners(address(this), burner, network.addr);
                }
            }
        }

        uint48 captureTimestamp = uint48(vm.getBlockTimestamp() - 1);
        for (uint256 i; i < confirmedNetworkVaults.length; ++i) {
            for (uint256 j; j < confirmedNetworkOperators[confirmedNetworkVaults[i]].length; ++j) {
                address slasher = ISymbioticVault(confirmedNetworkVaults[i]).slasher();
                uint256 slashableStake = ISymbioticBaseSlasher(slasher).slashableStake(
                    subnetwork, confirmedNetworkOperators[confirmedNetworkVaults[i]][j], captureTimestamp, new bytes(0)
                );
                if (slashableStake == 0) {
                    continue;
                }
                _slash_SymbioticCore({
                    who: middleware,
                    vault: confirmedNetworkVaults[i],
                    subnetwork: subnetwork,
                    operator: confirmedNetworkOperators[confirmedNetworkVaults[i]][j],
                    amount: slashableStake,
                    captureTimestamp: captureTimestamp
                });
            }
        }

        for (uint256 i; i < confirmedNetworkVaults.length; ++i) {
            address burner = ISymbioticVault(confirmedNetworkVaults[i]).burner();

            if (symbioticBurnerRouterFactory.isEntity(burner)) {
                console2.log("Burner Router:", burner);

                address receiver = ISymbioticBurnerRouter(burner).networkReceiver(network.addr);
                uint256 balanceBefore = IERC20(collateral).balanceOf(receiver);
                if (ISymbioticBurnerRouter(burner).balanceOf(receiver) > 0) {
                    _anyoneTriggerTransfer_SymbioticBurners(address(this), burner, receiver);
                }
                uint256 balanceAfter = IERC20(collateral).balanceOf(receiver);
                console2.log("Collateral received by:", receiver, "-", balanceAfter - balanceBefore);
            } else {
                console2.log("Burner:", burner);
                console2.log("Burner's balance:", IERC20(collateral).balanceOf(burner));
            }
        }
    }
}
