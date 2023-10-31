const { ethers } = require('ethers');
const rollup_address = "0xDf70259674Ad261b3d7fe7F604dcAA634500c534";
const BlobRollup_Artifact = require("../artifacts/contracts/BlobRollup.sol/BlobRollup.json");
const Rollup_Artifact = require("./Rollup.json");
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


    // let httpProvider = new ethers.providers.JsonRpcProvider(
    //     "https://rpc.dencun-devnet-9.ethpandaops.io"
    // );

    // let receipt = await httpProvider.getTransactionReceipt("0x3768ee7c4cf2ac937231f377414e97f5b835c8a90f46835c5c8a6445ed47697d");
    // console.log("receipt:" + JSON.stringify(receipt));

    await evaluate_y();

}


async function prove() {
    const blobRollup = new ethers.Contract(
        rollup_address,
        BlobRollup_Artifact.abi,
        createSigner(),
    )

    let pointEvalInput = "0x017cbc58b7f607e607f97b6080ee0786ce62c801fe8cda72835cf37c0e04b64d267762ab802631f2395fde004449c850cf5be841876a9042c0d2eea57b1922f001f466ec2ee1275c6ca9b69759e40d176637080b2dce9593b508d70218b2a19f98f17d0610266bb586865d4f9789d7ffc04cf35123263c482ba2bd86ef01bf9e128a200b02d17d7f2718a994c5276010a8656431abf5c54523f19671a1329d4664d577ab37dafd0fe1ed98bae13797a4dc6333f4ebc19b77793f9d1b7081e266";
    // console.log(ethers.utils.arrayify(pointEvalInput));
    let result = await blobRollup.evaluation(ethers.utils.arrayify(pointEvalInput));
    console.log("result:" + result);
}

async function evaluate_y(){
    let httpProvider = new ethers.providers.JsonRpcProvider(
        "https://rpc.dencun-devnet-9.ethpandaops.io"
    );
    const signer = new ethers.Wallet("780fcdff557b07a8cc6b47345438aa86e3c4317672565b26a7d0944329ff3ce3", httpProvider);
    const blobRollup = new ethers.Contract(
        "0xb2A771c6BD8E6b21DB7E025132a26f21Fd01661b",
        BlobRollup_Artifact.abi,
        signer,
    )

    let result = await blobRollup.storageBatchs(2);
    console.log("result:" + result);

    // let txs_rlp = "0xf8e1f86f80843b9aca00825208943acb110ebd55649be498dcc641947bfd5e7faad589056bc75e2d631000008083019ecda09b2247436104dccb28ac74bebf9723971e7987dbee76d6e191e22694101922c6a004ff9d79f16497a86835a2e1b2835ee863b4575307807c14e1bda2a09e3b4d0ef86e01843b9aca00825208940425266311aa5858625cd399eadbbfab183494f7888ac7230489e800008083019ecea033ba3571cf65812a9ebd5f58ef5aa21c56fcf955e024fc1fee90a4353cdd04cda01e27d2a06f98c13330f32de49468d0a6b63df3c77fc52d52dfc9c2e451d4dfa1";

    // let reslut = await blobRollup.getPointValue(ethers.utils.arrayify(txs_rlp));
    // console.log("result:" + result);

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