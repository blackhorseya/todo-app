package exit

import (
	"fmt"
	"os"
)

// Er print error msg and exit by status code 1
func Er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
