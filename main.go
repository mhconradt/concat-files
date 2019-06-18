package main

import (
	"io/ioutil"
	"log"
	"os"
)

type Args struct {
	TargetPath string
	SourcePaths []string
}

func ParseArgs() Args {
	target := os.Args[1]
	sources := os.Args[2:]
	return Args{TargetPath: target, SourcePaths: sources }
}

func GetDataFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", path)
	}
	return data
}

func ReadAllFiles(paths []string) [][]byte {
	aggregate := make([][]byte, 0)
	for _, filepath := range paths {
		aggregate = append(aggregate, GetDataFromFile(filepath))
	}
	return aggregate
}

func WriteAllToFile(path string, allData [][]byte) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Error reading file.")
	}
	defer f.Close()
	for _, eachData := range allData {
		n, err := f.Write(eachData)
		if err != nil {
			log.Fatalf("Error writing data to file.")
		}
		log.Printf("%n bytes written to %v", n, path)
	}

	log.Printf("Mission Accomplished! Good night!")
}

func main() {
	args := ParseArgs()
	allFileData := ReadAllFiles(args.SourcePaths)
	WriteAllToFile(args.TargetPath, allFileData)
}
