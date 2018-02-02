package stringutil

import "regexp"

// ExtractValue extracts value from raw message using regexp
// It returns string
func ExtractValue(bodyRaw string, regexString string) string {
	re := regexp.MustCompile(regexString)
	regexpResult := re.FindAllStringSubmatch(bodyRaw, -1)

	if len(regexpResult) > 0 {
		return regexpResult[0][1]
	}

	return "0"
}

// BuildSingleResultLine build single line (Prometheus format)
// It returns string
func BuildSingleResultLine(fullyURL string, value string, customMetricString string) string {
	if customMetricString == "" {
		customMetricString = "extracted_value"
	}

	return customMetricString + "{url=\"" + fullyURL + "\"} " + value
}
