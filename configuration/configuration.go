package configuration

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

// Configuration structure
type Configuration struct {
	Sites           []string
	MetricString    string
	AvoidCache      string
	Path            string
	ListenerAddress string
	RegexpString    string
}

// ReadConfiguration read configuration from filename
// It returns Configuration struct
func ReadConfiguration(filename string) Configuration {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("Error reading configuration file: '%s'", err)
	}

	return configuration
}

// BuildConfigExample builds example of configuration file
func BuildConfigExample() *bytes.Buffer {
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
