package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
	"github.com/goscan/utils"
)


func main() {
	fmt.Println("Hello world!")
		
	scanCtx, err := utils.ParseArgs(flag.NewFlagSet("main", flag.ExitOnError))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	if scanCtx == nil {
		os.Exit(1)
	}

	fmt.Printf("BaseURL: %s\n", scanCtx.Target.BaseURL)
	fmt.Printf("Path: %s\n", scanCtx.Target.Path)
	fmt.Printf("Scanned? %t\n", scanCtx.Target.Scanned)
	fmt.Printf("Scan level: %d\n", scanCtx.Level)

	target := fmt.Sprintf("http://%s", scanCtx.Target.BaseURL)

	resp, err := http.Get(target)
	if err != nil {
		fmt.Printf("Error with HTTP GET request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	htmlPage, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", htmlPage)
	traverse(htmlPage)

	links := findLinks(htmlPage, nil)
	fmt.Printf("Found links: %+v\n", links)
}

func traverse(n *html.Node) {

	fmt.Printf("Node Type: %+v - Node Data: %+v - Node Attributes: %+v\n", n.Type, n.Data, n.Attr)

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		fmt.Printf("First child: %+v\n", child.Type)
		traverse(child)
	}
}

func findLinks(n *html.Node, links []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		links = findLinks(child, links)
	}

	return links
}
