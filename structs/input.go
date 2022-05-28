package inputs

type RetrieveNetworkFromRegions struct {
	ProviderName    string   `validate:"required"`
	ProviderProfile string   `validate:"required"`
	OutputFormat    string   `validate:"required"`
	Regions         []string `validate:"required"`
	Arguments       []string `validate:"required"`
}
