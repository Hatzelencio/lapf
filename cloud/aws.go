package provider

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
)

type ClientAWS struct {
	ICloudProvider
	cli *ec2.Client
}

func newAWSClient(region string) *ClientAWS {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &ClientAWS{cli: ec2.NewFromConfig(cfg)}
}

func (c *ClientAWS) RetrieveVpc() ([]CloudNetwork, error) {
	var networks []CloudNetwork

	resp, err := c.cli.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{})
	if err != nil {
		return []CloudNetwork{}, err
	}

	for _, vpc := range resp.Vpcs {
		networks = append(networks, CloudNetwork{
			Name:         *vpc.VpcId,
			ProviderName: "aws",
			CidrBlock:    *vpc.CidrBlock,
		})
	}
	return networks, nil
}
