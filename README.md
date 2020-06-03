# blockchian-go
A simple implementation of  POW blockchain with GO



```
(1)安装 glide 
   $ go get github.com/Masterminds/glide
   $ go install github.com/Masterminds/glide

(2) 生成可执行文件

$ glide install
$ go build -i -o ./blockchain-go main.go
(3)  运行


$ ./blockchain-go server --port <port>

如：
$ ./blockchain-go server --port 8080
$ ./blockchain-go server --port 8081

请求(postman/curl)：注册一个区块链节点
curl -sSL -X POST -d '["localhost:8081"]' http://localhost:8080/nodes/register


发送一笔交易
$ curl -sSL -X POST -d '{"sender":"Alice","recipient":"Bob","amount":100}' http://localhost:8080/transactions/new

挖矿生成新区块：
$ curl -sSL -X POST http://localhost:8080/mine

查看chain情况
$ curl -sSL -X GET http://localhost:8080/chain

```
