package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ChallengeResponse struct {
	Follow  string `json:"follow"`
	Message string `json:"message"`
}

func main() {
	url := "https://www.letsrevolutionizetesting.com/challenge.json"
	count := 0

	for {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error when making the HTTP request: %s\n", err.Error())
			return
		}

		defer response.Body.Close()

		var challenge ChallengeResponse
		err = json.NewDecoder(response.Body).Decode(&challenge)
		if err != nil {
			fmt.Printf("Error decoding the JSON response: %s\n", err.Error())
			return
		}

		count++
		fmt.Printf("New URL: %s\n", challenge.Follow)

		// Extract the ID from new URL
		parts := strings.Split(challenge.Follow, "=")
		id := parts[len(parts)-1]

		// Make the next URL with the extracted ID
		url = fmt.Sprintf("https://www.letsrevolutionizetesting.com/challenge.json?id=%s", id)

		if challenge.Message == "This is not the end" {
			// Continue the process
			continue
		} else {
			// Process completed
			fmt.Printf("Final message: %s\n", challenge.Message)
			break
		}
	}

	fmt.Println("Process completed. We had %i URLs in the process.", count)
}

// Result:
// ...
// New URL: https://www.letsrevolutionizetesting.com/challenge?id=702496729
// New URL: https://www.letsrevolutionizetesting.com/challenge?id=94985935
// New URL: https://www.letsrevolutionizetesting.com/challenge?id=951438732
// New URL: https://www.letsrevolutionizetesting.com/challenge?id=903777032
// New URL: https://www.letsrevolutionizetesting.com/challenge?id=503047199
// New URL:
// Final message: Congratulations! You've reached the end! You have passed our simple test and we would love to hear from you. Please save the code you used to a private GitHub gist and fill out the application form at https://jobs.lever.co/rainforest with the link to the gist. We'll be in touch shortly!
// Process completed. We had %i URLs in the process. 100
