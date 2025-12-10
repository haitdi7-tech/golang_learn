// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
//  创建一个名为Voting的合约，包含以下功能：
contract Voting{
    // 一个mapping来存储候选人的得票数
    mapping (address => uint256) votes;
    //记录投票者的投票情况
    mapping (address => address) users;
    address[] candidates;
    

    // 一个vote函数，允许用户投票给某个候选人
    function vote(address user,address candidate)external{
        if (votes[candidate] == 0){
            candidates.push(candidate);
        }
        users[user]=candidate;
    }

    // 一个getVotes函数，返回某个候选人的得票数
    function getVotes(address candidate)external view returns (uint256){
        return votes[candidate];
    }

    // 一个resetVotes函数，重置所有候选人的得票数
    function resetVotes()external {
        uint256 len = candidates.length;
        for (uint256 i = 0; i < len; i++) 
        {
            votes[candidates[i]]=0;
        }
    }

//     反转字符串 (Reverse String)
// 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
    function reverse(string memory str) public returns(string memory,uint len){
        bytes memory bs = bytes(str);
        uint256 len = bs.length;
        bytes memory rbs = new bytes(len);
        
        for(uint i = 0;i < len; i++){
            rbs[i] = bs[len - 1 -i];
        }
        return (string(rbs),len);
    }
}
