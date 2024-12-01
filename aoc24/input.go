package aoc24

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetInput(day int) (string, error) {

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", errors.New("You need to set the AOC_SESSION to your session cookie.")
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
	return string(data), nil
}
