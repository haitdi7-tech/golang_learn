// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract NewToken is IERC20 {
    mapping(address => uint256) balances; //账户余额
    mapping(address account => mapping(address sender => uint256 amount))
        private _allowances;
    address owner;
    uint256 private _totalSupply;

    event TransferNew(address indexed from, address indexed to, uint256 amount);
    event ApprovalNew(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    constructor() {
        owner = msg.sender;
    }

    function _mint(uint256 amount) external {
        require(msg.sender == owner, "must be owner");
        _totalSupply += amount;
        balances[msg.sender] += amount;
    }

    function balanceOf(
        address account
    ) external view override returns (uint256) {
        return balances[account];
    }

    function transfer(
        address to,
        uint256 amount
    ) external override returns (bool) {
        uint256 balance = balances[msg.sender];
        require(balance < amount, "Have not enouth amount");
        balances[msg.sender] = balance - amount;
        balances[to] += amount;
        emit TransferNew(msg.sender, to, amount);
    }

    function approve(
        address spender,
        uint256 value
    ) external override returns (bool) {
        _allowances[msg.sender][spender] = value;
        emit ApprovalNew(msg.sender, spender, value);
        return true;
    }

    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external override returns (bool) {
        require(
            _allowances[from][msg.sender] >= value,
            "not have enough amount"
        );
        balances[from] -= value;
        balances[to] += value;
    }

    function allowance(
        address owner,
        address spender
    ) external view override returns (uint256) {
        return _allowances[owner][spender];
    }

    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }
}

