package utils

import (
	"flag"
	"fmt"
	"strings"
	"os"
)


type ScanCtx struct {
	Target	Page
	Level	uint8
}

type Page struct {
	BaseURL		string
	Path		string
	Scanned		bool
}

func ParseArgs(flags *flag.FlagSet) (*ScanCtx, error) {
	fullTarget := flags.String("t", "", "Provide the IP of the target (e.g., 192.168.56.121)")
	scanLevel := flags.Int("l", 2, "The -l option defines how many levels deep you want to scan")

	err := flags.Parse(os.Args[1:])

	if err != nil {
		return nil, fmt.Errorf("error parsing arguments: %v", err)
	}

	if len(os.Args) < 2 {
		flags.Usage()
		return nil, nil
	}

	tarStr := strings.TrimPrefix(*fullTarget, "http://")
	splits := strings.SplitN(tarStr, "/", 2)

	base := splits[0]

	var path string
	if len(splits) >= 2 {
		path = splits[1]
	}

	target := Page {
		BaseURL:	base,
		Path:		path,
		Scanned: 	false,
	}

	return &ScanCtx{
		Target: 	target,
		Level:		uint8(*scanLevel),
	}, nil

}
