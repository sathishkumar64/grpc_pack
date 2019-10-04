package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sathishkumar64/grpc_pack/streamexp"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Server is not reachable :%v", err)
	}
	client := streamexp.NewReadMyStoryServiceClient(con)
	//	err = getStory(context.Background(), client)
	/*	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}*/

	//err = ListOfStory(context.Background(), client)
	/*if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}*/

	err = StoryStatus(context.Background(), client)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func getStory(ctx context.Context, client streamexp.ReadMyStoryServiceClient) error {
	vo, err := client.GetStory(ctx, &streamexp.RequestReadStory{Course: "GO"})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(vo)
	return err
}

//ListOfStory service is used to display list of service
func ListOfStory(ctx context.Context, client streamexp.ReadMyStoryServiceClient) error {
	stream, err := client.ListOfStory(ctx, &streamexp.RequestReadStory{Course: "GO"})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	for {
		course, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListOfStory(_) = _, %v", client, err)
		}
		log.Println(course)
	}
	return err
}

//StoryStatus service is used to display list of service
func StoryStatus(ctx context.Context, client streamexp.ReadMyStoryServiceClient) error {
	stream, err := client.StoryStatus(ctx)
	courseList := []string{"GO", "Java", "HTML", "Mule"}

	for _, crs := range courseList {

		log.Println("CRS::::", crs)
		if err := stream.Send(&streamexp.RequestReadStory{Course: crs}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, crs, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Story Status: %v", reply)
	return err
}

//StoryStatus service is used to display list of service
func FullStory(ctx context.Context, client streamexp.ReadMyStoryServiceClient) error {

}