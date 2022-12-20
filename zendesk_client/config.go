package zendesk_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	SubDomain string `yaml:"subdomain,omitempty" mapstructure:"subdomain"`
	Email     string `yaml:"email,omitempty" mapstructure:"email"`
	Token     string `yaml:"token,omitempty" mapstructure:"token"`
}
