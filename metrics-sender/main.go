package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"log"
	"time"
)

func main() {
	// Create an OTLP Metric Exporter
	ctx := context.Background()
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithEndpoint("localhost:4317"), otlpmetricgrpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create OTLP metric exporter: %v", err)
	}

	// Create a Metric Reader
	metricReader := metric.NewPeriodicReader(exporter)

	// Create a Metric Provider
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metricReader),
	)

	// Set the global Meter Provider
	otel.SetMeterProvider(meterProvider)

	// Get a Meter
	meter := otel.Meter("example-meter")

	// Create a Counter instrument
	counter, err := meter.Int64Counter(
		"example_counter",
	)
	if err != nil {
		log.Fatalf("Failed to create counter: %v", err)
	}

	// Periodically record metrics
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	log.Println("Sending metrics to localhost:4317...")
	for i := 0; i < 50; i++ {
		<-ticker.C
		counter.Add(ctx, 1)
		log.Println("Metric recorded")
	}

	// Shutdown the Metric Provider to flush remaining metrics
	if err := meterProvider.Shutdown(ctx); err != nil {
		log.Printf("Failed to shutdown MeterProvider: %v", err)
	}

	log.Println("Shutdown complete")
}
