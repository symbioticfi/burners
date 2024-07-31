// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {SelfDestruct} from "src/contracts/SelfDestruct.sol";

import {IrETH_Burner} from "src/interfaces/burners/rETH/IrETH_Burner.sol";
import {IRocketTokenRETH} from "src/interfaces/burners/rETH/IRocketTokenRETH.sol";

contract rETH_Burner is IrETH_Burner {
    /**
     * @inheritdoc IrETH_Burner
     */
    address public immutable COLLATERAL;

    constructor(address collateral) {
        COLLATERAL = collateral;
    }

    /**
     * @inheritdoc IrETH_Burner
     */
    function triggerBurn(uint256 amount) external {
        IRocketTokenRETH(COLLATERAL).burn(amount);

        uint256 ethToBurn = address(this).balance;
        new SelfDestruct{value: ethToBurn}();

        emit TriggerBurn(msg.sender, amount, ethToBurn);
    }

    receive() external payable {}
}