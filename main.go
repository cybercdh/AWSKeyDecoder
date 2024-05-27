/*
read in a list of AWSKeyIds from stdin and converts them to AWS Account Ids
*/
package main

import (
	"bufio"
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var verbose bool
var concurrency int
var keys = make(chan string, 100)

func AWSAccountFromAWSKeyID(AWSKeyID string) (int64, error) {
	trimmedAWSKeyID := AWSKeyID[4:]

	base32Decoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	x, err := base32Decoder.DecodeString(strings.ToUpper(trimmedAWSKeyID))
	if err != nil {
		return 0, err
	}

	y := x[:6]

	z := binary.BigEndian.Uint64(append(make([]byte, 2), y...))

	mask, err := hex.DecodeString("7fffffffff80")
	if err != nil {
		return 0, err
	}
	maskInt := binary.BigEndian.Uint64(append(make([]byte, 2), mask...))

	e := (z & maskInt) >> 7
	return int64(e), nil
}

func processAWSKeyID(AWSKeyID string) {
	accountID, err := AWSAccountFromAWSKeyID(AWSKeyID)
	if err != nil {
		if verbose {
			fmt.Println("Error processing %s: %v", AWSKeyID, err)
		}
		return
	}

	result := ""

	if verbose {
		result = fmt.Sprintf("AWS Key ID: %s -> Account ID: %012d", AWSKeyID, accountID)
	} else {
		result = fmt.Sprintf("%012d", accountID)
	}

	fmt.Println(result)
}

func main() {

	flag.IntVar(&concurrency, "c", 20, "set the concurrency level")
	flag.BoolVar(&verbose, "v", false, "Get more info on attempts")

	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for key := range keys {
				processAWSKeyID(key)
			}
		}()
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		AWSKeyID := scanner.Text()
		keys <- AWSKeyID
	}

	close(keys)
	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
