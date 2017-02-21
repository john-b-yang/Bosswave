Reference: 
import (
	"fmt"
	"time"
	"os"

	bw2 "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	bwClient, err := bw2.Connect("")
	if err != nil {
		fmt.Printf("Failed to connect to Bosswave agent: %v\n", err)
		os.Exit(1)
	}
	for {
		fmt.Println("Hello World")
		time.Sleep(10*time.Second)
	}
}
