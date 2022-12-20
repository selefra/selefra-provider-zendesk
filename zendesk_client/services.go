package zendesk_client

import (
	"context"

	"github.com/nukosuke/go-zendesk/zendesk"
)

func Connect(ctx context.Context, config *Config) (*zendesk.Client, error) {
	// You can set custom *http.Client here
	client, err := zendesk.NewClient(nil)
	if err != nil {
		return nil, err
	}

	// example.zendesk.com
	err = client.SetSubdomain(config.SubDomain)
	if err != nil {
		return nil, err
	}

	// Authenticate with API token
	client.SetCredential(zendesk.NewAPITokenCredential(config.Email, config.Token))

	return client, nil
}
