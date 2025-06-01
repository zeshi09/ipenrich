package parser

import (
	"bufio"
	"os"
	"regexp"
)

func ReadingFileForIP(logFile string) ([]string, error) {
	f, err := os.Open(logFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`) // regex for matching IP addresses
	unique := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, ip := range matches {
			unique[ip] = true
		}
	}

	var result []string

	for ip := range unique {
		result = append(result, ip)
	}

	return result, nil
}

func ReadingFileForHashes(logFile string) ([]string, error) {
	f, err := os.Open(logFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	md5Regex := regexp.MustCompile(`\b[a-fA-F0-9]{32}\b`)
	sha1Regex := regexp.MustCompile(`\b[a-fA-F0-9]{40}\b`)
	sha256Regex := regexp.MustCompile(`\b[a-fA-F0-9]{64}\b`)

	unique := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, ip := range matches {
			unique[ip] = true
		}
	}

	var result []string

	for ip := range unique {
		result = append(result, ip)
	}

	return result, nil
}
