package main

import (
	test "demo/grpc/proto"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/lithammer/shortuuid/v3"
)

func main() {
	req := &test.TestRequest{}
	saga := dtmgrpc.NewSagaGrpc("localhost:36790", shortuuid.New()).
		Add("localhost:50052/test.Test/TransIn", "localhost:50052/test.Test/TransInCompensate", req).
		Add("localhost:50052/test.Test/TransOut", "localhost:50052/test.Test/TransOutCompensate", req)
	if err := saga.Submit(); err != nil {
		fmt.Println("submit error:", err)
	}
}
