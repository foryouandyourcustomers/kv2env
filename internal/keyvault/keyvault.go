package keyvault

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/foryouandyourcustomers/kv2env/internal/config"
	log "github.com/sirupsen/logrus"
	"path"

	"github.com/foryouandyourcustomers/kv2env/internal/secrets"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
)

var (
	authorizer autorest.Authorizer
)

// initialize keyvault authorizer
func init() {
	log.SetLevel(config.Cfg.LogLevel)
	// first try to get authorizer from cli
	var err error
	log.Debug("Try to get azure login via cli / identity")
	authorizer, err = kvauth.NewAuthorizerFromCLI()
	if err != nil {
		log.Debug("Unable to get authorizer from cli, trying from env")
		authorizer, err = kvauth.NewAuthorizerFromEnvironment()
		if err != nil {
			log.Fatal("Unable to login to azure")
		}
	}
}

// GetKeyvaultEntries - return all secrets in the given keyvault
func GetKeyvaultEntries(kv string) (secrets.SecretList, error) {
	c := keyvault.New()
	c.Authorizer = authorizer

	siter, err := c.GetSecretsComplete(context.Background(), "https://"+kv+".vault.azure.net", nil)
	if err != nil {
		log.Fatalf("unable to get list of secrets: %v\n", err)
	}

	var s secrets.SecretList
	for siter.NotDone() {
		i := siter.Value()

		key := path.Base(*i.ID)
		b, err := c.GetSecret(context.Background(), "https://"+kv+".vault.azure.net", key, "")
		if err != nil {
			log.Fatalf("unable to retrieve secrets value: %v\n", err)
		}
		s.Secrets = append(s.Secrets, secrets.Secret{Key: key, Value: *b.Value})
		siter.Next()
	}
	return s, nil
}
