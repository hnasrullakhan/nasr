package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

func (t *Cluster) MarshalJSON() ([]byte, error) {
	fmt.Println("MarshalJSON called")
	var j interface{}
	b, _ := bson.Marshal(t)
	bson.Unmarshal(b, &j)
	return json.Marshal(&j)
}

func (t *Cluster) UnmarshalJSON(b []byte) error {
	fmt.Println("UnmarshalJSON called")
	var j map[string]interface{}
	json.Unmarshal(b, &j)
	b, _ = bson.Marshal(&j)
	return bson.Unmarshal(b, t)
}

func Flatten(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := Flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}
func EncodeCredentials(username string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func queryHXRESTApi(clusterRequestUrl string, clusterUsername string, clusterPassword string) []byte {

	// Ignores certificates which can not be validated
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// create a HTTP client
	var httpClient = http.Client{Transport: tr}

	// create a http Request pointer
	var req *http.Request
	var err error

	// Defines the HTTP Request
	// send a GET to the SpringPath API
	req, err = http.NewRequest("GET", clusterRequestUrl, nil)
	if err != nil {
		fmt.Println("Error: GET failed : ", err)
		return nil
	}

	// before the request is send set the HTTP Header key "Authorization" with
	// the value of base64 encoded Username and Password
	req.Header.Set("Authorization", "Basic "+EncodeCredentials(clusterUsername, clusterPassword))

	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error: Do Failed: ", err)
		return nil
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ReadAll failed: ", err)
		return nil
	}

	cerr := resp.Body.Close()
	if cerr != nil {
		fmt.Println("Error: closing http resp, ", cerr)
	}

	return responseData
}

type Cluster struct {
	ClusterUUID  string                 `json:"cluster_uuid" bson:"cluster_uuid"`
	ClusterName  string                 `json:"cluster_name" bson:"cluster_name"`
	DynamicProps map[string]interface{} `json:"-" bson:",inline"`
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func main() {
	summaryResponse := queryHXRESTApi("https://"+"10.193.33.175"+"/rest/summary", "root", "Cisco123")

	buf := []byte(`{"cluster_uuid": "90a13c89-bdd7-4df5-9690-17980d918937","cluster_name": "skasanav-hx-dev", "summary":`)
	end := []byte(`}`)

	result1 := append(buf[:], summaryResponse[:]...)
	result2 := append(result1[:], end[:]...)

	fmt.Println(jsonPrettyPrint(string(result2)))

	fmt.Println("******************")
	c := Cluster{}
	json.Unmarshal(result2, &c)
	fmt.Println(c.ClusterUUID)
	fmt.Println("******************")
	fmt.Println(c.DynamicProps["summary"].(map[string]interface{})["uptime"])
	fmt.Println("******************")
	fmt.Println(c.DynamicProps["summary"].(map[string]interface{})["resiliencyInfo"].(map[string]interface{})["state"])
	fmt.Println("******************")

	fmt.Println(Flatten(c.DynamicProps)["summary.resiliencyDetails.HEALTH_REASON"])
	fmt.Println("***************")
	b2, _ := json.Marshal(Flatten(c.DynamicProps))
	fmt.Println(jsonPrettyPrint(string(b2)))

}
