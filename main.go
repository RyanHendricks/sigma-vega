package main

import (
	"log"
	"runtime"
	"sigma-vega/cmd"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
