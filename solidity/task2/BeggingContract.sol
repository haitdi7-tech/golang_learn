// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BeggingContract{
    mapping (address => uint256) donates;
    uint256 totalDonate;
    address owner;

    event Donation(address indexed from, uint256 amount);

    constructor(){
        owner = msg.sender;
    }



    modifier onlyOwner() {
        //调用者必须是合约所有者
        require(msg.sender == owner,"Caller is not the owner");
        _;
    }

    //向合约转账
    function donate() public payable  {
        donates[msg.sender] += msg.value;
        totalDonate += msg.value;
        emit Donation(msg.sender, msg.value);
    }

    function withdraw() public onlyOwner {
        //uint256 = totalDonate;
        payable(owner).transfer(totalDonate);
        totalDonate = 0;
    }

    function getDonation(address donater) public view returns (uint256){
        return donates[donater];
    }
}
