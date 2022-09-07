package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// flags
	caseSensitive := flag.Bool("c", false, "Using this flag will search with case sensitivity enabled")
	recursive := flag.Bool("r", false, "Using this flag will search element on folder and on all of its sub-folders")
	searchString := flag.String("s", "", "The program will use the specified search string")
	flag.Parse()

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	// Get directory elements
	var folderElements []os.DirEntry
	err = getDirectoryContent(currentPath, *recursive, &folderElements)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	// Search
	var searchedElements []os.DirEntry
	if *searchString != "" {
		for _, element := range folderElements {
			elementStr := ""
			if *caseSensitive {
				elementStr = element.Name()
			} else {
				elementStr = strings.ToLower(element.Name())
			}
			if strings.Contains(elementStr, *searchString) {
				searchedElements = append(searchedElements, element)
			}
		}
	} else {
		for _, element := range folderElements {
			searchedElements = append(searchedElements, element)
		}
	}

	// Printing
	for _, element := range searchedElements {
		if element.IsDir() {
			color.Cyan(element.Name())
		} else {
			color.White(element.Name())
		}
	}
}

func getDirectoryContent(path string, recursive bool, elements *[]os.DirEntry) error {
	directoryContents, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, element := range directoryContents {
		*elements = append(*elements, element)
		if element.IsDir() && recursive {
			getDirectoryContent(filepath.Join(path, element.Name()), recursive, elements)
		}
	}

	return nil
}
