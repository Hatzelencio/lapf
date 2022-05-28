package application

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func New() {
	app := &cli.App{
		Name:  "Overlapping Finder",
		Usage: "Overlapping Finder (a.k.a \"lapf\") is a binary-tool made and built in golang to find possible overlapping CIDR Block notation through cloud providers (AWS)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "provider",
				Usage:       "Cloud Provider",
				DefaultText: "aws",
				Value:       "aws",
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "Output format",
				DefaultText: "text, json",
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
	}
}
