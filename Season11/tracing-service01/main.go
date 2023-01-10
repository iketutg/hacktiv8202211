package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"log"
	"net/http"
	"time"

	sdktracer "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	service           = "service01"
	environment       = "development"
	id                = 1
	urlTracerProvider = "http://localhost:14268/api/traces"
)

func tracerProv() (*sdktracer.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(urlTracerProvider)))
	if err != nil {
		return nil, err
	}

	provider := sdktracer.NewTracerProvider(
		sdktracer.WithBatcher(exporter),
		sdktracer.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", environment),
			attribute.Int("ID", id),
		)),
	)

	return provider, nil

}

func main() {
	prov, err := tracerProv()

	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(prov)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	//clean shutdown flush opentelementry when exit
	defer func(ctx context.Context) {
		_, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := prov.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	tr := prov.Tracer("Component-Utama")

	ctx, span := tr.Start(ctx, "SPAN-SERVICE-HELLO")
	defer span.End()

	httpHandlerAPI := func(w http.ResponseWriter, r *http.Request) {
		tr := otel.Tracer("API-HANDLER")
		_, span := tr.Start(ctx, "hello")
		span.SetAttributes(attribute.Key("KeyAttrTest").String("Value didalalam http handler"))
		defer span.End()
		fmt.Fprintf(w, "Hello OpenTracing !!!! ")
	}

	otelhandler := otelhttp.NewHandler(http.HandlerFunc(httpHandlerAPI), "Hello")

	http.Handle("/api", otelhandler)

	log.Println("Listen Server localhost :8081 ")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
