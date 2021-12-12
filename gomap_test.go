package gomap_test

import (
	"fmt"
	"testing"

	"github.com/JustinTimperio/gomap"
)

func TestScanIP(t *testing.T) {
	results, err := gomap.ScanIP("127.0.0.1", "tcp", false, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range results.Results {
		fmt.Println(v.Port)
	}
}
