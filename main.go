package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/samluiz/goinit/internal/cli"
)

const banner string = `
#####    #####     ######  ##   ##   ######  ######## 
##   ##  ### ###     ##    ###  ##     ##    ######## 
##       ##   ##     ##    #### ##     ##       ##    
## ####  ##   ##     ##    #######     ##       ##    
##   ##  ##   ##     ##    ## ####     ##       ##    
##   ##  ### ###     ##    ##  ###     ##       ##    
 #####    #####    ######  ##   ##   ######     ##`

var displayVersion bool

func GetVersionInfo() versionInfo {
	if len(goinitversion) != 0 && len(goversion) != 0 {
		return versionInfo{
			goinitversion: goinitversion,
			goversion:     goversion,
		}
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		return versionInfo{
			goinitversion: info.Main.Version,
			goversion:     runtime.Version(),
		}
	}
	return versionInfo{
		goinitversion: "(unknown)",
		goversion:     runtime.Version(),
	}
}

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

func init() {
	parseFlag(os.Args)
	flag.Parse()
}

func main() {
	versionInfo := GetVersionInfo()
	fmt.Print(banner)
	if displayVersion {
		fmt.Printf("\nversion: %s\ngo Version: %s\n", versionInfo.goinitversion, versionInfo.goversion)
		return
	}
	cli.Run(os.Stdin)
}
