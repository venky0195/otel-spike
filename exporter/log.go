package exporter

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"google.golang.org/grpc"
)

type projectData struct {
	projectUid string
	endpoint   string
}

func pushLogs(
	ctx context.Context,
	td plog.Logs,
) error {
	fmt.Println("request within exporter")
	rls := td.ResourceLogs()
	endpoint := ""
	rls.RemoveIf(func(rl plog.ResourceLogs) bool {
		resourceAttibute := rl.Resource().Attributes()
		projectUidFromResourceAttribute, isPresent := resourceAttibute.Get("project_id")

		if !isPresent {
			return true
		} else {
			endpoint = getDataFromRedis(ctx, projectUidFromResourceAttribute.AsString())
			fmt.Println("value of endpoint from reids is", endpoint)
			return false
		}
	})
	if endpoint == "" {
		fmt.Println("No endpoint is set, No logs to forward")
		return nil
	} else {
		err := forwardLogsToGRPC(endpoint, ctx, td)
		return err
	}

}

func forwardLogsToGRPC(endpoint string, ctx context.Context, td plog.Logs) error {
	// Set up a gRPC connection to the server.
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client using the connection.
	logClient := plogotlp.NewGRPCClient(conn)

	// Convert the pdata.Logs to an ExportRequest.
	request := plogotlp.NewExportRequestFromLogs(td)

	// Make a gRPC call to the server.
	_, err = logClient.Export(ctx, request)
	if err != nil {
		log.Printf("gRPC call failed: %v", err)
		return err
	}

	return nil
}

func getDataFromRedis(ctx context.Context, projectUid string) string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	endpoint, _ := rdb.HGet("project", projectUid).Result()

	return endpoint
}
