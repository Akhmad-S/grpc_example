package tutorial

import (
	context "context"
	"grpctutorial/proto-gen/tutorial"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// We define a TutorialService struct that implements the server interface.
type TutorialService struct {
	tutorial.UnimplementedTutorialServer
}

// We implement the SayHello method of the TutorialService interface.
func (s *TutorialService) SayHello(ctx context.Context, in *tutorial.HelloRequest) (*tutorial.HelloReply, error) {
	return &tutorial.HelloReply{
		Message: "Hello, " + in.GetName()}, nil
}

// We implement the RollDice method of the TutorialService interface.
func (s *TutorialService) RollDice(ctx context.Context, req *tutorial.RollDiceRequest) (*tutorial.RollDiceResponse, error) {
	var res tutorial.RollDiceResponse

	if req.Num < 0{
		return nil, grpc.Errorf(codes.InvalidArgument, "number should be positive")
	}
	
	for i := 0; i < int(req.Num); i++ {
		res.Dice = append(res.Dice, int32(rand.Intn(100)))
	}

	return &res, nil
}
