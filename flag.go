package aphgrpc

import (
	"github.com/urfave/cli"
)

// NatsFlag returns a cli.Flag slice for using in the command
// line arguments
func NatsFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "nats-host",
			EnvVar: "NATS_SERVICE_HOST",
			Usage:  "nats messaging server host",
		},
		cli.StringFlag{
			Name:   "nats-port",
			EnvVar: "NATS_SERVICE_PORT",
			Usage:  "nats messaging server port",
		},
	}
}
