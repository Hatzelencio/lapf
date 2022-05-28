package provider

import (
	"errors"
	"fmt"
	"strings"
)

type ICloudProvider interface {
	RetrieveVpc() ([]CloudNetwork, error)
}

type CloudNetwork struct {
	Name         string
	ProviderName string
	CidrBlock    string
}

func NewCloudProvider(providerName, region string) (ICloudProvider, error) {
	if strings.ToLower(providerName) == "aws" {
		return newAWSClient(region), nil
	}
	return nil, errors.New(fmt.Sprintf("cloud provider: %v Does not exist."))
}
