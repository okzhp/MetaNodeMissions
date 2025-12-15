// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MyERC20 {
    string public name;
    string public symbol;
    uint256 public totalSupply;
    uint8 public immutable decimals;

    address public immutable owner;

    mapping(address => uint256) private balanceMap;
    mapping(address =>mapping(address => uint256)) private allowanceMap;

    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Approval(address indexed from, address indexed to, uint256 amount);

    modifier onlyOwner() {
        require(msg.sender == owner, "not owner");
        _;
    }

    constructor(string memory _name, string memory _symbol, uint256 _totalSupply, uint8 _decimals) {
        require(_decimals <= 18, "invalid _decimals");
        
        name = _name;
        symbol = _symbol;
        totalSupply = _totalSupply * 10 ** _decimals;
        decimals = _decimals;

        owner = msg.sender;
        balanceMap[msg.sender] = totalSupply;

        emit Transfer(address(0), msg.sender, totalSupply);
    }


    function balanceOf(address account) public view returns (uint256) {
        return balanceMap[account];
    }

    function allowance(address from, address spender) public view returns (uint256) {
        return allowanceMap[from][spender];
    }

    function transfer(address to, uint256 amount) public returns (bool) {
        require(to != address(0), "invalid to account");
        require(amount > 0, "invalid amount");
        require(balanceMap[msg.sender] >= amount, "insufficient balance!");

        balanceMap[msg.sender] -= amount;
        balanceMap[to] += amount;

        emit Transfer(msg.sender, to, amount);

        return true;
    }

    function approve(address to, uint256 amount) external returns (bool) {
        require(to != address(0), "invalid to account");
        require(amount > 0, "invalid amount");

        allowanceMap[msg.sender][to] = amount;

        emit Approval(msg.sender, to, amount);

        return true;
    }

    function transferFrom(address from, address to, uint256 amount) public returns (bool) {
        require(from != address(0), "invalid from account");
        require(to != address(0), "invalid to account");
        require(amount > 0, "invalid amount");

        require(balanceMap[from] >= amount, "insufficient balance");
        require(allowanceMap[from][msg.sender] >= amount, "insufficient allowance!");

        balanceMap[from] -= amount;
        balanceMap[to] += amount;

        allowanceMap[from][msg.sender] -= amount;

        emit Transfer(from, to, amount);

        return true;

    }

    function mint(address to, uint256 amount) public onlyOwner {
        require(to != address(0), "invalid to account");
        require(amount > 0, "invalid amount");

        balanceMap[to] += amount;
        totalSupply += amount;

        emit Transfer(address(0), to, amount);
    }

    function burn(uint256 amount) public {
        require(amount > 0, "invalid amount");
        require(balanceMap[msg.sender] >= amount, "insufficient balance");

        balanceMap[msg.sender] -= amount;
        totalSupply -= amount;

        emit Transfer(msg.sender, address(0), amount);
    }

}