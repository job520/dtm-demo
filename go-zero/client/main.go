package main

import (
	"demo/go-zero/rpc/types/test"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
	"github.com/lithammer/shortuuid/v3"
)

func main() {
	req := &test.Req{}
	saga := dtmgrpc.NewSagaGrpc("etcd://localhost:2379/dtmservice", shortuuid.New()).
		Add("localhost:8080/test.Test/TransIn", "localhost:8080/test.Test/TransInCompensate", req).
		Add("localhost:8080/test.Test/TransOut", "localhost:8080/test.Test/TransOutCompensate", req)
	if err := saga.Submit(); err != nil {
		fmt.Println("submit error:", err)
	}
}
