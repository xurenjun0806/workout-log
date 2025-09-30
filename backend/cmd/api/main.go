package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	rest_exercise "github.com/xurenjun0806/workout-log/backend/interfaces/rest/exercise"
	"github.com/xurenjun0806/workout-log/backend/interfaces/rest/middleware"
	usecase_exercise "github.com/xurenjun0806/workout-log/backend/usecase/exercise"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func main() {
	e := echo.New()
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// TODO: useCaseはいったんダミーで
	rest_exercise.NewExerciseHandler(e, &usecase_exercise.UseCase{})

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address))
}
