package main

import (
	"fmt"

	"github.com/perph/perph-agent/cmd/agent"
)

func main() {
	agent.Execute()
	fmt.Println("shutting down")
}
