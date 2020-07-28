package util

import (
	"bufio"
	"github.com/go-errors/errors"
	"log"
	"os"
	"strings"
)

// FileExists checks if path exists and is a file
func FileExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return !fileInfo.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errors.Errorf("failed to check if file exists %s", path)
}

// returns true/false if a file exists, or crashes on error
func FileMustExist(path string) bool {
	ex, err := FileExists(path)
	if err != nil {
		log.Panicf("file exists failed: %s", err)
		return false
	}
	return ex
}

// DirExists checks if path exists and is a directory
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DoesFileContainString(s string, pathToFile string) (bool, int, error) {
	replacer := strings.NewReplacer("~", os.Getenv("HOME"))
	pathToFile = replacer.Replace(pathToFile)
	Logger().Debugf("Looking for text : %s", s)
	Logger().Debugf("In File          : %s", pathToFile)

	f, err := os.Open(pathToFile)
	if err != nil {
		errors.Errorf("Error Opening file: %s, Error: %s", pathToFile, err)
		return false, -1, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			return true, line, nil
		}
		line++
	}
	if someError := scanner.Err(); someError != nil {
		return false, -1, err
	}
	return false, -1, nil
}

func FindMatchesInFile(s string, pathToFile string) ([]int, error) {
	replacer := strings.NewReplacer("~", os.Getenv("HOME"))
	pathToFile = replacer.Replace(pathToFile)
	Logger().Debugf("Looking for text : %s", s)
	Logger().Debugf("In File          : %s", pathToFile)
	var listLineNumbers []int
	f, err := os.Open(pathToFile)
	if err != nil {
		errors.Errorf("Error Opening file: %s, Error: %s", pathToFile, err)
		return listLineNumbers, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			listLineNumbers = append(listLineNumbers, line)
		}
		line++
	}
	if someError := scanner.Err(); someError != nil {
		return listLineNumbers, err
	}
	Logger().Debugf("Found %s", listLineNumbers)
	return listLineNumbers, nil
}

// StripTrailingSlash removes any trailing forward slashes on the URL
func StripTrailingSlash(url string) string {
	if url[len(url)-1:] == "/" {
		return url[0 : len(url)-1]
	}
	return url
}
