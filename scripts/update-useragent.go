package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func updateUserAgent(version string) error {
	providerFile := "internal/provider/provider.go"

	// Read file
	content, err := ioutil.ReadFile(providerFile)
	if err != nil {
		return fmt.Errorf("error reading provider.go: %v", err)
	}

	// Create backup
	err = ioutil.WriteFile(providerFile+".backup", content, 0644)
	if err != nil {
		return fmt.Errorf("error creating backup: %v", err)
	}

	// UserAgent pattern
	userAgentPattern := regexp.MustCompile(`c\.UserAgent\s*=\s*"[^"]*"(?:\s*"[^"]*")*`)
	newUserAgent := fmt.Sprintf(`c.UserAgent = "MerakiTerraform/%s Cisco"`, version)

	var newContent string
	if userAgentPattern.Match(content) {
		// Replace existing UserAgent
		newContent = userAgentPattern.ReplaceAllString(string(content), newUserAgent)
	} else {
		// Insert before "data := MerakiProviderData"
		if strings.Contains(string(content), "data := MerakiProviderData") {
			newContent = strings.Replace(
				string(content),
				"data := MerakiProviderData",
				fmt.Sprintf("\t%s\n\n\tdata := MerakiProviderData", newUserAgent),
				1,
			)
		} else {
			return fmt.Errorf("could not find where to insert UserAgent")
		}
	}

	// Write updated content
	err = ioutil.WriteFile(providerFile, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing updated file: %v", err)
	}

	fmt.Printf("✅ UserAgent successfully updated to: MerakiTerraform/%s Cisco\n", version)
	return nil
}

func main() {
	// Version must be provided as argument
	if len(os.Args) < 2 {
		fmt.Println("❌ Error: version argument is required")
		fmt.Println("Usage: go run update-useragent-python.go <version>")
		os.Exit(1)
	}

	version := os.Args[1]
	fmt.Printf("Using version: %s\n", version)

	err := updateUserAgent(version)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Update completed")
}
