package parser

import (
	"bufio"
	"os"
	"regexp"
)

func ReadingFile(logFile string) ([]string, error) {
	f, err := os.Open(logFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
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
