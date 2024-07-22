// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {SelfDestruct} from "src/contracts/SelfDestruct.sol";

import {IDC_swETH_Burner} from "src/interfaces/burners/DC_swETH/IDC_swETH_Burner.sol";
import {ISwEXIT} from "src/interfaces/burners/DC_swETH/ISwEXIT.sol";

import {IDefaultCollateral} from "@symbiotic/collateral/interfaces/defaultCollateral/IDefaultCollateral.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {IERC721Receiver} from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

contract DC_swETH_Burner is IDC_swETH_Burner, IERC721Receiver {
    using Math for uint256;
    using EnumerableSet for EnumerableSet.UintSet;

    /**
     * @inheritdoc IDC_swETH_Burner
     */
    address public immutable COLLATERAL;

    /**
     * @inheritdoc IDC_swETH_Burner
     */
    address public immutable ASSET;

    /**
     * @inheritdoc IDC_swETH_Burner
     */
    address public immutable SWEXIT;

    EnumerableSet.UintSet private _requestIds;

    constructor(address collateral, address swEXIT) {
        COLLATERAL = collateral;

        ASSET = IDefaultCollateral(collateral).asset();

        SWEXIT = swEXIT;

        IERC20(ASSET).approve(SWEXIT, type(uint256).max);
    }

    /**
     * @inheritdoc IDC_swETH_Burner
     */
    function requestIdsLength() external view returns (uint256) {
        return _requestIds.length();
    }

    /**
     * @inheritdoc IDC_swETH_Burner
     */
    function requestIds(uint256 index, uint256 maxRequestIds) external view returns (uint256[] memory requestIds_) {
        uint256 length = Math.min(index + maxRequestIds, _requestIds.length()) - index;

        requestIds_ = new uint256[](length);
        for (uint256 i; i < length;) {
            requestIds_[i] = _requestIds.at(index);
            unchecked {
                ++i;
                ++index;
            }
        }
    }

    /**
     * @inheritdoc IDC_swETH_Burner
     */
    function triggerWithdrawal(uint256 maxRequests) external returns (uint256 firstRequestId, uint256 lastRequestId) {
        IDefaultCollateral(COLLATERAL).withdraw(address(this), IERC20(COLLATERAL).balanceOf(address(this)));
        uint256 amount = IERC20(ASSET).balanceOf(address(this));

        uint256 maxWithdrawalAmount = ISwEXIT(SWEXIT).withdrawRequestMaximum();
        uint256 requests = amount / maxWithdrawalAmount;
        if (amount % maxWithdrawalAmount >= ISwEXIT(SWEXIT).withdrawRequestMinimum()) {
            requests += 1;
        }
        requests = Math.min(requests, maxRequests);

        if (requests == 0) {
            revert InsufficientWithdrawal();
        }

        uint256 requestsMinusOne = requests - 1;
        firstRequestId = ISwEXIT(SWEXIT).getLastTokenIdCreated() + 1;
        lastRequestId = firstRequestId + requestsMinusOne;
        uint256 requestId = firstRequestId;
        for (; requestId < lastRequestId; ++requestId) {
            _requestIds.add(requestId);
            ISwEXIT(SWEXIT).createWithdrawRequest(maxWithdrawalAmount);
        }
        _requestIds.add(requestId);
        ISwEXIT(SWEXIT).createWithdrawRequest(
            Math.min(amount - requestsMinusOne * maxWithdrawalAmount, maxWithdrawalAmount)
        );

        emit TriggerWithdrawal(msg.sender, firstRequestId, lastRequestId);

        return (firstRequestId, lastRequestId);
    }

    /**
     * @inheritdoc IDC_swETH_Burner
     */
    function triggerBurn(uint256 requestId) external {
        if (!_requestIds.remove(requestId)) {
            revert InvalidRequestId();
        }

        ISwEXIT(SWEXIT).finalizeWithdrawal(requestId);

        new SelfDestruct{value: address(this).balance}();

        emit TriggerBurn(msg.sender, requestId);
    }

    /**
     * @inheritdoc IERC721Receiver
     */
    function onERC721Received(address, address, uint256, bytes calldata) external returns (bytes4) {
        return IERC721Receiver.onERC721Received.selector;
    }

    receive() external payable {}
}