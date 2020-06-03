package server
//定义节点注册
import (
    "net/http"
    "blockchain-go/node"
)

type nodesRegisteredResponse struct {
    Message    string `json:"message"`
    TotalNodes int    `json:"total_nodes"`
}

func newNodesRegisteredResponse(nodes []*node.Node) *nodesRegisteredResponse {
    return &nodesRegisteredResponse{
        Message:    "新节点注册",
        TotalNodes: len(nodes),
    }
}

func registerNode(writer http.ResponseWriter, request *http.Request) {
    var nodes []node.Node
    if !requestBodyToPayload(writer, request, &nodes) {
        return
    }

    for _, n := range nodes {
        blockchain.RegisterNode(n)
    }

    writeJsonResponse(writer, newNodesRegisteredResponse(blockchain.GetNodes()), http.StatusOK)
}

func getNodes(writer http.ResponseWriter, request *http.Request) {
    writeJsonResponse(writer, blockchain.GetNodes(), http.StatusOK)
}
