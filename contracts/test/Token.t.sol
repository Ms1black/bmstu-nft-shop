// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/Token.sol";

contract TokenTest is Test {
    MyToken token;
    address owner = address(1);

    function setUp() public {
        vm.prank(owner);
        token = new MyToken(1000); // 1000 токенов на старте
    }

    function testInitialBalance() public {
        assertEq(token.balanceOf(owner), 1000 * (10 ** token.decimals()));
    }

    function testTransferAllToOwner() public {
        address receiver = address(2);
        uint256 balance = token.balanceOf(owner);

        vm.prank(owner);
        token.transfer(receiver, balance);

        assertEq(token.balanceOf(receiver), balance);
        assertEq(token.balanceOf(owner), 0);
    }
}
