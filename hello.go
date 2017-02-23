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
	message := "Message"
	bwClient, err := bw2.Connect("")
	bwClient.SetEntityFromEnvironOrExit()

	if err != nil {
		fmt.Printf("Failed to connect to Bosswave agent: %v\n", err)
		os.Exit(1)
	}
	//message = "round 2"
	for {
		payload := bw2.CreateStringPayloadObject(message)
		error := bwClient.Publish(&bw2.PublishParams{

			URI:            "john/test",
			AutoChain:      true,


			PayloadObjects: []bw2.PayloadObject{payload},
		})
		if (error != nil) {
			fmt.Print(error)
			os.Exit(1)
		}
		time.Sleep(10*time.Second)
	}
}
