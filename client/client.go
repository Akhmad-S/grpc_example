package main

import (
	"context"
	"log"
	"time"

	"grpctutorial/proto-gen/tutorial"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := tutorial.NewTutorialClient(conn)

	for i := 0; i < 3; i++ {
		res, err := c.SayHello(context.Background(), &tutorial.HelloRequest{Name: "John"})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server (service SayHello): %s", res.Message)
		time.Sleep(time.Second * 1)
	}

	res, err := c.RollDice(context.Background(), &tutorial.RollDiceRequest{Num: 7})
	if err != nil {
		log.Fatalf("Error when calling RollDice: %s", err)
	}
	log.Printf("Response from server (service RollDice): %d", res.Dice)
}
