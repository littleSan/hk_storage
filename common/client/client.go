package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"hk_storage/core/configs"
	"log"
	"time"
)

var ethClient *ethclient.Client

func GetEthClient() *ethclient.Client {
start:
	if ethClient != nil {
		ethClient.Close()
		ethClient = nil
	}
	var err error
	ethClient, err = ethclient.Dial(configs.TomlConfig.Chain.RPCUrl)
	defer ethClient.Close()
	if err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
		goto start
	}
	return ethClient
}

var rpcClient *rpc.Client

func GetRpcClient() *rpc.Client {
start:
	if rpcClient != nil {
		rpcClient.Close()
		rpcClient = nil
	}
	var err error
	rpcClient, err = rpc.Dial(configs.TomlConfig.Chain.RPCUrl)
	defer rpcClient.Close()
	if err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
		goto start
	}
	return rpcClient
}
