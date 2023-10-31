package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morphism-labs/morphism-bindings/bindings"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
)

// get 100 blocks and traces from local node and write to file for test use
func main() {
	// filepath update
	// bPath := blockFilePath
	// tPath := traceFilePath

	l2Client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Crit("query L2 node error", "err", err)
	}
	// require.NoError("", err)
	blocks := []*types.Block{}
	traces := []*TraceInfo{}
	for i := blockStart; i <= blockEnd; i++ {
		block, err := l2Client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			log.Crit("query L2 block error", "err", err.Error())
		}
		blocks = append(blocks, block)
		trace, _ := l2Client.GetBlockTraceByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			log.Crit("query L2 block error", "err", err.Error())
		}
		traces = append(traces, &TraceInfo{
			Number: trace.Header.Number,
			Bls: &bindings.RollupBatchSignature{
				Signers:   trace.Header.BLSData.BLSSigners,
				Signature: trace.Header.BLSData.BLSSignature,
			},
			WithdrawTrieRoot: trace.WithdrawTrieRoot,
		})
	}

	var txs []*types.Transaction
	for _, blockInfo := range blocks {
		txs = append(txs, blockInfo.Transactions()...)
	}
	fmt.Println("txs.len: ", len(txs))

	encodeTxsAndWriteToFile("../resources/block_rlp_hex_vectors_large.json", txs)
	// decodedBlocks := ReadAndDecodeTxsFromFile("block_rlp_hex_vectors.json")

	encodeAndWriteToFile[*TraceInfo]("../resources/trace_rlp_vectors_large.json", traces)

}
