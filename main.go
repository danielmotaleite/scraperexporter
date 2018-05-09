package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/marceloalmeida/scraperexporter/configuration"
	"github.com/marceloalmeida/scraperexporter/nethelper"
	"github.com/marceloalmeida/scraperexporter/stringutil"
	"gopkg.in/tylerb/graceful.v1"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {

	configFile := flag.String("configuration-file", "conf.json", "Configuration file")
	generateConfigFile := flag.Bool("generate-config", false, "Generate configuration example file")
	flag.Parse()

	// @todo - not working as expected
	if *generateConfigFile == true {
		fmt.Println(configuration.BuildConfigExample())
		os.Exit(0)
	}

	config := configuration.ReadConfiguration(*configFile)

	var wg sync.WaitGroup

	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:        config.ListenerAddress,
			ReadTimeout: time.Duration(5) * time.Second,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf := &bytes.Buffer{}

				wg.Add(len(config.Sites))
				for _, site := range config.Sites {
					go func(site string) {
						defer wg.Done()

						fullURL := nethelper.BuildURL(site+config.Path, config.AvoidCache)
						bodyRaw := nethelper.GetHTTPBody(fullURL)
						extractedValue := stringutil.ExtractValue(bodyRaw, config.RegexpString)
						builtSingleResultLine := stringutil.BuildSingleResultLine(site, extractedValue, config.MetricString)

						buf.WriteString(builtSingleResultLine + "\n")
					}(site)
				}
				wg.Wait()

				io.Copy(w, buf)
			}),
		},
	}

	log.Println("Server is now listening on ", config.ListenerAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Listening server error: %s", err)
	}
}
