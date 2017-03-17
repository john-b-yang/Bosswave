// Reference:
// 1. Jack's Demo Repo: https://github.com/jhkolb/demosvc/blob/master/main.go
// 2. Documentation: https://godoc.org/gopkg.in/immesys/bw2bind.v3#BW2Client.SetEntity

// Steps:
// 1. check permission with setentity, set entity from environment variables, setentityfromenviron
// 2. Create string payload object (createstringpayloadobject?)
// 3. Replace fmt.println with bosswave publish (publish to namespace)
package main
import (
	"fmt"
	"time"
	"os"

	bw2 "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	message := "Message 1 2 3"
	message2 := "Message 4 5 6"
	bwClient, err := bw2.Connect("")
	bwClient.SetEntityFromEnvironOrExit()

	if err != nil {
		fmt.Printf("Failed to connect to Bosswave agent: %v\n", err)
		os.Exit(1)
	}

	for {
		payload := bw2.CreateStringPayloadObject(message)
		payload2 := bw2.CreateStringPayloadObject(message2)
		sendPayload(payload)
		sendPayload(payload2)
		time.Sleep(10*time.Second)
	}
}

func sendPayload(payload string) {
	err = bwClient.Publish(&bw2.PublishParams{
		URI:            "john/test",
		AutoChain:      true,
		PayloadObjects: []bw2.PayloadObject{payload, payload},
	})
	if (err != nil) {
		fmt.Print(err)
		os.Exit(1)
	}
}
