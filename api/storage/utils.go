package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func buildSort(sortField, sortOrder, sortFieldDefault, sortOrderDefault string) string {
	if sortField == "" {
		sortField = sortFieldDefault
	}

	if sortOrder == "desc" {
		sortOrder = "-"
	} else if sortOrder == "asc" {
		sortOrder = ""
	} else {
		sortOrder = sortOrderDefault
	}

	sort := fmt.Sprintf(`%s%s`, sortOrder, sortField)

	return sort
}

// PrintLogs prints the response status code and body content, can be used for debugging.
func PrintLogs(resp *http.Response, bodyContent []byte) {
	fmt.Println("=========================================")
	fmt.Println("Response StatusCode:", resp.StatusCode)

	var jsonBody interface{}
	if jsonErr := json.Unmarshal(bodyContent, &jsonBody); jsonErr == nil {
		prettyJSON, err := json.MarshalIndent(jsonBody, "", "  ")
		if err != nil {
			fmt.Println("Response Body (as string):", string(bodyContent))
		}

		fmt.Println("Response Body (as JSON):", string(prettyJSON))
	}

	fmt.Println("=========================================")
}
