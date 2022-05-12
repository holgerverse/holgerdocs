package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

/*
	Return all files in a directory specified by folderPath that match the fileSuffix
*/
func filesInDirectory(folderPath string, fileSuffix string) []fs.FileInfo {
	// Stores all files that match the fileSuffix
	var filteredFiles []fs.FileInfo

	unfilteredFiles, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	// Safe all file names, that match the fileSuffix
	for _, v := range unfilteredFiles {
		if strings.HasSuffix(v.Name(), fileSuffix) {
			filteredFiles = append(filteredFiles, v)
		}
	}

	return filteredFiles
}

/*
	Return the absolute path of the path specified by path
*/
func absolutePath(path string) string {
	var absolutePath string

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	return absolutePath
}
