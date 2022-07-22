// Return the current local time
package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	f = flag.String("f", "2006-01-02 15:04:05", "Date format strings like 2006-01-02 15:04")
)

func main() {
	flag.Parse()
	fmt.Println(time.Now().Format(*f))
}
