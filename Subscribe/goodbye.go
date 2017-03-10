// Reference:
// 1. Jack's Demo Repo: https://github.com/jhkolb/demosvc/blob/master/main.go
// 2. Documentation: https://godoc.org/gopkg.in/immesys/bw2bind.v3#BW2Client.SetEntity

// Learn how to use channels

// Steps:
// 1. check permission with setentity, set entity from environment variables, setentityfromenviron
// 2. Create string payload object (createstringpayloadobject?)
// 3. Replace fmt.println with bosswave publish (publish to namespace)
package main

import (
	"fmt"
	//"time"
	"os"

	bw2 "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	bwClient, err := bw2.Connect("")
	bwClient.SetEntityFromEnvironOrExit()

	if err != nil {
		fmt.Printf("Failed to connect to Bosswave agent: %v\n", err)
		os.Exit(1)
	}

	for {
		channel, subErr := bwClient.Subscribe(&bw2.SubscribeParams{
			URI:       "john/test",
			AutoChain: true,
		})

		if subErr != nil {
			fmt.Printf("Could not subscribe to URI")
			os.Exit(1)
		}

		msg := <- channel
		fmt.Println(msg)
	}

	//Form connection so that we can continuously check URI for updates

}
