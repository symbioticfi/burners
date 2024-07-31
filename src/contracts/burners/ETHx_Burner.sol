// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {SelfDestruct} from "src/contracts/SelfDestruct.sol";
import {UintRequests} from "src/contracts/UintRequests.sol";

import {IETHx_Burner} from "src/interfaces/burners/ETHx/IETHx_Burner.sol";
import {IStaderStakePoolsManager} from "src/interfaces/burners/ETHx/IStaderStakePoolsManager.sol";
import {IUserWithdrawalManager} from "src/interfaces/burners/ETHx/IUserWithdrawalManager.sol";
import {IStaderConfig} from "src/interfaces/burners/ETHx/IStaderConfig.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

contract ETHx_Burner is UintRequests, IETHx_Burner {
    using Math for uint256;

    /**
     * @inheritdoc IETHx_Burner
     */
    address public immutable COLLATERAL;

    /**
     * @inheritdoc IETHx_Burner
     */
    address public immutable STADER_CONFIG;

    /**
     * @inheritdoc IETHx_Burner
     */
    address public immutable USER_WITHDRAW_MANAGER;

    /**
     * @inheritdoc IETHx_Burner
     */
    address public immutable STAKE_POOLS_MANAGER;

    constructor(address collateral, address staderConfig) {
        COLLATERAL = collateral;

        STADER_CONFIG = staderConfig;
        USER_WITHDRAW_MANAGER = IStaderConfig(STADER_CONFIG).getUserWithdrawManager();
        STAKE_POOLS_MANAGER = IStaderConfig(STADER_CONFIG).getStakePoolManager();

        IERC20(COLLATERAL).approve(USER_WITHDRAW_MANAGER, type(uint256).max);
    }

    /**
     * @inheritdoc IETHx_Burner
     */
    function triggerWithdrawal(uint256 maxRequests) external returns (uint256 firstRequestId, uint256 lastRequestId) {
        uint256 amount = IERC20(COLLATERAL).balanceOf(address(this));

        uint256 maxWithdrawalAmount = IStaderStakePoolsManager(STAKE_POOLS_MANAGER).previewDeposit(
            IStaderConfig(STADER_CONFIG).getMaxWithdrawAmount()
        );
        uint256 minWithdrawalAmount = IStaderStakePoolsManager(STAKE_POOLS_MANAGER).previewDeposit(
            IStaderConfig(STADER_CONFIG).getMinWithdrawAmount()
        ) + 1;

        uint256 requests = amount / maxWithdrawalAmount;
        if (amount % maxWithdrawalAmount >= minWithdrawalAmount) {
            requests += 1;
        }
        requests = Math.min(requests, maxRequests);

        if (requests == 0) {
            revert InsufficientWithdrawal();
        }

        uint256 requestsMinusOne = requests - 1;
        firstRequestId = IUserWithdrawalManager(USER_WITHDRAW_MANAGER).nextRequestId();
        lastRequestId = firstRequestId + requestsMinusOne;
        uint256 requestId = firstRequestId;
        for (; requestId < lastRequestId; ++requestId) {
            _addRequestId(requestId);
            IUserWithdrawalManager(USER_WITHDRAW_MANAGER).requestWithdraw(maxWithdrawalAmount, address(this));
        }
        _addRequestId(requestId);
        IUserWithdrawalManager(USER_WITHDRAW_MANAGER).requestWithdraw(
            Math.min(amount - requestsMinusOne * maxWithdrawalAmount, maxWithdrawalAmount), address(this)
        );

        emit TriggerWithdrawal(msg.sender, firstRequestId, lastRequestId);

        return (firstRequestId, lastRequestId);
    }

    /**
     * @inheritdoc IETHx_Burner
     */
    function triggerBurn(uint256 requestId) external {
        _removeRequestId(requestId);

        IUserWithdrawalManager(USER_WITHDRAW_MANAGER).claim(requestId);

        new SelfDestruct{value: address(this).balance}();

        emit TriggerBurn(msg.sender, requestId);
    }

    receive() external payable {}
}