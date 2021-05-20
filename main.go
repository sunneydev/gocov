package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/tools/cover"
)

func main() {
	if len(os.Args) < 2 {
		panic("coverprofile path is required!")
	}
	fn := os.Args[1]
	if _, err := os.Stat(fn); err != nil {
		panic(fn + " does not exist")
	}
	profiles, err := getProfiles(fn)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total coverage: %v\n", calculate(profiles))
	os.Remove(fn)
}

func getProfiles(path string) (profiles []*cover.Profile, err error) {
	return cover.ParseProfiles(path)
}

func calculate(profiles []*cover.Profile) string {
	var (
		statements int
		covered    int
	)

	for _, profile := range profiles {
		for _, block := range profile.Blocks {
			statements += block.NumStmt
			if block.Count > 0 {
				covered += block.NumStmt
			}
		}
	}

	return strconv.Itoa(int((float64(covered)/float64(statements))*100.0)) + "%"
}
