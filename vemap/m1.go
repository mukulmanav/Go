package vemap

import "fmt"

func getpackagename() string {
	return vemap
}

func createmap() {
	statepopulation := map[string]int{
		"california": 39256,
		"Texas":      2785,
		"Florida":    2061,
		"Ohio":       1161,
	}
	fmt.Println(statepopulation)

}
