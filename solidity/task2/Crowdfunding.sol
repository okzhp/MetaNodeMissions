// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract Crowdfunding is Ownable{
    enum State {
        Preparing,
        Funding,
        Success,
        Fail
    }

    //众筹名称
    string public name;
    //描述
    string public description;
    //众筹目标金额
    uint256 public immutable targetAmount;
    //众筹截止时间
    uint256 public deadLineTime;
    //众筹持续时间
    uint256 public immutable durationMinutes;

    //众筹状态
    State public state;

    //捐赠详情
    mapping(address => uint256) fundMap;
    //当前众筹金额
    uint256 public currentAmount;

    //最少捐献金额 0.001 eth
    uint256 public constant MIN_FUND = 0.001 ether;

    // 事件：记录状态变化和投资
    event StateChanged(State newState);
    event Contributed(address indexed contributor, uint256 amount);


    modifier inState(State _state) {
        require(state == _state, "invalid funding state");
        _;
    }

    constructor(string memory _name, string memory _description, uint256 _targetAmount, uint256 _durationMinutes) Ownable(msg.sender) {
        require(_targetAmount > 0, "targetAmount must bigger than 0");
        require(_durationMinutes > 0, "durationMinutes must bigger than 0");

        name = _name;
        description = _description;
        targetAmount = _targetAmount;
        durationMinutes = _durationMinutes;
        state = State.Preparing;
        emit StateChanged(State.Preparing);
    }

    //开始募集
    function startFunding() public onlyOwner inState(State.Preparing) {
        deadLineTime = block.timestamp + (durationMinutes * 1 minutes);
        state = State.Funding;
        emit StateChanged(State.Funding);
    }

    //募集资金
    function funding() public payable inState(State.Funding) {
        require(msg.value >= MIN_FUND, "smaller than MIN_FUND");
        require(block.timestamp < deadLineTime, "over deadLineTime");

        fundMap[msg.sender] += msg.value;
        currentAmount += msg.value;

        emit Contributed(msg.sender, msg.value);
    }

    //众筹成功，合约所有者提取众筹资金
    function withDraw() public onlyOwner inState(State.Success) {
        (bool success,) = payable(msg.sender).call{value: address(this).balance}("");
        require(success, "withDraw transfer fail");
    }

    
    function finalize() public inState(State.Funding) {
        require(block.timestamp >= deadLineTime, "before _deadLineTime");
        if(currentAmount >= targetAmount) {
            state = State.Success;
            emit StateChanged(State.Success);
        } else {
            state = State.Fail;
            emit StateChanged(State.Fail);
        }
    }

    //众筹失败，用户提取资金
    function reFund() public inState(State.Fail) {
        require(fundMap[msg.sender] > 0, "insufficient balance");
        uint256 value = fundMap[msg.sender];
        require(address(this).balance >= value, "insufficient contract balance");

        fundMap[msg.sender] = 0;
        (bool success, ) = payable(msg.sender).call{value: value}("");
        require(success, "reFund transfer fail");
    }

    function queryFund(address _address) public view returns (uint256) {
        return fundMap[_address];
    }

}