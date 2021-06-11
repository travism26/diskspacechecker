package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// This is the interface we will build off //
type scanner interface {
	scan() error
	test() string
}

// These are our test objects //
type BasicScanner struct {
	pathToScan string
	outputFile string
}

type LargeFileFinder struct {
	pathToScan  string
	outputFile  string
	minFileSize int64
}

// BASIC SCANNER //
func NewBasicScanner(scanPathIn, outputFileIn string) scanner {
	newBasicScanner := &BasicScanner{
		pathToScan: scanPathIn,
		outputFile: outputFileIn,
	}
	return newBasicScanner
}

func (s *BasicScanner) scan() error {
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
	return err
}

func (s *BasicScanner) test() string {
	fmt.Println("Hello from the BasicScanner side")
	return "Hello world from BasicScanner"
}

// LargeFile SCANNER //
func NewLargeFileFinder(scanPathIn, outputFileIn string, minimumFileSizeIn int64) scanner {
	newLargeFileFinder := &LargeFileFinder{
		pathToScan:  scanPathIn,
		outputFile:  outputFileIn,
		minFileSize: minimumFileSizeIn,
	}
	return newLargeFileFinder
}

func (l *LargeFileFinder) scan() error {
	err := filepath.Walk(l.pathToScan,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Size() >= l.convertToMB(l.minFileSize) {
				fmt.Println(path, info.Size())
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return err
}
func (l *LargeFileFinder) test() string {
	fmt.Println("Hello from the large file finder")
	return "looking for large files"
}

// Quick helper function to convert user input MB into eqivalant (in bytes)
func (l *LargeFileFinder) convertToMB(userIn int64) int64 {
	output := userIn * (1024 * 1024)
	return output
}
