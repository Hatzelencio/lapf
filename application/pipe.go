package application

import (
	"encoding/json"
	"errors"
	"os"
)

var info os.FileInfo

func init() {
	var err error
	info, err = os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
}
func hasPipeContent() bool {
	return !(info.Mode()&os.ModeNamedPipe == 0)
}

func retrieveJSONFromStdin() ([]*ResultIsContainsOverlap, error) {
	var resultFromStdin []*ResultIsContainsOverlap
	err := json.NewDecoder(os.Stdin).Decode(&resultFromStdin)
	if err != nil {
		return nil, errors.New("bad input from stdin. JSON array does not found")
	}

	return resultFromStdin, nil
}

func retrieveCIDRBlockFromResults(input []*ResultIsContainsOverlap) []string {
	var stdinArguments []string
	m := make(map[string]bool)

	for _, item := range input {
		if m[item.CurrentCidr] {
			continue
		}

		m[item.CurrentCidr] = true
		stdinArguments = append(stdinArguments, item.CurrentCidr)
	}
	return stdinArguments
}

func retrieveStdinArguments() (stdinResults []*ResultIsContainsOverlap, arguments []string, err error) {
	stdinResults, err = retrieveJSONFromStdin()
	if err != nil {
		return nil, nil, err
	}
	arguments = retrieveCIDRBlockFromResults(stdinResults)

	return stdinResults, arguments, nil
}
