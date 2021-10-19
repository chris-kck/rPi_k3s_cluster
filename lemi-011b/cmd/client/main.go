package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"
	"github.com/sss-eda/lemi-011b/pkg/infrastructure/rest"
	"github.com/sss-eda/lemi-011b/pkg/infrastructure/serial"

	tarm "github.com/tarm/serial"
)

func main() {
	ctx := context.Background()

	restURL := os.Getenv("LEMI011B_CLIENT_REST_URL")
	if restURL == "" {
		log.Fatal("No environment variable for rest url")
	}
	serialName := os.Getenv("LEMI011B_CLIENT_SERIAL_PORT")
	if serialName == "" {
		log.Fatal("No environment variable for serial port name")
	}
	serialBaud := os.Getenv("LEMI011B_CLIENT_SERIAL_BAUD")
	if serialBaud == "" {
		log.Fatal("No environment variable for serial baud")
	}
	serialBaudInt, err := strconv.Atoi(serialBaud)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := rest.NewClient(restURL)
	if err != nil {
		log.Fatal(err)
	}

	service, err := acquisition.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}

	port, err := tarm.OpenPort(&tarm.Config{
		Name: serialName,
		Baud: serialBaudInt,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctrl, err := serial.NewController(1, service)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(ctrl.Run(ctx, port))
}
