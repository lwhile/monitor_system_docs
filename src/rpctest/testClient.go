package rpctest

import (
	"fmt"
	"log"
	"net/rpc"
)

func StartClient() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error：", err)
	}
	fmt.Println(args.A, args.B, reply)
}
