package nethelper

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// GetHTTPBody Get Http body from an URL
// It returns string
func GetHTTPBody(url string) string {
	request, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Panicf("Error making HTTP request: %s", err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("Error reading HTTP request body: %s", err)
		return ""
	}

	return string(body)
}

// BuildURL build encoded URL with optional query string to avoid cache
// It returns string
func BuildURL(urlRaw string, avoidCache ...string) string {
	u, err := url.Parse(urlRaw)

	if err != nil {
		log.Fatalf("Error building URL: %s", err)
	}

	q := u.Query()

	if len(avoidCache) > 0 && avoidCache[0] == "true" {
		queryString := strconv.FormatInt(time.Now().Unix(), 10)
		q.Set("z"+queryString, queryString)
	}

	u.RawQuery = CompatibleRFC3986Encode(q.Encode())

	return u.String()
}

// CompatibleRFC3986Encode Compatible with RFC 3986.
func CompatibleRFC3986Encode(str string) string {
	resultStr := str
	resultStr = strings.Replace(resultStr, "+", "%2B", -1)
	return resultStr
}
