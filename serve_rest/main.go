package serve_rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/answer-map/answer-protobuf"
	"github.com/answer-map/answer-service/app"
)

func run(config *app.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterAnswerServiceHandlerFromEndpoint(ctx, mux, config.GRPCHTTP.Address(), opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(config.RESTHTTP.Address(), mux)
}

func main() {
	configAsBytes, err := os.ReadFile(`config/local.json`)
	if err != nil {
		log.Fatalf("failed to load config, error = %v", err)
	}

	var config app.Config
	if err := json.Unmarshal(configAsBytes, &config); err != nil {
		log.Fatalf("failed to unmarshal config, error = %v", err)
	}

	if err := run(&config); err != nil {
		log.Fatal(err)
	}
}
