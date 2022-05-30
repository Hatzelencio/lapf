package provider

import (
	"errors"
	"fmt"
	inputs "overlapping-finder/structs"
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

func NewCloudProvider(in *inputs.Ipv4Command, region string) (ICloudProvider, error) {
	if strings.ToLower(in.ProviderName) == "aws" {
		return newAWSClient(in.ProviderProfile, region), nil
	}
	return nil, errors.New(fmt.Sprintf("cloud provider: %v Does not exist."))
}
