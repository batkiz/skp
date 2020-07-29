package lib

import (
	. "github.com/batkiz/skp/utils"
	"io/ioutil"
	"path/filepath"
)

var BucketsDir = filepath.Join(ScoopDir, "buckets")

func GetLocalBuckets() []string {
	dirs, err := ioutil.ReadDir(BucketsDir)
	HandleErr(err)

	var buckets []string
	for _, dir := range dirs {
		if dir.IsDir() {
			buckets = append(buckets, dir.Name())
		}
	}

	return buckets
}


