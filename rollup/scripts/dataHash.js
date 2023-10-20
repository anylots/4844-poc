const { ethers } = require('ethers');
const rollup_address = "0x13a1F1832D02366042661dA6A14239C90A8126D4";
const BlobRollup_Artifact = require("../artifacts/contracts/BlobRollup.sol/BlobRollup.json");
const fs = require("fs");
require("dotenv").config({ path: ".env" });


/**
 * Update l2 gas price.
 * Use this command to execute: node .\scripts\authManage.js
 */
async function main() {
    // const blobRollup = new ethers.Contract(
    //     rollup_address,
    //     BlobRollup_Artifact.abi,
    //     createSigner(),
    // )

    // // let tx = await blobRollup.populateTransaction.submit();
    // // console.log("tx:" + JSON.stringify(tx));

    // let result = await blobRollup.getDataHash();
    // console.log("result:" + result);


    let httpProvider = new ethers.providers.JsonRpcProvider(
        "https://rpc.dencun-devnet-9.ethpandaops.io"
    );

    let receipt = await httpProvider.getTransactionReceipt("0x3768ee7c4cf2ac937231f377414e97f5b835c8a90f46835c5c8a6445ed47697d");
    console.log("receipt:" + JSON.stringify(receipt));

}

/**
 * Use priveteKey and rpc create Signer for L2.
 * 
 * @returns 
 */
function createSigner() {
    let httpProvider = new ethers.providers.JsonRpcProvider(
        "https://rpc.dencun-devnet-9.ethpandaops.io"
    );
    const signer = new ethers.Wallet("780fcdff557b07a8cc6b47345438aa86e3c4317672565b26a7d0944329ff3ce3", httpProvider);
    return signer;
}


/**
 * Load environment variables 
 * 
 * @param {*} entry 
 * @returns 
 */
function requireEnv(entry) {
    if (process.env[entry]) {
        return process.env[entry]
    } else {
        throw new Error(`${entry} not defined in .env`)
    }
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});