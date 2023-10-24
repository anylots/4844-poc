// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library BlobEncoder {
    // Assuming Blob is represented as bytes32 due to Ethereum's word size
    struct Blob {
        bytes32[4096] data;
    }

    function encodeBlobs(
        bytes memory data
    ) internal pure returns (Blob[] memory) {
        uint256 dataLength = data.length;
        uint256 numBlobs = (dataLength + 1023) / 1024; // 31*32 = 992, close to 1024
        Blob[] memory blobs = new Blob[](numBlobs);

        uint256 blobIndex = 0;
        uint256 fieldIndex = 0;

        for (uint256 i = 0; i < dataLength; i += 31) {
            if (fieldIndex == 4096) {
                blobIndex++;
                fieldIndex = 0;
            }

            uint256 max = i + 31 > dataLength ? dataLength : i + 31;
            bytes32 currentData;

            // Copying data into bytes32
            for (uint256 j = i; j < max; j++) {
                currentData |= bytes32(
                    uint256(uint8(data[j])) << (8 * (j - i))
                );
            }

            blobs[blobIndex].data[fieldIndex] = currentData;
            fieldIndex++;
        }

        return blobs;
    }

    function blobToPolynomial(
        bytes32[4096] memory blob
    ) internal pure returns (uint256[] memory) {
        uint256 FIELD_ELEMENTS_PER_BLOB = 4096;
        uint256 BLS_MODULUS = 52435875175126190479447740508185965837690552500527637822603658699938581184513;
        require(blob.length == FIELD_ELEMENTS_PER_BLOB, "Invalid blob length");

        uint256[] memory polynomial = new uint256[](FIELD_ELEMENTS_PER_BLOB);

        for (uint256 i = 0; i < FIELD_ELEMENTS_PER_BLOB; i++) {
            uint256 fieldElement = uint256(blob[i]);
            require(
                fieldElement < BLS_MODULUS,
                "Field element must be less than BLS modulus"
            );
            polynomial[i] = fieldElement;
        }

        return polynomial;
    }
}
