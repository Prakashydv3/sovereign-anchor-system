// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title AnchorRegistry
/// @notice Immutable L1 anchor storage. No execution logic. Storage only.
contract AnchorRegistry {

    struct Anchor {
        bytes32 stateHash;
        bytes32 parentHash;
        uint256 timestamp;
    }

    // anchorId => Anchor
    mapping(bytes32 => Anchor) private anchors;

    event AnchorStored(bytes32 indexed anchorId, bytes32 stateHash, bytes32 parentHash, uint256 timestamp);

    /// @notice Store a new anchor. Reverts if anchorId already exists.
    function store(bytes32 anchorId, bytes32 stateHash, bytes32 parentHash) external {
        require(anchors[anchorId].timestamp == 0, "anchor exists");
        anchors[anchorId] = Anchor(stateHash, parentHash, block.timestamp);
        emit AnchorStored(anchorId, stateHash, parentHash, block.timestamp);
    }

    /// @notice Retrieve an anchor by its ID.
    function get(bytes32 anchorId) external view returns (bytes32 stateHash, bytes32 parentHash, uint256 timestamp) {
        Anchor memory a = anchors[anchorId];
        return (a.stateHash, a.parentHash, a.timestamp);
    }

    /// @notice Check if an anchor exists.
    function exists(bytes32 anchorId) external view returns (bool) {
        return anchors[anchorId].timestamp != 0;
    }
}
