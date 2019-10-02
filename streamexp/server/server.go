package main

import (
	"context"
	"log"
	"net"

	"github.com/sathishkumar64/grpc_pack/streamexp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadMyStoryServiceServer as interface
type ReadMyStoryServiceServer struct {
}

//GetStory is used for get story service
func (r ReadMyStoryServiceServer) GetStory(ctx context.Context, req *streamexp.RequestReadStory) (*streamexp.ReadStoryVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStory not implemented")
}

/*func (r ReadMyStoryServiceServer) ListOfStory(req *streamexp.RequestReadStory, srv streamexp.ReadMyStoryService_ListOfStoryServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOfStory not implemented")
}
func (r ReadMyStoryServiceServer) StoryStatus(srv streamexp.ReadMyStoryService_StoryStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StoryStatus not implemented")
}
func (r ReadMyStoryServiceServer) FullStory(srv streamexp.ReadMyStoryService_FullStoryServer) error {
	return status.Errorf(codes.Unimplemented, "method FullStory not implemented")
}
*/
func main() {

	var obj ReadMyStoryServiceServer

	srv := grpc.NewServer()

	streamexp.RegisterReadMyStoryServiceServer(srv, obj)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 :%v", err)
	}
	log.Fatal(srv.Serve(lis))

}
