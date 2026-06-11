// compilation: go build -o MapNTLMHashes MapNTLMHashes.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func mapHashes(potFilePath, dumpFilePath string) int {
	potLines, err := readLines(potFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening pot file: %v\n", err)
		os.Exit(1)
	}

	dumpLines, err := readLines(dumpFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening dump file: %v\n", err)
		os.Exit(1)
	}

	// Build hash->password map from pot file
	hashMap := make(map[string]string)
	for _, entry := range potLines {
		pos := strings.Index(entry, ":")
		if pos != -1 {
			hash := entry[:pos]
			password := entry[pos+1:]
			hashMap[hash] = password
		}
	}

	// Match NTLM hashes from dump file against hashMap
	counter := 0
	for _, entry := range dumpLines {
		fields := strings.Split(entry, ":")
		if len(fields) < 4 {
			continue
		}
		ntlmHash := fields[3]
		if password, found := hashMap[ntlmHash]; found {
			fmt.Printf("%s:%s\n", entry, password)
			counter++
		}
	}

	return counter
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("[+] Usage: ./MapNTLMHashes potfile NTDS-DUMP")
		os.Exit(1)
	}

	total := mapHashes(os.Args[1], os.Args[2])
	fmt.Printf("[+] Total Matches: %d\n", total)
}
