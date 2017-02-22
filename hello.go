// Reference:
// 1. Jack's Demo Repo: https://github.com/jhkolb/demosvc/blob/master/main.go
// 2. Documentation: https://godoc.org/gopkg.in/immesys/bw2bind.v3#BW2Client.SetEntity

// Steps:
// 1. check permission with setentity, set entity from environment variables, setentityfromenviron
// 2. Create string payload object (createstringpayloadobject?)
// 3. Replace fmt.println with bosswave publish (publish to namespace)

import (
	"fmt"
	"time"
	"os"

	bw2 "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	message := "Message"
	bwClient, err := bw2.Connect("")
	bwClient.SetEntityFromEnvironOrExit()

	if err != nil {
		fmt.Printf("Failed to connect to Bosswave agent: %v\n", err)
		os.Exit(1)
	}

	for {
		payload := bw2.CreateStringPayloadObject(message)
		bwClient.PublishOrExit(&bw2.PublishParams{
			URI:            parameters.MustString("to"),
			AutoChain:      true,
			PayloadObjects: []bw2.PayloadObject{payload},
		})
		time.Sleep(10*time.Second)
	}
}
