package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

	separator = ","
)

const (
	//usage              = "Provide only the [starting] and [stopping] positions"
	//errInvalidPosition = "Wrong positions. Starting index >= 0 && stopping pos < %d\n"
	errInvalidInput = "Invalid input"
)

func splitAndMapData(data string, separator string) ([]string, []string, []string, []string, []string, error) {
	var (
		locations []string
		sizes     []string
		beds      []string
		baths     []string
		prices    []string
	)

	eachRow := strings.Split(data, "\n")

	for _, row := range eachRow {
		eachValue := strings.Split(row, separator)

		if len(eachValue) != 5 {
			return nil, nil, nil, nil, nil, fmt.Errorf("Invalid data format\n")
		}

		locations = append(locations, eachValue[0])
		sizes = append(sizes, eachValue[1])
		beds = append(beds, eachValue[2])
		baths = append(baths, eachValue[3])
		prices = append(prices, eachValue[4])
	}

	return locations, sizes, beds, baths, prices, nil
}

func filteringSlices(housingData [][]string, headers []string) {
	for i := range housingData[0] {
		if i == 0 {
			for _, header := range headers {
				fmt.Printf("%-15s", header)
			}
			fmt.Println()
		}

		for j := range housingData {
			fmt.Printf("%-15s", housingData[j][i])
		}
		fmt.Println()
	}
}

func main() {
	housingData := make([][]string, 0)
	var headers []string

	locations, sizes, beds, baths, prices, err := splitAndMapData(data, separator)

	if err != nil {
		fmt.Println(errInvalidInput)
		return
	}

	housingData = append(housingData, [][]string{locations[1:]}...)
	housingData = append(housingData, [][]string{sizes[1:]}...)
	housingData = append(housingData, [][]string{beds[1:]}...)
	housingData = append(housingData, [][]string{baths[1:]}...)
	housingData = append(housingData, [][]string{prices[1:]}...)

	headers = append(headers, locations[0], sizes[0], beds[0], baths[0], prices[0])

	indexes := map[string]int{
		"Location": 0,
		"Size":     1,
		"Beds":     2,
		"Baths":    3,
		"Price":    4,
	}

	start := indexes["Location"]
	end := indexes["Price"] + 1

	args := os.Args[1:]

	if len(args) > 2 {
		fmt.Println("Error")
		return
	}

	if len(args) > 0 {
		startCommand := args[0]

		if _, ok := indexes[startCommand]; !ok {
			start = indexes["Location"]
		} else {
			start = indexes[startCommand]
		}
	}

	if len(args) > 1 {
		endCommand := args[1]

		if _, ok := indexes[endCommand]; !ok {
			end = indexes["Price"] + 1
		} else {
			end = indexes[endCommand] + 1
		}
	}

	if start < end {
		headers = headers[start:end]
		housingData = housingData[start:end]
		filteringSlices(housingData, headers)
	} else if start > end {
		var newHeaders []string
		var newhousingData [][]string

		newHeaders = append(newHeaders, headers[start+1:]...)
		newHeaders = append(newHeaders, headers[:end]...)

		newhousingData = append(newhousingData, housingData[start+1:]...)
		newhousingData = append(newhousingData, housingData[:end]...)

		filteringSlices(newhousingData, newHeaders)
	}
}
