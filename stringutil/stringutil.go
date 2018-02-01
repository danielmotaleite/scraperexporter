package stringutil

import "regexp"

func ExtractValue(bodyRaw string, regexString string) string {
	re:= regexp.MustCompile(regexString)
	regexpResult := re.FindAllStringSubmatch(bodyRaw, -1)

	if len(regexpResult) > 0 {
		return regexpResult[0][1]
	}

	return "0"
}

func BuildSingleResultLine(fullyUrl string, value string, customMetricString string) string {
	if customMetricString == "" {
		customMetricString = "extracted_value"
	}

	return customMetricString + "{url=\"" + fullyUrl + "\"} " + value
}
