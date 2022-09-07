package provider

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
	"os"
)

type ClientAWS struct {
	ICloudProvider
	cli *ec2.Client
}

func newAWSClient(profile, region string) *ClientAWS {
	var cfg aws.Config
	var err error

	if len(profile) == 0 {
		fmt.Fprintln(os.Stderr, "[WARN] AWS_PROFILE variable did not define")
		fmt.Fprintln(os.Stderr, "[WARN] ensure if was set a valid access_keys")
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
			config.WithSharedConfigProfile(profile),
		)
	}

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
