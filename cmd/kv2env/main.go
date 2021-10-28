package main

import (
	"github.com/foryouandyourcustomers/kv2env/internal/cmd"
	"github.com/foryouandyourcustomers/kv2env/internal/config"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func init() {
	// setup logging
	log.SetOutput(os.Stdout)
	log.SetLevel(config.Cfg.LogLevel)
}

func main() {
	log.Debug("check for parameters")
	if len(os.Args[1:]) < 1 {
		log.Fatal("No arguments to execute given")
	}
	commandToExecute, lookErr := exec.LookPath(os.Args[1])
	if lookErr != nil {
		log.Fatalf("unable to to locate command: %v\n", lookErr)
	}

	paramtersForCommand := []string{commandToExecute}
	for _, parameter := range os.Args[2:] {
		paramtersForCommand = append(paramtersForCommand, parameter)
	}

	cmd.Execute(commandToExecute, paramtersForCommand)
}
