package main

import (
	"context"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting SMTP output example")

	service.RunCLI(context.Background())
}
