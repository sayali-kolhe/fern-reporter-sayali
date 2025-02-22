package handlers

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"github.com/guidewire/fern-reporter/fernreporter_pb"
	"github.com/guidewire/fern-reporter/pkg/models"
)

type GrpcHandler struct {
	fernreporter_pb.UnimplementedFernReporterServiceServer
	db *gorm.DB
}

func NewGrpcHandler(db *gorm.DB) *GrpcHandler {
	return &GrpcHandler{db: db}
}

func (g *GrpcHandler) Ping(ctx context.Context, in *fernreporter_pb.PingRequest) (*fernreporter_pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &fernreporter_pb.PingResponse{Message: "Pong: " + in.GetMessage()}, nil
}

func (g *GrpcHandler) GetTestRunByID(ctx context.Context, req *fernreporter_pb.GetTestRunByIDRequest) (*fernreporter_pb.GetTestRunByIDResponse, error) {
	var testRun models.TestRun
	log.Printf("Received GetTestRunByID request for id: %v", req.GetId())
	id := req.GetId()
	result := g.db.Where("id = ?", id).First(&testRun)
	if result.Error != nil {
		log.Printf("Error for GetTestRunByID request for id: %v, Err: %v", req.GetId(), result.Error.Error())
		return nil, result.Error
	}

	response := &fernreporter_pb.GetTestRunByIDResponse{
		TestRun: &fernreporter_pb.TestRun{
			Id:              testRun.ID,
			TestProjectName: testRun.TestProjectName,
			TestSeed:        testRun.TestSeed,
			StartTime:       timestamppb.New(testRun.StartTime),
			EndTime:         timestamppb.New(testRun.EndTime),
			//SuiteRuns: testRun.SuiteRuns,

			// Add other fields here
		},
	}
	return response, nil
}
