package main

import (
	"encoding/json"
	"example.com/custom-otel-playground/config"
	"example.com/custom-otel-playground/model"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	promServer      = "localhost"
	port            = "9090"
	protocol        = "http"
	limitSeries     = 5
	promQueryStatus = "success"
)

var (
	promResponse     model.PrometheusQueryResponse
	promTSDBResponse model.PromTSDBResponse

	multiTenant      = false
	querySeriesCount = "count({product=\"abc\", __name__!~\"up|scrape_.*\"})"
	studioSeries     = make(map[string]string)
	studioSeriesInt  = make(map[string]int)
)

func decodePractice() {
	obj := make(map[string]string)
	//var obj interface{}

	yamlFile, err := os.ReadFile("/custom-otel-playground/script/processor.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v\n", err)
	}

	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		fmt.Printf("Unmarshal: %v\n", err)
	}

	var cfg config.Config

	err = mapstructure.Decode(obj, &cfg)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
	}
	fmt.Println(obj)
	fmt.Println(cfg.PrometheusTarget)
}

func queryResponse() *model.PrometheusQueryResponse {
	promURL := fmt.Sprintf("%s://%s:%s/api/v1/query", protocol, promServer, port)

	params := url.Values{}
	params.Add("query", querySeriesCount)

	promQuery, _ := url.ParseRequestURI(promURL)
	promQuery.RawQuery = params.Encode()

	promQueryStr := fmt.Sprintf("%v", promQuery)

	fmt.Println(promQueryStr)

	resp, err := http.Get(promQueryStr)
	if err != nil {
		fmt.Printf("error on http call: %s\n", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error tp read body: %s\n", err)
		return nil
	}

	err = json.Unmarshal(body, &promResponse)
	if err != nil {
		fmt.Printf("error tp unmarshall body: %s\n", err)
		return nil
	}

	return &promResponse
}

func countActiveSeries() {
	//multiTenant = true
	//querySeriesCount = "count({__name__=~\"ext_.*\"}) by (studio)"

	queryResult := queryResponse()
	if multiTenant == true {
		listStudios := []string{"studio1", "studio2", "studio3"}
		for _, result := range queryResult.Data.Result {
			studioSeries[result.Metric["studio"]] = fmt.Sprintf("%v", result.Value[1])
		}

		for _, studio := range listStudios {
			currentSeries, err := strconv.Atoi(studioSeries[studio])
			if err != nil {
				fmt.Printf("error to convert string to in: %s\n", err)
				return
			}

			if currentSeries >= limitSeries {
				fmt.Printf("filter metrics, studio: %s\n", studio)
			} else {
				fmt.Printf("keep metrics, studio: %s\n", studio)
			}
		}
		return
	}

	currentSeries, err := strconv.Atoi(fmt.Sprintf("%v", queryResult.Data.Result[0].Value[1]))
	if err != nil {
		fmt.Printf("error to convert string to in: %s\n", err)
		return
	}

	if currentSeries >= limitSeries {
		fmt.Printf("filter metrics, current series: %d | limit: %d\n", currentSeries, limitSeries)
	} else {
		fmt.Printf("keep metrics, current series: %d | limit: %d\n", currentSeries, limitSeries)
	}
}

func tsdbQueryResponse() *model.PromTSDBResponse {
	promURL := fmt.Sprintf("%s://%s:%s/api/v1/status/tsdb", protocol, promServer, port)
	fmt.Println(promURL)

	resp, err := http.Get(promURL)
	if err != nil {
		fmt.Printf("error on http call: %s\n", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error tp read body: %s\n", err)
		return nil
	}

	err = json.Unmarshal(body, &promTSDBResponse)
	if err != nil {
		fmt.Printf("error tp unmarshall body: %s\n", err)
		return nil
	}

	return &promTSDBResponse
}

func main() {
	//decodePractice()

	//multiTenant = true
	//tsdbResult := tsdbQueryResponse()
	queryResult := queryResponse()

	//if multiTenant {
	//	listStudios := []string{"studio1", "studio2", "studio3"}
	//	for _, result := range tsdbResult.Data.SeriesCountByLabelValuePair {
	//		studioPattern := `^studio=.*$`
	//		re := regexp.MustCompile(studioPattern)
	//
	//		if re.MatchString(result.Name) {
	//			parts := strings.SplitN(result.Name, "=", 2)
	//			studioSeriesInt[parts[1]] = result.Value
	//		}
	//	}
	//
	//	for _, studio := range listStudios {
	//		currentSeries := studioSeriesInt[studio]
	//
	//		if currentSeries >= limitSeries {
	//			fmt.Printf("filter metrics, studio: %s\n", studio)
	//		} else {
	//			fmt.Printf("keep metrics, studio: %s\n", studio)
	//		}
	//	}
	//	return
	//}

	if queryResult == nil || queryResult.Status != promQueryStatus {
		fmt.Println("failed when execute query")
		return
	}

	var currentSeries int

	if queryResult.Status == promQueryStatus && len(queryResult.Data.Result) > 0 {
		for _, result := range queryResult.Data.Result {
			numSeriesStr, ok := result.Value[1].(string)
			if !ok {
				fmt.Println("failed to parse interface to string")
				return
			}

			numSeriesInt, err := strconv.Atoi(numSeriesStr)
			if err != nil {
				fmt.Println("failed to string parse to int, error: ", err)
				return
			}

			currentSeries += numSeriesInt
		}
	}

	fmt.Println(currentSeries)

	//currentSeries := tsdbResult.Data.HeadStats.NumSeries - 5
	//currentSeries := queryResult.Data.Result[0]

	if currentSeries >= limitSeries {
		fmt.Printf("filter metrics, current series: %d | limit: %d\n", currentSeries, limitSeries)
	} else {
		fmt.Printf("keep metrics, current series: %d | limit: %d\n", currentSeries, limitSeries)
	}
}
