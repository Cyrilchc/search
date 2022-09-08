package main

import (
	"flag"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"strings"
)

var searchedFiles = 0

func main() {
	tm.Clear()
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	// flags
	caseSensitive := flag.Bool("c", false, "Using this flag will search with case sensitivity enabled")
	recursive := flag.Bool("r", false, "Using this flag will search element on folder and on all of its sub-folders")
	searchString := flag.String("s", "", "The specified search string will be used to search files and directories")
	displayFiles := flag.Bool("f", false, "Using this flag will search files")
	displayDirectories := flag.Bool("d", false, "Using this flag will search directories")
	path := flag.String("p", currentPath, "Using this flag will search in the provided path")
	flag.Parse()

	if *path != "" {
		pathExist, _ := pathExists(*path)
		if !pathExist {
			fmt.Println("please provide a valid directory path")
			os.Exit(1)
		} else {
			currentPath = *path
		}
	}

	// Get directory elements
	var folderElements []os.DirEntry
	err = getDirectoryContent(currentPath, *recursive, *displayFiles, *displayDirectories, &folderElements)
	if err != nil {
		fmt.Println(err.Error())
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

func getDirectoryContent(path string, recursive bool, displayFiles bool, displayDirectories bool, elements *[]os.DirEntry) error {
	directoryContents, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, element := range directoryContents {
		searchedFiles++
		tm.MoveCursor(1, 1)
		_, err := tm.Printf("Searched elements : %v", searchedFiles)
		if err != nil {
			return err
		}
		tm.Flush()
		if !displayFiles && !displayDirectories {
			*elements = append(*elements, element)
		} else {
			if displayFiles {
				if !element.IsDir() {
					*elements = append(*elements, element)
				}
			}
			if displayDirectories {
				if element.IsDir() {
					*elements = append(*elements, element)
				}
			}
		}

		if element.IsDir() && recursive {
			err := getDirectoryContent(filepath.Join(path, element.Name()), recursive, displayFiles, displayDirectories, elements)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
