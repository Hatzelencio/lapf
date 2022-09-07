package provider

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
	"os"
	"strings"
)

type ClientAWS struct {
	ICloudProvider
	ec2 *ec2.Client
	sts *sts.Client
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

	return &ClientAWS{ec2: ec2.NewFromConfig(cfg), sts: sts.NewFromConfig(cfg)}
}

func (c *ClientAWS) RetrieveVpc() ([]CloudNetwork, error) {
	var networks []CloudNetwork

	resp, err := c.ec2.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{})
	if err != nil {
		return []CloudNetwork{}, err
	}

	for _, vpc := range resp.Vpcs {
		var vpcName string

		for _, tag := range vpc.Tags {
			if strings.EqualFold(*tag.Key, "Name") {
				vpcName = *tag.Value
			}
		}
		networks = append(networks, CloudNetwork{
			Id:        *vpc.VpcId,
			Name:      vpcName,
			CidrBlock: *vpc.CidrBlock,
		})
	}
	return networks, nil
}

func (c *ClientAWS) RetrieveAccountInfo() (CloudAccount, error) {
	var account CloudAccount

	resp, err := c.sts.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		return CloudAccount{}, err
	}

	account.Id = *resp.Account
	account.ProviderName = "aws"

	return account, nil
}
