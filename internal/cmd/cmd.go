package cmd

import (
	"github.com/foryouandyourcustomers/kv2env/internal/config"
	"github.com/foryouandyourcustomers/kv2env/internal/keyvault"
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
)

// Execute executes the root command.
func Execute(cmd string, params []string) {
	log.Debugf("Retrieve all keyvault secrets for keyvault: %s", config.Cfg.KeyVault)
	secrets, err := keyvault.GetKeyvaultEntries(config.Cfg.KeyVault)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Debug("Convert secret names to env vars")
	secrets.Secrets = secrets.TransformToEnvironmentNames()
	if len(config.Cfg.FilterPrefix) > 0 {
		log.Debugf("Filter retrieved secrets for prefix %s", config.Cfg.FilterPrefix)
		secrets.Secrets = secrets.FilterSecretsByPrefix(config.Cfg.FilterPrefix)
		if config.Cfg.PrefixRemove {
			log.Debug("Remove prefix from retrieved secrets")
			secrets.Secrets = secrets.RemoveSecretsPrefix(config.Cfg.FilterPrefix)
		} else if len(config.Cfg.PrefixReplace) > 0 {
			log.Debugf("Replace prefix with %s", config.Cfg.PrefixReplace)
			secrets.Secrets = secrets.ReplaceSecretsPrefix(config.Cfg.FilterPrefix, config.Cfg.PrefixReplace)
		}
	}

	if config.Cfg.PrefixTfVar {
		log.Debug("Prefix secrets with 'TF_VAR_'")
		secrets.Secrets = secrets.PrefixSecretsWithTfVar()
	}

	log.Debug("Prepare environment for command")
	cmdEnv := make([]string, 0)
	if config.Cfg.PassParentEnv {
		log.Debug("Add current environment variables to next processes env")
		cmdEnv = os.Environ()
	}
	for _, s := range secrets.Secrets {
		cmdEnv = append(cmdEnv, s.ToEnv())
	}

	log.Debugf("Execute command %s", cmd)
	execErr := syscall.Exec(cmd, params, cmdEnv)
	if execErr != nil {
		panic(execErr)
	}
}
