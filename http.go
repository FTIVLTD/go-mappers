package mappers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

/*
HTTP - http(s) mapper
*/
type HTTP struct {
	Host    string
	Port    int
	Path    string
	headers map[string]string
}

func (h *HTTP) AddHeader(key, value string) {
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[key] = value
}

/*
Load — loading data: GET request
*/
func (h *HTTP) Load(params interface{}) (interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", h.prepareURL(params), nil)
	for key, v := range h.headers {
		req.Header.Add(key, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

/*
Update — loading data: PUT request
*/
func (h *HTTP) Update(params, payload interface{}) (interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", h.prepareURL(params), nil)
	for key, v := range h.headers {
		req.Header.Add(key, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (h *HTTP) prepareURL(params interface{}) string {
	var query string
	if params != nil {
		query = prepareQuery(params)
	}
	protocol := "http://"
	if h.Port == 443 {
		protocol = "https://"
	}
	url := protocol + h.Host + ":" + strconv.FormatInt(int64(h.Port), 10) + h.Path + "?" + query
	fmt.Println("Query: ", url)
	return url

}

func prepareQuery(params interface{}) string {
	var q []string
	if v, ok := params.(map[string]interface{}); ok == true {
		for key, value := range v {
			q = append(q, key+"="+toString(value))
		}

	}
	return strings.Join(q, "&")
}

func toString(v interface{}) string {
	if d, ok := v.(string); ok == true {
		return d
	}
	if d, ok := v.([]byte); ok == true {
		return string(d)
	}
	if d, ok := v.(int64); ok == true {
		return strconv.FormatInt(d, 10)
	}
	if d, ok := v.(int); ok == true {
		return strconv.FormatInt(int64(d), 10)
	}
	if d, ok := v.(float64); ok == true {
		return strconv.FormatFloat(d, 'f', 6, 64)
	}
	a, _ := json.Marshal(v)
	return string(a)
}
