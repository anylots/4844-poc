package main

import (
	"io"
	"math/big"
	"os"

	"github.com/morphism-labs/morphism-bindings/bindings"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rlp"
)

type TraceInfo struct {
	Number           *big.Int
	Bls              *bindings.RollupBatchSignature
	WithdrawTrieRoot common.Hash
}

var (
	blockFilePath = "block.json"
	traceFilePath = "trace.json"
	blockStart    = 9730
	blockEnd      = 10738
	TestData      *TestDataInfo
)

type TestDataInfo struct {
	BlsCnt  int
	BlsMark map[int]bool // index in Blocks -> has bls
	Traces  []*TraceInfo
	Blocks  []*types.Block
}

type BlockOrTrace interface {
	*types.Block | *TraceInfo
}

func encodeTxsAndWriteToFile(path string, txs []*types.Transaction) {
	bs, err := rlp.EncodeToBytes(txs)
	if err != nil {
		log.Crit("encode txs to bytes error", "err", err)
	}
	data := hexutil.Encode(bs)
	outputFile, err := os.Create(path)
	if err != nil {
		log.Crit("save txs ti file error", "err", err)
	}
	outputFile.WriteString(data)
}

func ReadAndDecodeTxsFromFile(path string) []types.Transaction {
	bs := readBsFromFile(path)
	str := string(bs)

	data := hexutil.MustDecode(str)
	blocks := []types.Transaction{}
	err := rlp.DecodeBytes(data, &blocks)
	if err != nil {
		panic(err)
	}
	return blocks
}

func encodeAndWriteToFile[T BlockOrTrace](path string, blocks []T) {
	bs, err := rlp.EncodeToBytes(blocks)
	if err != nil {
		panic(err)
	}
	writeBsToFile(path, bs)
}
func writeBsToFile(path string, bs []byte) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	_, _ = file.Write(bs)
	defer file.Close()
}
func readBsFromFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	bs, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return bs
}
func ReadAndDecodeBlocksFromFile[T BlockOrTrace](path string) []T {
	bs := readBsFromFile(path)
	blocks := []T{}
	err := rlp.DecodeBytes(bs, &blocks)
	if err != nil {
		panic(err)
	}
	return blocks
}

func GetTestData(path string) *TestDataInfo {
	bPath := path + blockFilePath
	tPath := path + traceFilePath
	if TestData != nil {
		return TestData
	}
	blocks := ReadAndDecodeBlocksFromFile[*types.Block](bPath)
	traces := ReadAndDecodeBlocksFromFile[*TraceInfo](tPath)
	info := TestDataInfo{}
	info.Traces = traces
	info.Blocks = blocks
	info.BlsMark = make(map[int]bool)
	for i := 0; i < len(info.Traces); i++ {
		if len(info.Traces[i].Bls.Signature) > 0 {
			info.BlsCnt++
			info.BlsMark[i] = true
		}
	}
	TestData = &info
	return TestData
}
