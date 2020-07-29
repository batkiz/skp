package commands

import (
	"fmt"
	"github.com/batkiz/skp/lib"
)

func Bucket(args []string) {
	if args[0] == "list" {
		buckets := lib.GetLocalBuckets()
		for _, bucket := range buckets {
			fmt.Println(bucket)
		}
	}
}
