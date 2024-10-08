// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {ISUSDe} from "../../../interfaces/burners/sUSDe/ISUSDe.sol";

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract sUSDe_Miniburner is OwnableUpgradeable {
    address private immutable _COLLATERAL;

    address private immutable _USDE;

    constructor(
        address collateral
    ) {
        _disableInitializers();

        _COLLATERAL = collateral;
        _USDE = ISUSDe(collateral).asset();
    }

    function initialize(
        uint256 amount
    ) external initializer {
        __Ownable_init(msg.sender);

        ISUSDe(_COLLATERAL).cooldownShares(amount);
    }

    function triggerClaim() external onlyOwner {
        ISUSDe(_COLLATERAL).unstake(owner());
    }
}
