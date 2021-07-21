package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	ns, err := net.LookupHost("www.google.com.tw")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "--%s\n", n)
	}
}
