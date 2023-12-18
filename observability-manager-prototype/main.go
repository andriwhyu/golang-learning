package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Access Policy & Token Struct
type AccessPolicy struct {
	Name        string              `json:"name"`
	DisplayName string              `json:"displayName"`
	Scopes      []string            `json:"scopes"`
	Realms      []RealmAccessPolicy `json:"realms"`
}

type AccessPolicyResult struct {
	AccessPolicyID   string `json:"id"`
	AccessPolicyName string `json:"name"`
}

type AccessPolicyTokenRequestBody struct {
	AccessPolicyID string `json:"accessPolicyId"`
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
}

type AccessPolicyTokenResult struct {
	TokenID        string `json:"id"`
	AccessPolicyID string `json:"accessPolicyId"`
	TokenName      string `json:"name"`
	Token          string `json:"token"`
}

type RealmAccessPolicy struct {
	Type        string                `json:"type"`
	Identifier  string                `json:"identifier"`
	LabelPolicy []LabelSelectorPolicy `json:"labelPolicies"`
}

type LabelSelectorPolicy struct {
	Selector string `json:"selector"`
}

// Data Source Struct
type PrometheusRequestBody struct {
	Name           string             `json:"name"`
	Type           string             `json:"type"`
	Access         string             `json:"access"`
	Url            string             `json:"url"`
	User           string             `json:"user"`
	Database       string             `json:"database"`
	BasicAuth      bool               `json:"basicAuth"`
	BasicAuthUser  string             `json:"basicAuthUser"`
	JsonData       JsonDataPrometheus `json:"jsonData"`
	SecureJsonData SecureJsonData     `json:"secureJsonData"`
}

type LokiRequestBody struct {
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	Access         string         `json:"access"`
	Url            string         `json:"url"`
	BasicAuth      bool           `json:"basicAuth"`
	BasicAuthUser  string         `json:"basicAuthUser"`
	JsonData       JsonDataLoki   `json:"jsonData"`
	SecureJsonData SecureJsonData `json:"secureJsonData"`
}

type DataSourceResponse struct {
	DatasourceID int    `json:"id"`
	Name         string `json:"name"`
}

type DataSourcePermission struct {
	DatasourceID int           `json:"datasourceId"`
	Enabled      bool          `json:"enabled"`
	Permissions  []Permissions `json:"permissions"`
}

type TeamPermission struct {
	TeamID     int `json:"teamId"`
	Permission int `json:"permission"`
}

type Permissions struct {
	PermissionsID  int    `json:"id"`
	DatasourceID   int    `json:"datasourceId"`
	UserID         int    `json:"userId,omitempty"`
	TeamID         int    `json:"teamId,omitempty"`
	BuiltInRole    string `json:"builtInRole,omitempty"`
	Permission     int    `json:"permission"`
	PermissionName string `json:"permissionName"`
}

type JsonDataPrometheus struct {
	HttpMethod        string `json:"httpMethod"`
	PrometheusType    string `json:"prometheusType"`
	TimeInterval      string `json:"timeInterval"`
	Timeout           int    `json:"timeout"`
	PrometheusVersion string `json:"prometheusVersion"`
}

type JsonDataLoki struct {
	MaxLines          int                 `json:"maxLines"`
	DerivedFieldsLoki []DerivedFieldsLoki `json:"derivedFields,omitempty"`
}

type DerivedFieldsLoki struct {
	DatasourceUid   string `json:"datasourceUid,omitempty"`
	MatcherRegex    string `json:"matcherRegex,omitempty"`
	Name            string `json:"name,omitempty"`
	Url             string `json:"url,omitempty"`
	UrlDisplayLabel string `json:"urlDisplayLabel,omitempty"`
}

type SecureJsonData struct {
	BasicAuthPassword string `json:"basicAuthPassword"`
}

