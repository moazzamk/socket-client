package main

import (
	"flag"
	"socket-client/socketclient"
	"fmt"
)

// main is cli socket client
func main() {
	hostPtr := flag.String(`h`, `host`, `Hostname or IP with port. Example: abc.com:9043`)
	inputPtr := flag.String(`i`, `input`, `Input to send to the server. Example: <rxml>...</rxml>`)

	flag.Parse()

	if *hostPtr == `` {
		flag.Usage()
	}

	fmt.Println(socketclient.GetResponse(*hostPtr, *inputPtr))
}
