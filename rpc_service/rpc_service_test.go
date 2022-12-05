package rpc_service

import (
	"os"
	"testing"

	"google.golang.org/appengine/log"
)

func TestRPCService(t *testing.T) {

	t.Run("Should create a new RPC service", func(t *testing.T) {
		var uniswapAddr = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
		var alchemyConnection = "https://eth-mainnet.ws.alchemyapi.io/v2/" + os.Getenv("ALCHEMY_API_KEY")
		_, err := NewRPC(alchemyConnection, uniswapAddr)
		if err != nil {
			log.Errorf("there was an error creating the RPC connection")
		}
	})
}
