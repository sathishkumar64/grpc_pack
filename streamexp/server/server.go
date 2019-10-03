package main

import (
	"context"
	"encoding/json"
	"github.com/sathishkumar64/grpc_pack/streamexp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
)

// ReadMyStoryServiceServer as interface
type ReadServer struct {
}
func generateStub() []byte{
	var exampleData = []byte(`[{  
        "Course": "HTML",
        "TutorName": "BABU",
 		"Status": true
    },
 	{  
         "Course": "JAVA",
        "TutorName": "BALU",
 		"Status": true
    },
	{
         "Course": "MongoDB",
        "TutorName": "Kumar",
 		"Status": true
    },
	{
        "Course": "GO",
        "TutorName": "Sathish",
 		"Status": true
    },
	{
		"Course": "Graphql",
        "TutorName": "Sathish",
 		"Status": true
    },
	{
        "Course": "SpringBoot",
        "TutorName": "Arun",
 		"Status": true
    },
	{
         "Course": "Mule",
        "TutorName": "Basav",
 		"Status": true
    },
	{
       "Course": "Jmeter",
        "TutorName": "Dhamu",
 		"Status": true
    },
	{
        "Course": "JSP",
        "TutorName": "Manish",
 		"Status": true
    }]`)
	return exampleData
}



//GetStory is used for get story service
func (r ReadServer) GetStory(ctx context.Context, req *streamexp.RequestReadStory) (*streamexp.ReadStoryVO, error) {
	log.Println("Getting inside.................",req.GetCourse())
	switch req.GetCourse() {
	case "GO":
		aStudent := &streamexp.ReadStoryVO{
			Course:   "GO",
			TutorName: "Sathish",
			Status: true,
		}
		return aStudent, nil
	case "Java":
		aStudent := &streamexp.ReadStoryVO{
			Course:   "Java",
			TutorName: "BALU",
			Status: true,
		}
		return aStudent, nil
	default :
		aStudent := &streamexp.ReadStoryVO{
			Course:   "HTML",
			TutorName: "BABU",
			Status: true,
		}
		return aStudent, nil
	}
}

func (r ReadServer) ListOfStory(req *streamexp.RequestReadStory, srv streamexp.ReadMyStoryService_ListOfStoryServer) error {
	log.Println("Getting inside of ListOfStory.................",req.GetCourse())
	var savedFeatures []*streamexp.ReadStoryVO
	var data []byte
		data = generateStub()

	if err := json.Unmarshal(data, &savedFeatures); err != nil {
		log.Fatalf("Failed to load default Course Data: %v", err)
	}
	for _, course := range savedFeatures {
			if err := srv.Send(course); err != nil {
				return err
			}
	}
	return nil




}

func (r ReadServer) StoryStatus(srv streamexp.ReadMyStoryService_StoryStatusServer) error {
	log.Println("Getting inside of StoryStatus.................")
	var savedFeatures []*streamexp.ReadStoryVO
	var data []byte
	data = generateStub()

	if err := json.Unmarshal(data, &savedFeatures); err != nil {
		log.Fatalf("Failed to load default Course Data: %v", err)
	}
	/*var (
			crsstatus bool
			tutorName string
		)*/


	for{

		courseObj, err := srv.Recv()
		if courseObj.Course != "" {
			/*for _, course := range savedFeatures {
				if course.Course == courseObj.Course {
					fmt.Printf("data.................%v %v ", course.Status, course.TutorName)
						crsstatus = course.Status
						tutorName = course.TutorName

				}
			}*/

			if err == io.EOF {
				log.Println("courseObj.Course................", courseObj.Course)
				return srv.SendAndClose(&streamexp.ReadStoryVO{Course: courseObj.Course, TutorName: "Test", Status: false,})
			}

			if err != nil {
				log.Fatalf("Failed reciving data %v", err)
				return err
			}
		}
	}
	return nil
}

func (r ReadServer) FullStory(srv streamexp.ReadMyStoryService_FullStoryServer) error {
	return status.Errorf(codes.Unimplemented, "method FullStory not implemented")
}

func main() {
	var obj ReadServer
	srv := grpc.NewServer()
	streamexp.RegisterReadMyStoryServiceServer(srv, obj)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 :%v", err)
	}
	log.Fatal(srv.Serve(lis))
}
