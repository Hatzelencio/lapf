package application

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

var cliAcceptedOutputFormat = []string{"text", "json"}
var cliAcceptedCloudProvider = []string{"aws"}

func New() {
	app := &cli.App{
		Name:  "Overlapping Finder",
		Usage: "Overlapping Finder (a.k.a \"lapf\") is a binary-tool made and built info golang to find possible overlapping CIDR Block notation through cloud providers (AWS)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "provider",
				Usage:       "Cloud Provider",
				DefaultText: strings.Join(cliAcceptedCloudProvider, ", "),
				Value:       cliAcceptedCloudProvider[0],
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "Output format",
				DefaultText: strings.Join(cliAcceptedOutputFormat, ", "),
				Value:       "text",
			},
		},
		Commands: commands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "ipv4",
			Description: "Find overlapping CIDR Block over IPV4",
			Usage:       "lapf ipv4 192.168.0.0/24",
			ArgsUsage:   "[optional | --profile <aws>] [target | <192.168.0.0/24> ... <N>]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "profile",
					Usage:   "Env Variable AWS Profile",
					EnvVars: []string{"AWS_PROFILE"},
				},
				&cli.StringSliceFlag{
					Name:    "region",
					Aliases: []string{"r"},
					Value: cli.NewStringSlice(
						"us-east-1", "us-east-2", "us-west-1", "us-west-2"),
				},
			},
			Action: newOverlappingFinder,
		},
		{
			Name:        "ensure",
			Description: "Ensures the IPv4 or CIDR Block is private, following RFC 1918",
			Subcommands: []*cli.Command{
				{
					Name:        "cidr",
					Description: "Ensures the CIDR Block is private, following RFC 1918",
					Usage:       "lapf ensure cidr 192.168.0.0/24",
					ArgsUsage:   "[optional | --show-ip-list] [target | <192.168.0.0/16> ... <N>]",
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "show-ip-list",
							Aliases: []string{"show", "list"},
							Usage:   "Includes the list of IPs out of private context.",
						},
					},
					Action: isCIDRBlockPrivate,
				},
			},
		},
	}
}
