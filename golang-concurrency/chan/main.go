package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type response struct {
	id  int
	err error
}

func main() {
	rand.Seed(time.Now().Unix())

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)

	chanForResp := make(chan response)
	go RPCCall(ctx, chanForResp)
	resp := <-chanForResp
	fmt.Println(resp.id, resp.err)
}

func RPCCall(ctx context.Context, ch chan<- response) {
	select {
	case <-ctx.Done():
		ch <- response{
			id:  0,
			err: fmt.Errorf("timeout expired"),
		}
	case <-time.After(time.Second * time.Duration(rand.Intn(5))):
		ch <- response{
			id:  rand.Int(),
			err: nil,
		}
	}
}
