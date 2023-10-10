package exporter

// import (
// 	"bytes"
// 	"context"
// 	"fmt"

// 	"go.opentelemetry.io/collector/pdata/plog"
// 	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// const apiEndpointURL = "grpc0.0.0.0:4320"

// func pushLogs(
// 	ctx context.Context,
// 	td plog.Logs,
// ) (err error) {
// 	jsonMarshaler := &plog.JSONMarshaler{}
// 	marshaledLog, err := jsonMarshaler.MarshalLogs(td)

// 	fmt.Println(bytes.NewBuffer(marshaledLog))
// 	// do not make api call as no logs are present
// 	if td.LogRecordCount() == 0 {
// 		return
// 	}

// 	req := plogotlp.NewExportRequestFromLogs(td)
// 	fmt.Println(req)

// 	// resp, respErr := e.metricExporter.Export(e.enhanceContext(ctx), req, e.callOptions...)
// 	// if err := processError(respErr); err != nil {
// 	// 	return err
// 	// }
// 	// partialSuccess := resp.PartialSuccess()
// 	// if !(partialSuccess.ErrorMessage() == "" && partialSuccess.RejectedDataPoints() == 0) {
// 	// 	return consumererror.NewPermanent(fmt.Errorf("OTLP partial success: \"%s\" (%d rejected)", resp.PartialSuccess().ErrorMessage(), resp.PartialSuccess().RejectedDataPoints()))
// 	// }
// 	return nil
// 	// jsonMarshaler := &plog.JSONMarshaler{}
// 	// marshaledLog, err := jsonMarshaler.MarshalLogs(td)

// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// if err := sendLogsToAPI(marshaledLog); err != nil {
// 	// 	fmt.Printf("Error sending logs to API: %v\n", err)
// 	// 	return err
// 	// }

// 	// fmt.Println("Logs sent to API successfully")
// 	// return nil
// }

// // func sendLogsToAPI(logData []byte) error {
// // 	// Create an HTTP request
// // 	req, err := http.NewRequest("POST", apiEndpointURL, bytes.NewBuffer(logData))
// // 	if err != nil {
// // 		return err
// // 	}
// // 	req.Header.Set("Content-Type", "application/json")
// // 	// Create an HTTP client
// // 	client := &http.Client{}

// // 	// Send the request
// // 	resp, err := client.Do(req)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	defer resp.Body.Close()

// // 	// Check the response and handle any errors

// // 	if resp.StatusCode != http.StatusOK {
// // 		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
// // 	}

// // 	return nil
// // }

// // func pushLogs(
// // 	ctx context.Context,
// // 	td plog.Logs,
// // ) (err error) {
// // 	fmt.Println("within exporter")
// // 	// do not make api call as no logs are present
// // 	if td.LogRecordCount() == 0 {
// // 		return
// // 	}
// // 	jsonMarshaler := &plog.JSONMarshaler{}
// // 	marshaledLog, err := jsonMarshaler.MarshalLogs(td)

// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	if err := sendLogsToAPI(marshaledLog); err != nil {
// // 		fmt.Printf("Error sending logs to API: %v\n", err)
// // 		return err
// // 	}

// // 	fmt.Println("Logs sent to API successfully")
// // 	return nil
// // }

// // func sendLogsToAPI(logData []byte) error {
// // 	// Create an HTTP request
// // 	req, err := http.NewRequest("POST", apiEndpointURL, bytes.NewBuffer(logData))
// // 	if err != nil {
// // 		return err
// // 	}
// // 	req.Header.Set("Content-Type", "application/json")
// // 	// Create an HTTP client
// // 	client := &http.Client{}

// // 	// Send the request
// // 	resp, err := client.Do(req)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	defer resp.Body.Close()

// // 	// Check the response and handle any errors

// // 	if resp.StatusCode != http.StatusOK {
// // 		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
// // 	}

// // 	return nil
// // }

// func processError(err error) error {
// 	if err == nil {
// 		// Request is successful, we are done.
// 		return nil
// 	}

// 	// We have an error, check gRPC status code.
// 	st := status.Convert(err)
// 	if st.Code() == codes.OK {
// 		// Not really an error, still success.
// 		return nil
// 	}

// 	// Now, this is this a real error.

// 	// retryInfo := getRetryInfo(st)

// 	// if !shouldRetry(st.Code(), retryInfo) {
// 	// 	// It is not a retryable error, we should not retry.
// 	// 	return consumererror.NewPermanent(err)
// 	// }

// 	// // Check if server returned throttling information.
// 	// throttleDuration := getThrottleDuration(retryInfo)
// 	// if throttleDuration != 0 {
// 	// 	// We are throttled. Wait before retrying as requested by the server.
// 	// 	return exporterhelper.NewThrottleRetry(err, throttleDuration)
// 	// }

// 	// Need to retry.

// 	return err
// }

// // func newExporterTest(cfg component.Config, set exporter.CreateSettings) (*baseExporter, error) {
// // 	oCfg := cfg.(*config)

// // 	if oCfg.Endpoint == "" {
// // 		return nil, errors.New("OTLP exporter config requires an Endpoint")
// // 	}

// // 	userAgent := fmt.Sprintf("%s/%s (%s/%s)",
// // 		set.BuildInfo.Description, set.BuildInfo.Version, runtime.GOOS, runtime.GOARCH)

// // 	return &baseExporter{config: oCfg, settings: set.TelemetrySettings, userAgent: userAgent}, nil
// // }
