package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	header = "Location,Size,Beds,Baths,Price"
	data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

	separator = ","
)

func splitAndMapData(data string, separator string) ([]string, []int, []int, []int, []int, error) {
	var (
		locations []string
		sizes     []int
		beds      []int
		baths     []int
		prices    []int
	)

	eachRow := strings.Split(data, "\n")

	for _, row := range eachRow {
		eachValue := strings.Split(row, separator)

		if len(eachValue) != 5 {
			return nil, nil, nil, nil, nil, fmt.Errorf("Invalid data format\n")
		}

		locations = append(locations, eachValue[0])

		size, err := strconv.Atoi(eachValue[1])
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		sizes = append(sizes, size)

		bed, err := strconv.Atoi(eachValue[2])
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		beds = append(beds, bed)

		bath, err := strconv.Atoi(eachValue[3])
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		baths = append(baths, bath)

		price, err := strconv.Atoi(eachValue[4])
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		prices = append(prices, price)

	}

	return locations, sizes, beds, baths, prices, nil
}

func divider(spaceLength int) {
	fmt.Printf("%s\n", strings.Repeat("=", spaceLength*15))
}

func averageIntInput(num []int) float64 {
	var sum float64

	for _, v := range num {
		sum += float64(v)
	}

	return sum / float64(len(num))
}

func main() {
	locations, sizes, beds, baths, prices, err := splitAndMapData(data, separator)

	if err != nil {
		fmt.Println(err)
		return
	}

	headers := strings.Split(header, separator)
	for _, header := range headers {
		fmt.Printf("%-15s", header)
	}
	fmt.Println()

	divider(len(headers))
	for i := 0; i < len(locations); i++ {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d\n", locations[i], sizes[i], beds[i], baths[i], prices[i])
	}

	sizesAvg := averageIntInput(sizes)
	bedsAvg := averageIntInput(beds)
	bathsAvg := averageIntInput(baths)
	pricesAvg := averageIntInput(prices)

	divider(len(headers))
	fmt.Printf("%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", "", sizesAvg, bedsAvg, bathsAvg, pricesAvg)
}
