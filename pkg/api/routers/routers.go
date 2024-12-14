package routers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/guidewire/fern-reporter/config"
	"github.com/guidewire/fern-reporter/fernreporter_pb"
	"github.com/guidewire/fern-reporter/pkg/api/handlers"
	"github.com/guidewire/fern-reporter/pkg/auth"
	"github.com/guidewire/fern-reporter/pkg/db"

	"github.com/gin-gonic/gin"
)

var (
	testRun *gin.RouterGroup
)

func RegisterRouters(router *gin.Engine) {
	handler := handlers.NewHandler(db.GetDb())

	authEnabled := config.GetAuth().Enabled

	conn, err := grpc.Dial("grpc-server:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpcClient := fernreporter_pb.NewFernReporterServiceClient(conn)

	var api *gin.RouterGroup
	if authEnabled {
		api = router.Group("/api", auth.ScopeMiddleware())
	} else {
		api = router.Group("/api")
	}

	api.Use()
	{
		testRun = api.Group("/testrun/")
		testRun.GET("/", handler.GetTestRunAll)
		testRun.GET("/:id", handler.GetTestRunByID)
		//	func(c *gin.Context) {
		//	GetTestRunByIDHandler(c, grpcClient)
		//})
		testRun.POST("/", handler.CreateTestRun)
		testRun.PUT("/:id", handler.UpdateTestRun)
		testRun.DELETE("/:id", handler.DeleteTestRun)

		testReport := api.Group("/reports")
		testReport.GET("/projects/", handler.GetProjectAll)
		testReport.GET("/summary/:name/", handler.GetTestSummary)
		testReport.GET("/testruns/", handler.ReportTestRunAll)
		testReport.GET("/testruns/:id/", handler.ReportTestRunById)
	}

	var reports *gin.RouterGroup
	if authEnabled {
		reports = router.Group("/reports/testruns", auth.ScopeMiddleware())
	} else {
		reports = router.Group("/reports/testruns")
	}

	reports.Use()
	{
		reports.GET("/", handler.ReportTestRunAllHTML)
		reports.GET("/:id", handler.ReportTestRunByIdHTML)
	}

	var ping *gin.RouterGroup
	if authEnabled {
		ping = router.Group("/ping", auth.ScopeMiddleware())
	} else {
		ping = router.Group("/ping")
	}

	ping.Use()
	{
		ping.GET("/", func(c *gin.Context) {
			PingHandler(c, grpcClient)
		})
	}
	insights := router.Group("/insights")
	{
		insights.GET("/:name", handler.ReportTestInsights)
	}
}

// PingHandler handles HTTP requests and uses the gRPC client to make a gRPC call
func PingHandler(c *gin.Context, grpcClient fernreporter_pb.FernReporterServiceClient) {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the gRPC call
	response, err := grpcClient.Ping(ctx, &fernreporter_pb.PingRequest{Message: "Ping"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond to the HTTP request with the gRPC response
	c.JSON(http.StatusOK, gin.H{"response": response.Message})
}

// GetTestRunByIDHandler handles HTTP requests and uses the gRPC client to make a gRPC call
func GetTestRunByIDHandler(c *gin.Context, grpcClient fernreporter_pb.FernReporterServiceClient) {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the gRPC call
	testRunId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response, err := grpcClient.GetTestRunByID(ctx, &fernreporter_pb.GetTestRunByIDRequest{
		Id: testRunId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond to the HTTP request with the gRPC response
	c.JSON(http.StatusOK, gin.H{"response": response.TestRun})
}
