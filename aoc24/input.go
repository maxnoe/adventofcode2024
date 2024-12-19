package aoc24

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func hash(session string) string {
	hash := sha256.Sum256([]byte(session))
	return hex.EncodeToString(hash[:])
}

func cachePath(day int, session string) string {
	return fmt.Sprintf(".input_cache/%s/%d.input", hash(session), day)
}

func cacheStore(day int, session string, input string) {
	p := cachePath(day, session)
	os.MkdirAll(filepath.Dir(p), 0o755)
	err := os.WriteFile(p, []byte(input), 0o644)
	if err != nil {
		log.Printf("Error storing input in cache: %v", err)
	} else {
		log.Printf("Successfully stored input in cache")
	}
}

func isFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func cacheLoad(day int, session string) (string, bool) {
	p := cachePath(day, session)
	if !isFile(p) {
		return "", false
	}
	input, err := os.ReadFile(p)
	if err != nil {
		log.Printf("Error loading input from cache: %v", err)
		return "", false
	}

	log.Printf("Successfully loaded input from cache")
	return string(input), true
}

func GetInput(day int) (string, error) {

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", errors.New("You need to set the AOC_SESSION to your session cookie.")
	}

	input, success := cacheLoad(day, session)
	if success {
		return input, nil
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Received status code: %d", res.StatusCode))
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	input = string(data)
	cacheStore(day, session, input)
	return input, nil
}
