package main

import (
	"context"
	"embed"
	"html/template"
	"log"
	"net"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"github.com/guidewire/fern-reporter/fernreporter_pb"
	"github.com/guidewire/fern-reporter/pkg/graph/resolvers"
	"github.com/guidewire/fern-reporter/pkg/models"
	"github.com/guidewire/fern-reporter/pkg/utils"

	"github.com/guidewire/fern-reporter/pkg/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/guidewire/fern-reporter/config"
	"github.com/guidewire/fern-reporter/pkg/api/routers"
	"github.com/guidewire/fern-reporter/pkg/auth"
	"github.com/guidewire/fern-reporter/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed pkg/views/test_runs.html
//go:embed pkg/views/insights.html
var testRunsTemplate embed.FS

type server struct {
	fernreporter_pb.UnimplementedFernReporterServiceServer
	db *gorm.DB
}

func main() {
	initConfig()
	initDb()
	initServer()
}

func initConfig() {
	if _, err := config.LoadConfig(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func initDb() {
	db.Initialize()
}

func initServer() {
	serverConfig := config.GetServer()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	if config.GetAuth().Enabled {
		checkAuthConfig()
		configJWTMiddleware(router)
	} else {
		log.Println("Auth is disabled, JWT Middleware is not configured.")
	}

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	funcMap := template.FuncMap{
		"CalculateDuration": utils.CalculateDuration,
		"FormatDate":        utils.FormatDate,
	}

	templ, err := template.New("").Funcs(funcMap).ParseFS(testRunsTemplate, "pkg/views/test_runs.html", "pkg/views/insights.html")
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}
	router.SetHTMLTemplate(templ)

	// router.LoadHTMLGlob("pkg/views/*")
	routers.RegisterRouters(router)

	router.POST("/query", GraphqlHandler(db.GetDb()))
	router.GET("/", PlaygroundHandler("/query"))

	// Run grpc server
	go initGrpcServer()

	err = router.Run(serverConfig.Port)
	if err != nil {
		log.Fatalf("error starting routes: %v", err)
	}

}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GraphqlHandler(gormdb *gorm.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: gormdb}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func checkAuthConfig() {
	if config.GetAuth().ScopeClaimName == "" {
		log.Fatal("Set SCOPE_CLAIM_NAME environment variable or add a default value in config.yaml")
	}
	if config.GetAuth().JSONWebKeysEndpoint == "" {
		log.Fatal("Set AUTH_JSON_WEB_KEYS_ENDPOINT environment variable or add a default value in config.yaml")
	}
}

func configJWTMiddleware(router *gin.Engine) {
	authConfig := config.GetAuth()
	ctx := context.Background()

	keyFetcher, err := auth.NewDefaultJWKSFetcher(ctx, authConfig.JSONWebKeysEndpoint)
	if err != nil {
		log.Fatalf("Failed to create JWKS fetcher: %v", err)
	}

	jwtValidator := &auth.DefaultJWTValidator{}

	router.Use(auth.JWTMiddleware(authConfig.JSONWebKeysEndpoint, keyFetcher, jwtValidator))
	log.Println("JWT Middleware configured successfully.")
}

func initGrpcServer() {
	grpcPort := ":50051" // Change as needed
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port %s: %v", grpcPort, err)
	}
	grpcServer := grpc.NewServer()
	fernreporter_pb.RegisterFernReporterServiceServer(grpcServer, &server{db: db.GetDb()})
	log.Printf("Starting gRPC server on port %s...", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}

	log.Println("This will never print unless the server shuts down.")

}

func (s *server) Ping(ctx context.Context, in *fernreporter_pb.PingRequest) (*fernreporter_pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &fernreporter_pb.PingResponse{Message: "Pong: " + in.GetMessage()}, nil
}

func (s *server) GetTestRunByID(ctx context.Context, req *fernreporter_pb.GetTestRunByIDRequest) (*fernreporter_pb.GetTestRunByIDResponse, error) {
	var testRun models.TestRun
	log.Printf("Received GetTestRunByID request for id: %v", req.GetId())
	id := req.GetId()
	result := s.db.Where("id = ?", id).First(&testRun)
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
