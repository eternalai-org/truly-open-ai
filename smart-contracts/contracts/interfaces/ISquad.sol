// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ISquad {
    event MoveAgentToSquad(uint256 indexed toSquadId, uint256 agentIds);
    event SquadTransferred(
        address indexed from,
        address indexed to,
        uint256 indexed squadId
    );

    error Unauthorized();
    error InvalidSquadId();
    error InvalidData();
    error InvalidAgentId();

    function moveAgentToSquad(
        address _caller,
        uint256 _agentId,
        uint256 _toSquadId
    ) external;
}
