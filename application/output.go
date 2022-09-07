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
		var originAccount = fmt.Sprintf("%s %s", item.CloudAccount.ProviderName, item.CloudAccount.Id)
		var originNetwork = fmt.Sprintf("%s = %s [%s]", item.CloudNetwork.Id, item.CloudNetwork.Name, item.CloudNetwork.CidrBlock)
		if item.IsOverlap {
			fmt.Println(fmt.Sprintf("[X][%v]\tis overlapping at: \"%v\" (%v)", item.CurrentCidr, originAccount, originNetwork))
		} else {
			fmt.Println(fmt.Sprintf("[ ][%v]\tis not overlapping at: \"%v\" (%v)", item.CurrentCidr, originAccount, originNetwork))
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
