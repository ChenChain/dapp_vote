// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Vote {
    // 候选人
    struct Candidate {
        string name;
        uint voteCount;
    }

    address public owner;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action");
        _;
    }

    Candidate[] public candidates;
    mapping(address => bool) public hasVoted;

    event VoteCast(address indexed voter, uint candidateIndex);

    // 添加候选人 只有部署者能添加
    function addCandidate(string memory _name) public onlyOwner {
        candidates.push(Candidate({name: _name, voteCount: 0}));
    }

    // 投票
    function vote(uint _candidateIndex) public {
        require(!hasVoted[msg.sender], "You have already voted");
        require(_candidateIndex < candidates.length, "Invalid candidate index");

        hasVoted[msg.sender] = true;
        candidates[_candidateIndex].voteCount += 1;

        emit VoteCast(msg.sender, _candidateIndex); // 记录事件日志
    }

    // 获取候选人信息
    function getCandidates() public view returns (Candidate[] memory) {
        return candidates;
    }
}
