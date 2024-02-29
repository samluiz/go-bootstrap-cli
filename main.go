package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/samluiz/goinit/internal/cli"
)

var (
	displayVersion bool
	version        string
)

func helpMessage() {
	fmt.Println("Usage: goinit [flags]")
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func parseFlag(args []string) {
	flag.Usage = helpMessage
	flag.BoolVar(&displayVersion, "v", false, "display current version")
	if err := flag.CommandLine.Parse(args); err != nil {
		log.Fatal(err)
	}
}

func getBanner() string {
	banner, err := os.ReadFile("banner.txt")
	if err != nil {
		return "Error reading banner.txt"
	}
	return string(banner)
}

func init() {
	parseFlag(os.Args)
	flag.Parse()
}

func main() {
	banner := getBanner()
	fmt.Println(banner)
	if displayVersion {
		fmt.Println(version)
		return
	}
	cli.Run(os.Stdin)
}
