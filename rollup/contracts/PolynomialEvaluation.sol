// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";

library PolynomialEvaluation {
    uint256 constant FIELD_ELEMENTS_PER_BLOB = 4096;
    uint256 constant BLS_MODULUS = 52435875175126190479447740508185965837690552500527637822603658699938581184513;

    function evaluate_polynomial_in_evaluation_form(
        uint256[] memory polynomial,
        uint256 z
    ) internal pure returns (uint256) {
        require(polynomial.length == FIELD_ELEMENTS_PER_BLOB, "Invalid polynomial length");

        uint256 width = polynomial.length;
        uint256 inverse_width = modular_inverse(width, BLS_MODULUS);
        uint256 result = 0;

        for (uint256 i = 0; i < width; i++) {
            uint256 reversed_i = bit_reversal_permutation(i);
            uint256 a = (polynomial[reversed_i] * reversed_i) % BLS_MODULUS;
            uint256 b = (BLS_MODULUS + z - reversed_i) % BLS_MODULUS;
            result += (a * modular_inverse(b, BLS_MODULUS)) % BLS_MODULUS;
        }

        result = (result * ((z**width - 1) / width)) % BLS_MODULUS;
        return result;
    }

    function modular_inverse(uint256 a, uint256 modulus) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }
        uint256 m = modulus - 2;
        uint256 result = 1;
        while (m > 0) {
            if (m % 2 == 1) {
                result = (result * a) % modulus;
            }
            a = (a * a) % modulus;
            m = m / 2;
        }
        return result;
    }

    function bit_reversal_permutation(uint256 i) internal pure returns (uint256) {
        uint256 n = 0;
        for (uint256 j = 0; j < 12; j++) {  // log2(FIELD_ELEMENTS_PER_BLOB) == 12
            n <<= 1;
            n |= i & 1;
            i >>= 1;
        }
        return n;
    }
}
