package handlers

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
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
		TestRun: convertTestRunToProto(testRun),
	}
	return response, nil
}

func (g *GrpcHandler) GetTestRunAll(ctx context.Context, empty *emptypb.Empty) (*fernreporter_pb.GetTestRunAllResponse, error) {
	var testRuns []models.TestRun
	g.db.Find(&testRuns)

	var protoTestRuns []*fernreporter_pb.TestRun
	for _, testRun := range testRuns {
		protoTestRuns = append(protoTestRuns, convertTestRunToProto(testRun))
	}
	response := &fernreporter_pb.GetTestRunAllResponse{
		TestRuns: protoTestRuns,
	}

	return response, nil
}

// Helper Functions:

// Convert TestRun struct
func convertTestRunToProto(testRun models.TestRun) *fernreporter_pb.TestRun {
	return &fernreporter_pb.TestRun{
		Id:              testRun.ID,
		TestProjectName: testRun.TestProjectName,
		TestSeed:        testRun.TestSeed,
		StartTime:       timestamppb.New(testRun.StartTime),
		EndTime:         timestamppb.New(testRun.EndTime),
		SuiteRuns:       convertSuiteRunsToProto(testRun.SuiteRuns),
	}
}

// Convert a slice of SuiteRun structs
func convertSuiteRunsToProto(suiteRuns []models.SuiteRun) []*fernreporter_pb.SuiteRun {
	var protoSuiteRuns []*fernreporter_pb.SuiteRun
	for _, suiteRun := range suiteRuns {
		protoSuiteRuns = append(protoSuiteRuns, &fernreporter_pb.SuiteRun{
			Id:        suiteRun.ID,
			TestRunId: suiteRun.TestRunID,
			SuiteName: suiteRun.SuiteName,
			StartTime: timestamppb.New(suiteRun.StartTime),
			EndTime:   timestamppb.New(suiteRun.EndTime),
			SpecRuns:  convertSpecRunsToProto(suiteRun.SpecRuns),
		})
	}
	return protoSuiteRuns
}

// Convert a slice of SpecRun structs
func convertSpecRunsToProto(specRuns []models.SpecRun) []*fernreporter_pb.SpecRun {
	var protoSpecRuns []*fernreporter_pb.SpecRun
	for _, specRun := range specRuns {
		protoSpecRuns = append(protoSpecRuns, &fernreporter_pb.SpecRun{
			Id:              specRun.ID,
			SuiteId:         specRun.SuiteID,
			SpecDescription: specRun.SpecDescription,
			Status:          specRun.Status,
			Message:         specRun.Message,
			Tags:            convertTagsToProto(specRun.Tags),
			StartTime:       timestamppb.New(specRun.StartTime),
			EndTime:         timestamppb.New(specRun.EndTime),
		})
	}
	return protoSpecRuns
}

// Convert a slice of Tag structs
func convertTagsToProto(tags []models.Tag) []*fernreporter_pb.Tag {
	var protoTags []*fernreporter_pb.Tag
	for _, tag := range tags {
		protoTags = append(protoTags, &fernreporter_pb.Tag{
			Id:   tag.ID,
			Name: tag.Name,
		})
	}
	return protoTags
}
