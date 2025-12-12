// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    mapping(address => uint) votes;
    address[] voteAddress;

    function vote(address to) public {
        if(votes[to] == 0) {
            voteAddress.push(to);
        }
        votes[to] += 1;
    }

    function getVotes(address account) public view returns (uint)  {
        return votes[account];
    }

    function restoreVotes() public {
        uint len = voteAddress.length;
        for (uint i = 0; i < len; i++) 
        {
            votes[voteAddress[i]] = 0;
        }
        delete voteAddress;
    }

}