func createAccessPolicy(newGame, envLabel, accessPolicyNaming string, scopes []string) AccessPolicyResult {
	var accessPolicyResult AccessPolicyResult

	labelSelector := LabelSelectorPolicy{fmt.Sprintf("{environment_name=\"%s\", game_namespace=\"%s\"}", envLabel, newGame)}
	realmPolicy := RealmAccessPolicy{"stack", os.Getenv("GRAFANA_CENTRAL_STACK_ID"), []LabelSelectorPolicy{labelSelector}}
	accessPolicy := AccessPolicy{accessPolicyNaming, accessPolicyNaming, scopes, []RealmAccessPolicy{realmPolicy}}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(accessPolicy)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://grafana.com/api/v1/accesspolicies?region=us", payloadBuf)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_API_KEY")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &accessPolicyResult)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	fmt.Println(string(resp.Status))
	// fmt.Println(client)

	return accessPolicyResult
}

func createTokenAccessPolicy(newGame, envLabel, accessPolicyNaming string, scopes []string, accessPolicyResult AccessPolicyResult) AccessPolicyTokenResult {
	var accessPolicyTokenResult AccessPolicyTokenResult

	accessPolicyTokenRequestBody := AccessPolicyTokenRequestBody{accessPolicyResult.AccessPolicyID, accessPolicyNaming, accessPolicyNaming}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(accessPolicyTokenRequestBody)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://grafana.com/api/v1/tokens?region=us", payloadBuf)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_API_KEY")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &accessPolicyTokenResult)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	fmt.Println(string(resp.Status))
	// fmt.Println(client)

	return accessPolicyTokenResult
}

func createPrometheusDataSource(newGame, envLabel, accessPolicyNaming string, scopes []string, accessPolicyToken AccessPolicyTokenResult) DataSourceResponse {
	var dataSourceResponse DataSourceResponse

	datasourceName := fmt.Sprintf("%s-%s", "prom", accessPolicyNaming)

	secureJson := SecureJsonData{accessPolicyToken.Token}
	jsonData := JsonDataPrometheus{"POST", "Mimir", "60s", 150, "2.3.0"}

	prometheusReq := PrometheusRequestBody{
		datasourceName,
		"prometheus",
		"proxy",
		fmt.Sprintf("https://%s/api/prom", os.Getenv("GRAFANA_CENTRAL_METRICS_URL")),
		"",
		"",
		true,
		os.Getenv("GRAFANA_CENTRAL_METRICS_USER_ID"),
		jsonData,
		secureJson,
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(prometheusReq)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/datasources", os.Getenv("GRAFANA_STACK_URL")), payloadBuf)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_SERVICE_ACCOUNT_TOKEN")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &dataSourceResponse)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	fmt.Println(string(resp.Status))
	// fmt.Println(client)

	return dataSourceResponse
}

func createLokiDataSource(newGame, envLabel, accessPolicyNaming string, scopes []string, accessPolicyToken AccessPolicyTokenResult) DataSourceResponse {
	var dataSourceResponse DataSourceResponse

	datasourceName := fmt.Sprintf("%s-%s", "loki", accessPolicyNaming)

	secureJson := SecureJsonData{accessPolicyToken.Token}
	derivedFields := DerivedFieldsLoki{
		DatasourceUid: "grafanacloud-traces",
		MatcherRegex:  `[tT]race_?[iI][dD]"?[:=]"?(\w+)`,
		Name:          "traceID",
		Url:           `${__value.raw}`,
	}

	jsonData := JsonDataLoki{MaxLines: 1000, DerivedFieldsLoki: []DerivedFieldsLoki{derivedFields}}

	lokiReq := LokiRequestBody{
		Name:           datasourceName,
		Type:           "loki",
		Access:         "proxy",
		Url:            fmt.Sprintf("https://%s", os.Getenv("GRAFANA_CENTRAL_LOGS_URL")),
		BasicAuth:      true,
		BasicAuthUser:  os.Getenv("GRAFANA_CENTRAL_LOGS_USER_ID"),
		JsonData:       jsonData,
		SecureJsonData: secureJson,
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(lokiReq)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/datasources", os.Getenv("GRAFANA_STACK_URL")), payloadBuf)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_SERVICE_ACCOUNT_TOKEN")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &dataSourceResponse)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	fmt.Println(string(resp.Status))
	// fmt.Println(client)

	return dataSourceResponse
}

