package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	// "regexp"

	// "net/netip"
	"flag"
)

type ScanCtx struct {
	Level	uint8
}

type Page struct {
	BaseURL		string
	Path		string
	Scanned		bool
}

func main() {
	fmt.Println("Hello world!")
	
	var fullTarget string
	flag.StringVar(&fullTarget, "t", "", "Provide the IP of the target (e.g., 192.168.56.121)")
	levelInt := flag.Int("l", 2, "The -l option defines how many levels deep you want to scan")
	flag.Parse()

	tarStr := strings.TrimPrefix(fullTarget, "http://")
	splits := strings.SplitN(tarStr, "/", 2)
	fmt.Println(splits)
	base := splits[0]
	path := splits[1]
	fmt.Println(tarStr)
	fmt.Println("base: ", base)
	fmt.Println("path: ", path)

	fmt.Println(*levelInt)
	target := fmt.Sprintf("http://%s", tarStr)

	resp, err := http.Get(target)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)
	bodyStr := string(bodyByte[:])
	lines := strings.Split(bodyStr, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}
