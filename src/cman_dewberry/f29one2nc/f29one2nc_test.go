package main

import (
	"os"
	"path/filepath"
	"testing"
)

/* func init() {
	os.Mkdir("./nc_f291_Sonny_script", 0777)
} */

func TestFindF291Files(t *testing.T) {
	result := findF291Files("./testfiles/")
	// 4 because it counts the dir as well
	if len(result) != 4 {
		t.Error("Got ", result)
	}
}

func TestF2912nc(t *testing.T) {
	testFiles := []string{
		"./testfiles/ptgc1_200307",
		"./testfiles/42002_198512",
		"./testfiles/46070_200909",
	}
	f2912nc(testFiles)
	for _, file := range testFiles {
		file, err := os.Stat(filepath.Base(file) + ".nc")
		if err != nil {
			t.Error("File not found, got", file)
		}
	}
}
