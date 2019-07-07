package api

import "fmt"

var verAPI = "0.0.1"

//PrintVer -- printing version API
func PrintVer() string {
	return fmt.Sprintf("Version: %s", verAPI)
}
