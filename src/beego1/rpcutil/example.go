package main

import (
	rpc "beego1/rpcutil"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	var server = rpc.NewJsonRpcServer()

	server.Register(new(Arith))
	server.HandleHTTP("/test/")

	go func() {
		var err = http.ListenAndServe("0.0.0.0:12345", nil)

		if err != nil {
			log.Fatal("Serve Http:", err)
		}
	}()

	time.Sleep(time.Second)

	var client, err = rpc.DialJsonRpc("tcp", "127.0.0.1:12345", "/test/")

	if err != nil {
		log.Fatal("Dialing:", err)
	}

	var args = &Args{7, 8}
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal("Call:", err)
	}

	log.Printf("Result: %d * %d = %d", args.A, args.B, reply)

	fmt.Scanln()
}
