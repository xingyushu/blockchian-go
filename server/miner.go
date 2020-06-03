package server

import (
    "net/http"
    "blockchain-go/block"
    "blockchain-go/miner"
)

type minedResponse struct {
    Message      string               `json:"message"`
    Index        int                  `json:"index"`
    Transactions []*block.Transaction `json:"transactions"`
    Nonce        int                  `json:"nonce"`
    PreviousHash string               `json:"previous_hash"`
}

func newMinedResponse(blk *block.Block) *minedResponse {
    return &minedResponse{
        Message:      "挖出新区块",
        Index:        blk.GetIndex(),
        Transactions: blk.GetTransactions(),
        Nonce:        blk.GetNonce(),
        PreviousHash: string(blk.GetPreviousHash()),
    }
}

func mine(writer http.ResponseWriter, request *http.Request) {
    newBlock := miner.Mine(blockchain, recipientMinerAddress)

    writeJsonResponse(writer, newMinedResponse(newBlock), http.StatusOK)
}
