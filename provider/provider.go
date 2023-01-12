package provider

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-zendesk/zendesk_client"
	"github.com/spf13/viper"
	"os"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "zendesk",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var zendeskConfig zendesk_client.Configs

				err := config.Unmarshal(&zendeskConfig.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(zendeskConfig.Providers) == 0 {
					zendeskConfig.Providers = append(zendeskConfig.Providers, zendesk_client.Config{})
				}

				if zendeskConfig.Providers[0].SubDomain == "" {
					zendeskConfig.Providers[0].SubDomain = os.Getenv("ZENDESK_SUBDOMAIN")
				}

				if zendeskConfig.Providers[0].SubDomain == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing SubDomain in configuration")
				}

				if zendeskConfig.Providers[0].Email == "" {
					zendeskConfig.Providers[0].Email = os.Getenv("ZENDESK_EMAIL")
				}

				if zendeskConfig.Providers[0].Email == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing email in configuration")
				}

				if zendeskConfig.Providers[0].Token == "" {
					zendeskConfig.Providers[0].Token = os.Getenv("ZENDESK_TOKEN")
				}

				if zendeskConfig.Providers[0].Token == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing token in configuration")
				}

				clients, err := zendesk_client.NewClients(zendeskConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# subdomain: "<YOUR_SUBDOMAIN>"
# email: "<YOUR_EMAIL>"
# token: "<YOUR_ACCESS_TOKEN>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var client_config zendesk_client.Configs
				err := config.Unmarshal(&client_config.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
