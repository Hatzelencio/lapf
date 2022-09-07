package inputs

type Ipv4Command struct {
	ProviderName    string `validate:"required,provider"`
	ProviderProfile string
	OutputFormat    string   `validate:"required,output"`
	Regions         []string `validate:"required"`
	Arguments       []string `validate:"required,cidrv4,min=1"`
}

type EnsureCIDRv4Command struct {
	ShowIPList   bool
	OutputFormat string   `validate:"required,output"`
	Arguments    []string `validate:"required,cidrv4,min=1"`
}
