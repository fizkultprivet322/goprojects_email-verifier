package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}
