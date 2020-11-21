package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-server/pkg/api"
	"log"
	"strconv"
)

type MkShWalkerPathElement struct {
	// TODO: It might be worth defining negative indexes to correspond
	// to the fields "Cond", "Action", "Else", etc.
	Index   int
	Element interface{}
}

func main() {
	a := 333
	b := 31232
	const pi = 3.14321123

	flag.Parse()

	fmt.Println(a, b, pi)

	if flag.NArg() < 2 {
		log.Fatal("Not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
