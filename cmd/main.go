package main

import (
	"os"

	"github.com/samluiz/go-bootstrap-cli/internal/cli"
)

func main() {
	cli.Run(os.Stdin)
}
