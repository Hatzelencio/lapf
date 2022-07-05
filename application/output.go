package application

import (
	"encoding/json"
	"fmt"
	"strings"
)

func PrintOverlapResults(output string, results []*ResultIsContainsOverlap) {
	if strings.ToLower(output) == "json" {
		PrintResultsToJson(results)
	} else {
		printOverlapResultToText(results)
	}
}

func printOverlapResultToText(results []*ResultIsContainsOverlap) {
	for _, item := range results {
		if item.IsOverlap {
			fmt.Println(fmt.Sprintf("[X][%v]\tis overlapping at: \"%v\" (%v)", item.CurrentCidr, item.CloudNetwork.ProviderName, item.CloudNetwork.Name))
		} else {
			fmt.Println(fmt.Sprintf("[ ][%v]\tis not overlapping at: \"%v\" (%v)", item.CurrentCidr, item.CloudNetwork.ProviderName, item.CloudNetwork.Name))
		}
	}
}

func PrintIsCIDRPrivateResults(output string, results []*ResultIsPrivateCIDRBlock) {
	if strings.ToLower(output) == "json" {
		PrintResultsToJson(results)
	} else {
		printIsCIDRPrivateResultToText(results)
	}
}

func printIsCIDRPrivateResultToText(results []*ResultIsPrivateCIDRBlock) {
	for _, item := range results {
		if item.IsPrivate {
			fmt.Println(fmt.Sprintf("[✓][%v] is private", item.CIDR))
		} else {
			fmt.Println(fmt.Sprintf("[x][%v] is not private", item.CIDR))
		}
	}
}

func PrintResultsToJson(results interface{}) {
	e, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
}
