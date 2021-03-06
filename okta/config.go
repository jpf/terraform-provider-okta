package okta

import (
	"fmt"

	articulateOkta "github.com/articulate/oktasdk-go/okta"
	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/okta/cache"
)

// Config is a struct containing our provider schema values
// plus the okta client object
type Config struct {
	orgName     string
	domain      string
	apiToken    string
	retryCount  int
	parallelism int

	articulateOktaClient *articulateOkta.Client
	oktaClient           *okta.Client
	supplementClient     *ApiSupplement
}

func (c *Config) loadAndValidate() error {
	articulateClient, err := articulateOkta.NewClientWithDomain(nil, c.orgName, c.domain, c.apiToken)

	// add the Articulate Okta client object to Config
	c.articulateOktaClient = articulateClient

	if err != nil {
		return fmt.Errorf("[ERROR] Error creating Articulate Okta client: %v", err)
	}

	orgUrl := fmt.Sprintf("https://%v.%v", c.orgName, c.domain)

	config := okta.NewConfig().
		WithOrgUrl(orgUrl).
		WithToken(c.apiToken).
		WithCache(false).
		WithBackoff(true).
		WithRetries(int32(c.retryCount))
	client := okta.NewClient(config, nil, nil)
	c.supplementClient = &ApiSupplement{
		requestExecutor: okta.NewRequestExecutor(nil, cache.NewNoOpCache(), config),
	}

	// add the Okta SDK client object to Config
	c.oktaClient = client
	return nil
}
