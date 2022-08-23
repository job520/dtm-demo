package main

import (
	"fmt"
	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/lithammer/shortuuid/v3"
)

func main() {
	saga := dtmcli.NewSaga("http://localhost:36789/api/dtmsvr", shortuuid.New()).
		Add("http://localhost:8080/trans_in", "http://localhost:8080/trans_in_compensate", nil).
		Add("http://localhost:8080/trans_out", "http://localhost:8080/trans_out_compensate", nil)
	if err := saga.Submit(); err != nil {
		fmt.Println("submit error:", err)
	}
}
