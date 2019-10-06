package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sathishkumar64/grpc_pack/overhttpexp"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Server is not reachable :%v", err)
	}

	client := overhttpexp.NewNotesServiceClient(con)

	err = PostNotes(context.Background(), client)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//PostNotes as client service
func PostNotes(ctx context.Context, client overhttpexp.NotesServiceClient) error {
	vo, err := client.PostStory(ctx, &overhttpexp.RequestNotes{NotesInfo: "Have to complete GOLang ASAP",
		ContactInfo: "sathishkumar64@gmail.com",
		Priority:    1})
	if err != nil {
		log.Fatalf("%v.PostNotes(_) = _, %v: ", client, err)
	}
	log.Println(vo)
	return err
}
