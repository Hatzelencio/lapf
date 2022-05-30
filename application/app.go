package application

import (
	"github.com/briandowns/spinner"
	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v2"
	"log"
	"overlapping-finder/cloud"
	inputs "overlapping-finder/structs"
	"sync"
	"time"
)

var validate *validator.Validate

const (
	cliRegionName      = "region"
	cliProviderName    = "provider"
	cliProviderProfile = "profile"
	cliOutputFormat    = "output"
)

func newSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Retrieving CloudNetwork from regions"

	return s
}

func retrieveAllNetworkFromRegions(input *inputs.Ipv4Command) ([]provider.CloudNetwork, error) {
	var wg sync.WaitGroup
	var ch = make(chan *ResultDescribeAllNetwork)
	var cloudNetworks []provider.CloudNetwork

	for _, region := range input.Regions {
		wg.Add(1)
		go func(region string) {
			defer wg.Done()
			svc, err := provider.NewCloudProvider(input, region)
			if err != nil {
				log.Fatalf("%v", err)
			}
			result, err := svc.RetrieveVpc()
			ch <- &ResultDescribeAllNetwork{
				Networks: result,
				Region:   region,
				Err:      err,
			}
		}(region)
	}

	for range input.Regions {
		resultDescribeAllNetworks := <-ch
		if err := resultDescribeAllNetworks.Err; err != nil {
			return []provider.CloudNetwork{}, err
		}
		cloudNetworks = append(cloudNetworks, resultDescribeAllNetworks.Networks...)
	}

	wg.Wait()

	return cloudNetworks, nil
}

func newOverlappingFinder(c *cli.Context) error {
	input := &inputs.Ipv4Command{
		ProviderName:    c.String(cliProviderName),
		ProviderProfile: c.String(cliProviderProfile),
		OutputFormat:    c.String(cliOutputFormat),
		Regions:         c.StringSlice(cliRegionName),
		Arguments:       c.Args().Slice(),
	}

	if err := newInputValidate(input); err != nil {
		return err
	}

	s := newSpinner()
	s.Start()

	networks, err := retrieveAllNetworkFromRegions(input)

	if err != nil {
		log.Fatalf("Something wrong happened: %v", err)
	}
	s.Stop()

	s.Suffix = " Ensuring CloudNetwork CIDR Blocks"
	s.Start()

	results, err := ensureCIDRBlock(networks, input.Arguments)
	if err != nil {
		log.Fatalf("Something wrong happened: %v", err)
	}
	s.Stop()

	PrintResults(input.OutputFormat, results)

	return nil
}
