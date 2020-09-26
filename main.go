package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/maxidelgado/skeleton-go/api"
	"github.com/maxidelgado/skeleton-go/config"
	"github.com/maxidelgado/skeleton-go/domain/example"
	"github.com/maxidelgado/skeleton-go/respository/dataaccess"
	"github.com/maxidelgado/toolkit-go/pkg/router"
	"github.com/maxidelgado/toolkit-go/pkg/router/adapter"
)

var release bool

func init() {
	release = os.Getenv("DEPLOY_MODE") == "release"
}

func main() {
	// setup router
	r := router.New(config.Get().Router)

	// setup repositories
	repo := dataaccess.New()

	// setup services
	svc := example.New(repo)

	// setup handlers
	exampleHandler := api.NewExampleHandler(svc)

	// register handlers
	r.RegisterHandlers(
		exampleHandler,
	)

	switch {
	case release:
		listenLambda(r.Engine())
	default:
		log.Fatal(r.Engine().Listen(":3000"))
	}
}

func listenLambda(app *fiber.App) {
	adt := adapter.New(app)
	handle := func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return adt.ProxyWithContext(ctx, req)
	}
	lambda.Start(handle)
}
