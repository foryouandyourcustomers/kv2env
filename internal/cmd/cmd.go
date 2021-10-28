package cmd

import (
	"fmt"
	"github.com/foryouandyourcustomers/kv2env/internal/config"
	"github.com/foryouandyourcustomers/kv2env/internal/keyvault"
	log "github.com/sirupsen/logrus"
)

// Execute executes the root command.
func Execute(cmd string) error {
	log.Debugf("Retrieve all keyvault secrets for keyvault: %s", config.Cfg.KeyVault)
	secrets, err := keyvault.GetKeyvaultEntries(config.Cfg.KeyVault)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(secrets)
	//return rootCmd.Execute()
	return nil
}
