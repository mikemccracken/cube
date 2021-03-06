package cli

import (
	"flag"
	"os"

	"github.com/anuvu/cube/config"
	"github.com/anuvu/cube/service"
)

// Cli wraps flags.
type Cli struct {
	Flags *flag.FlagSet
}

// New returns new instance of Cli.
func New() *Cli {
	return &Cli{flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
}

// Configure parses the flags.
func (c *Cli) Configure(ctx service.Context, store config.Store) error {
	return c.Flags.Parse(os.Args[1:])
}
