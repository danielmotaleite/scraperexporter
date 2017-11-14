package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"time"
	"strconv"
	"net/url"
	"regexp"
	"os"
	"encoding/json"
	"gopkg.in/tylerb/graceful.v1"
	"bytes"
	"io"
	"flag"
	"sync"
)

type Configuration struct {
	Sites []string
	MetricString string
	AvoidCache string
	Path string
	ListenerAddress string
	RegexpString string
}

func getHttpBody(url string) string {
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

func buildURL(urlRaw string, avoidCache ...string) string {
	u, err  := url.Parse(urlRaw)

	if err != nil {
		log.Fatalf("Error building URL: %s", err)
	}

	if len(avoidCache) > 0 && avoidCache[0] == "true" {
		queryString := strconv.FormatInt(time.Now().Unix(), 10)
		q := u.Query()
		q.Set("z" + queryString, queryString)
		u.RawQuery = q.Encode()
	}

	return u.String()
}

func extractValue(bodyRaw string, regexString string) string {
	re:= regexp.MustCompile(regexString)
	regexpResult := re.FindAllStringSubmatch(bodyRaw, -1)

	if len(regexpResult) > 0 {
		return regexpResult[0][1]
	}

	return "0"
}

func buildSingleResultLine(fullyUrl string, value string, custom_metric_string string) string {
	if custom_metric_string == "" {
		custom_metric_string = "extracted_value"
	}

	return custom_metric_string + "{url=\"" + fullyUrl + "\"} " + value
}

func readConfiguration(filename string) Configuration {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("Error reading configuration file: '%s'", err)
	}

	return configuration
}

func buildConfigExample() *bytes.Buffer {
	configuration := Configuration{}

	configuration.Sites = append(configuration.Sites, "https://www.jumia.com.ng")
	configuration.Sites = append(configuration.Sites, "https://www.jumia.com.eg")
	configuration.MetricString = "value"
	configuration.AvoidCache = "true"
	configuration.Path = "/products/"
	configuration.ListenerAddress = "0.0.0.0:8080"
	configuration.RegexpString = "(.*)"

	buf := &bytes.Buffer{}

	jsonObject := json.NewEncoder(buf)
	jsonObject.SetIndent("", "  ")
	jsonObject.Encode(configuration)

	return buf
}

func main() {

	configFile := flag.String("configuration-file", "conf.json", "Configuration file")
	generateConfigFile := flag.Bool("generate-config", false, "Generate configuration example file")
	flag.Parse()

	// @todo - not working as expected
	if *generateConfigFile == true {
		fmt.Println(buildConfigExample())
		os.Exit(0)
	}

	config := readConfiguration(*configFile)

	var wg sync.WaitGroup

	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr: config.ListenerAddress,
			ReadTimeout: time.Duration(5) * time.Second,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf := &bytes.Buffer{}

				wg.Add(len(config.Sites))
				for _, site := range config.Sites {
					go func(site string) {
						defer wg.Done()

						fullUrl := buildURL(site + config.Path, config.AvoidCache)
						bodyRaw := getHttpBody(fullUrl)
						extractedValue := extractValue(bodyRaw, config.RegexpString)
						builtSingleResultLine := buildSingleResultLine(site, extractedValue, config.MetricString)

						buf.WriteString(builtSingleResultLine + "\n")
					}(site)
				}
				wg.Wait()

				io.Copy(w,buf)
			}),
		},
	}

	log.Println("Server is now listening on ", config.ListenerAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Listening server error: %s", err)
	}
}
