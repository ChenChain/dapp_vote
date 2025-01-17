// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
contract Audi is ERC20 {

    address public audiOwner;
    uint256 public audiTotalSupply;

    // audi token
    constructor() ERC20("Audi", "ATK") {
        audiTotalSupply = 1000*1000*1000*1000;
        audiOwner = msg.sender;
        // 通过 _mint 函数铸造初始供应量的代币到部署合约的地址
        _mint(msg.sender, audiTotalSupply);
    }

    function Wish() public pure returns (string memory) {
        return "Wish U have one Audi Car";
    }
}