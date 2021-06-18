package cmd

import (
	scanobjects "diskspacecheck/cmd/scanObjects"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// This is the interface we will build off //
type scanner interface {
	scan() ([]scanobjects.ScanObject, error)
	test() string
}

// These are our test objects //
type BasicScanner struct {
	pathToScan string
	outputFile string
}

type LargeFileFinder struct {
	pathToScan   string
	outputFile   string
	minFileSize  int64
	scannedFiles []scanobjects.ScanObject
}

// BASIC SCANNER //
func NewBasicScanner(scanPathIn, outputFileIn string) scanner {
	newBasicScanner := &BasicScanner{
		pathToScan: scanPathIn,
		outputFile: outputFileIn,
	}
	return newBasicScanner
}

func (s *BasicScanner) scan() ([]scanobjects.ScanObject, error) {
	err := filepath.Walk(s.pathToScan,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return nil, err
}

func (s *BasicScanner) test() string {
	fmt.Println("Hello from the BasicScanner side")
	return "Hello world from BasicScanner"
}

// methods //
func scan(path string) ([]scanobjects.ScanObject, error) {
	var results []scanobjects.ScanObject
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path, info.Size())
			results = append(results, scanobjects.ScanObject{Path: path, FileInfo: info})
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return results, err
}
