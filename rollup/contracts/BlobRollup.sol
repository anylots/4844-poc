// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BlobHashGetterFactory {
    constructor() payable {
        bytes memory code = hex"6000354960005260206000F3";
        uint256 size = code.length;
        assembly {
            return(add(code, 0x020), size)
        }
    }
}

library BlobHashGetter {
    function getBlobHash(
        address getter,
        uint256 idx
    ) internal view returns (bytes32) {
        bool success;
        bytes32 blobHash;
        assembly {
            mstore(0x0, idx)
            success := staticcall(gas(), getter, 0x0, 0x20, 0x0, 0x20)
            blobHash := mload(0x0)
        }
        require(success, "failed to get blob hash");
        return blobHash;
    }
}

contract BlobRollup {
    struct BatchData {
        uint64 blockNumber; // Block number
        bytes blockWitness; // Contains each block
        bytes32 preStateRoot; // Pre-state root
        bytes32 postStateRoot; // Post-state root
        bytes32 withdrawalRoot; // withdraw trie tree
        BatchSignature signature; // Signature of the batch
    }

    struct BatchSignature {
        bytes[] signers; // all signers
        bytes signature; // aggregate signature
    }

    struct WithdrawalRootStore {
        uint64 batchIndex;
        uint256 timestamp;
    }

    struct BatchStore {
        uint64 blockNumber;
        bytes32 withdrawalRoot;
        bytes32 commitment;
        bytes32 stateRoot;
        uint256 originTimestamp;
    }
    /**
     * @notice Store the withdrawalRoot of each batch.
     */
    mapping(bytes32 => WithdrawalRootStore) public withdrawalRoots;

    /**
     * @notice Store batchs.
     */
    mapping(uint64 => BatchStore) public storageBatchs;

    /**
     * @notice Last batch sent by the sequencers
     */
    uint64 public lastBatchSequenced;

    /**
     * @notice The height of the latest batch that has been received.
     */
    uint64 public lastL2BlockNumber;

    address public hashGetter;
    bytes32 public dataHash;

    /**
     * @notice Store versionHash.
     */
    mapping(uint64 => bytes32) public batchToVersionHash;

    event Challenge(uint64 batchNum);

    constructor() {
        hashGetter = address(new BlobHashGetterFactory());
    }

    function submit(BatchData[] calldata batches) public {
        bytes32 _dataHash = BlobHashGetter.getBlobHash(hashGetter, uint256(0));
        dataHash = _dataHash;

        uint256 batchesNum = batches.length;
        uint64 currentBatchSequenced = lastBatchSequenced;

        for (uint256 i = 0; i < batchesNum; i++) {
            uint256 chainId = 99;
            (uint256 MAX_TXS, uint256 MAX_CALLDATA) = (14, 69750);
            uint256[] memory publicInput = _buildCommitment(
                MAX_TXS,
                MAX_CALLDATA,
                chainId,
                batches[i].blockWitness,
                true
            );

            bytes32 stateRoot = batches[i].postStateRoot;

            bytes32 commitmentHash;
            assembly {
                commitmentHash := keccak256(
                    add(publicInput, 32),
                    mul(mload(publicInput), 32)
                )
            }

            uint256 timeNow = block.timestamp;
            currentBatchSequenced++;

            // When this withdrawalRoot does not exist, store the relationship mapping between withdrawalRoot and batch Index
            if (withdrawalRoots[batches[i].withdrawalRoot].batchIndex == 0) {
                withdrawalRoots[
                    batches[i].withdrawalRoot
                ] = WithdrawalRootStore({
                    batchIndex: currentBatchSequenced,
                    timestamp: timeNow
                });
            }

            storageBatchs[currentBatchSequenced] = BatchStore(
                batches[i].blockNumber,
                batches[i].withdrawalRoot,
                commitmentHash,
                stateRoot,
                timeNow
            );
        }
        lastBatchSequenced = currentBatchSequenced;
        lastL2BlockNumber = batches[batchesNum - 1].blockNumber;
    }

    function challenge(uint64 batchNum) external {
        emit Challenge(batchNum);
    }

    function prove(BatchData[] calldata batches) external {
        
    }

    function _buildCommitment(
        uint256 MAX_TXS,
        uint256 MAX_CALLDATA,
        uint256 chainId,
        bytes calldata witness,
        bool clearMemory
    ) internal pure returns (uint256[] memory table) {
        // TODO
        return table;
    }

    function evaluation(bytes calldata input) public view {
        address precompile = address(0x0A);
        (bool ok, bytes memory output) = precompile.staticcall{gas: 50000}(
            input
        );
    }

    function getDataHash() public view returns (bytes32) {
        return dataHash;
    }
}
