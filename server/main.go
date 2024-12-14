// server.go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/guidewire/fern-reporter/fernreporter_pb"
	"github.com/guidewire/fern-reporter/pkg/models"
)

type server struct {
	fernreporter_pb.UnimplementedFernReporterServiceServer
	db *gorm.DB
}

//type servertestbyid struct {
//	gtid.UnimplementedTestRunServiceServer
//	db *gorm.DB
//}

func (s *server) Ping(ctx context.Context, in *fernreporter_pb.PingRequest) (*fernreporter_pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &fernreporter_pb.PingResponse{Message: "Pong: " + in.GetMessage()}, nil
}

func (s *server) GetTestRunByID(ctx context.Context, req *fernreporter_pb.GetTestRunByIDRequest) (*fernreporter_pb.GetTestRunByIDResponse, error) {
	var testRun models.TestRun
	id := req.GetId()
	result := s.db.Where("id = ?", id).First(&testRun)
	if result.Error != nil {
		return nil, result.Error
	}

	response := &fernreporter_pb.GetTestRunByIDResponse{
		TestRun: &fernreporter_pb.TestRun{
			Id: testRun.ID,

			// Add other fields here
		},
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fernreporter_pb.RegisterFernReporterServiceServer(s, &server{db: db})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
