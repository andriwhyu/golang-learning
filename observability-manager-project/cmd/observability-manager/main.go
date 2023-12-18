package main

import (
	"andriruslam/internal/config"
	"flag"
	"fmt"
	"os"
)

func main() {
	appConfig := &config.Config{}

	fmt.Println(*appConfig)

	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)
	}
}
