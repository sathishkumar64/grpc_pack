package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"github.com/sathishkumar64/grpc_pack/overhttpexp"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

var logger *zap.Logger

// InitLogger as logger
func InitLogger() {
	logger, _ = zap.NewProduction()
}

// NotesServer as interface
type NotesServer struct {
}

//PostStory as service
func (n NotesServer) PostStory(ctx context.Context, req *overhttpexp.RequestNotes) (*overhttpexp.NotesVO, error) {
	logger.Info("Getting inside...................",
		zap.String("Request input", req.NotesInfo))
	//log.Println("Getting inside.................", req.NotesInfo)
	notes := &overhttpexp.NotesVO{
		Desc:   req.NotesInfo,
		Id:     rand.Int31(),
		Status: true,
	}
	return notes, nil
}

func main() {
	InitLogger()
	defer logger.Sync()
	var obj NotesServer
	srv := grpc.NewServer()
	overhttpexp.RegisterNotesServiceServer(srv, obj)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		logger.Error("Error fetching url..", zap.Error(err))
		//log.Fatalf("Could not listen to :8888 :%v", err)
	}
	log.Fatal(srv.Serve(lis))
}
