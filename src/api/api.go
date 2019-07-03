package api

import "fmt"

var verAPI = "0.0.1"

func printVerAPI() string {
	return fmt.Sprintf("Version: %s", verAPI)
}
