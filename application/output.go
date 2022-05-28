package application

import (
	"encoding/json"
	"fmt"
	"strings"
)

func PrintResults(output string, results []*ResultIsContainsOverlap) {
	if strings.ToLower(output) == "json" {
		printResultsToJson(results)
	} else {
		printResultToText(results)
	}
}

func printResultsToJson(results []*ResultIsContainsOverlap) {
	e, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}

func printResultToText(results []*ResultIsContainsOverlap) {
	for _, item := range results {
		if item.IsOverlap {
			fmt.Println(fmt.Sprintf("[X][%v]\tis overlapping at: \"%v\" (%v)", item.CurrentCidr, item.CloudNetwork.ProviderName, item.CloudNetwork.Name))
		} else {
			fmt.Println(fmt.Sprintf("[ ][%v]\tis not overlapping at: \"%v\" (%v)", item.CurrentCidr, item.CloudNetwork.ProviderName, item.CloudNetwork.Name))
		}
	}
}
