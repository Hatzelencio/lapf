package inputs

type Ipv4Command struct {
	ProviderName    string   `validate:"required,provider"`
	ProviderProfile string   `validate:"required"`
	OutputFormat    string   `validate:"required,output"`
	Regions         []string `validate:"required"`
	Arguments       []string `validate:"required,cidrv4"`
}
