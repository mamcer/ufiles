package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"
)

func calculateHash(filePath string) (string, error) {
	var sha string

	file, err := os.Open(filePath)
	if err != nil {
		return sha, err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return sha, err
	}

	bytes := hash.Sum(nil)[:20]
	sha = hex.EncodeToString(bytes)
	return sha, nil
}

func main() {
	fmt.Println("process started")
	overallStart := time.Now()

	hash, err := calculateHash(os.Args[1])
	if err == nil {
		fmt.Printf("hash:%v\n", hash)
	}

	overallElapsed := time.Since(overallStart)
	fmt.Printf("process finished: %s\n", overallElapsed)
}