func getDataSourcePermission(datasource DataSourceResponse) DataSourcePermission {
	var dataSourcePermission DataSourcePermission

	client := http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s/api/datasources/%d/permissions", os.Getenv("GRAFANA_STACK_URL"), datasource.DatasourceID), nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_SERVICE_ACCOUNT_TOKEN")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &dataSourcePermission)

	if err != nil {
		fmt.Println(err)
	}

	return dataSourcePermission
}

func resetDataSourcePermission(datasourcePermission DataSourcePermission) {
	client := http.Client{}

	for i := 0; i < len(datasourcePermission.Permissions); i++ {
		permission := datasourcePermission.Permissions[i]

		if datasourcePermission.Permissions[i].BuiltInRole != "" {
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("https://%s/api/datasources/%d/permissions/%d", os.Getenv("GRAFANA_STACK_URL"), permission.DatasourceID, permission.PermissionsID), nil)

			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_SERVICE_ACCOUNT_TOKEN")))
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)

			if err != nil {
				log.Fatal(err)
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(body))
			fmt.Println(string(resp.Status))
		}
	}
}

func addPermission(datasource DataSourceResponse) {
	datasourcePermission := getDataSourcePermission(datasource)

	resetDataSourcePermission(datasourcePermission)

	// 1 = Query permission
	permission := 1
	teamID := 3

	teamPermission := TeamPermission{teamID, permission}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(teamPermission)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/datasources/%d/permissions", os.Getenv("GRAFANA_STACK_URL"), datasource.DatasourceID), payloadBuf)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_SERVICE_ACCOUNT_TOKEN")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	fmt.Println(string(resp.Status))
}

func getAllDataSource() []DataSourceResponse {
	var datasourceResponse []DataSourceResponse

	client := http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s//api/datasources", os.Getenv("GRAFANA_STACK_URL")), nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GRAFANA_ADMIN_SERVICE_ACCOUNT_TOKEN")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &datasourceResponse)

	if err != nil {
		fmt.Println(err)
	}

	return datasourceResponse
}

func initializeDataSourcePermission() {
	datasource := getAllDataSource()

	for _, d := range datasource {
		resetDataSourcePermission(getDataSourcePermission(d))
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading file")
	}

	mydir, err := os.Getwd()

	fmt.Println(mydir)

	// initializeDataSourcePermission()

	newGame := "andritesttest2"
	envLabel := fmt.Sprintf("%s-%s-%s", os.Getenv("CUSTOMER_NAME"), os.Getenv("PROJECT_NAME"), os.Getenv("ENVIRONMENT_NAME"))
	accessPolicyNaming := fmt.Sprintf("%s-%s", envLabel, newGame)
	scopes := []string{"metrics:read", "logs:read", "traces:read", "alerts:read"}

	fmt.Println("=============Create Access Policy=============")
	accessPolicyResult := createAccessPolicy(newGame, envLabel, accessPolicyNaming, scopes)
	fmt.Println("=============Create Token=============")
	accessPolicyToken := createTokenAccessPolicy(newGame, envLabel, accessPolicyNaming, scopes, accessPolicyResult)
	fmt.Println("=============Create Prometheus Data Source=============")
	promDatasource := createPrometheusDataSource(newGame, envLabel, accessPolicyNaming, scopes, accessPolicyToken)
	addPermission(promDatasource)
	fmt.Println("=============Create Loki Data Source=============")
	lokiDatasource := createLokiDataSource(newGame, envLabel, accessPolicyNaming, scopes, accessPolicyToken)
	addPermission(lokiDatasource)
	// new redis client

	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "",
	//})
	//
	//// test connection
	//pong, err := client.Ping().Result()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// return pong if server is online
	//fmt.Println(pong)
	//
	//err = client.Set("test", "user100", 0).Err()
	//if err != nil {
	//	panic(err)
	//}

}
