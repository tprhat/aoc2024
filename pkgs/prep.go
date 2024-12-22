package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run pkgs/prep.go prep_day <day>    - Prepare directory for puzzle")
		fmt.Println("  go run pkgs/prep.go get_input <day>   - Download input for puzzle")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "prep_day":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run . prep <day>")
			os.Exit(1)
		}
		prepDay(os.Args[2])
	case "get_input":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run . input <day>")
			os.Exit(1)
		}
		getInput(os.Args[2])
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands are:")
		fmt.Println("  go run pkgs/prep.go prep_day <day>    - Prepare directory for puzzle")
		fmt.Println("  go run pkgs/prep.go get_input <day>   - Download input for puzzle")

		os.Exit(1)
	}
}

func getInput(day string) {
	if len(day) == 1 {
		day = "0" + day
	}
	url := "https://adventofcode.com/2024/day/" + day + "/input"
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: os.Getenv("session"),
	}
	req.AddCookie(cookie)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var filePath string
	if dayInt, _ := strconv.Atoi(day); dayInt < 10 {
		day = "0" + day
	}
	filePath = "puzzles/day" + day + "/input.txt"

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	f.Write(body)
	f.Sync()
	f.Close()

	fmt.Printf("Successfully got the puzzle input at %s\n", filePath)
}

func prepDay(day string) {
	// Ensure day is two digits
	if dayInt, _ := strconv.Atoi(day); dayInt < 10 {
		day = "0" + day
	}
	puzzleDir := filepath.Join("puzzles", "day"+day)

	// Create puzzles directory if it doesn't exist
	if err := os.MkdirAll("puzzles", 0755); err != nil {
		fmt.Printf("Error creating puzzles directory: %v\n", err)
		os.Exit(1)
	}

	// Create the puzzle directory
	if err := os.MkdirAll(puzzleDir, 0755); err != nil {
		fmt.Printf("Error creating puzzle directory: %v\n", err)
		os.Exit(1)
	}
	// Check if puzzle directory exists and is not empty
	isEmpty, err := isDirEmpty(puzzleDir)
	if err != nil {
		fmt.Printf("Error checking directory: %v\n", err)
		os.Exit(1)
	}

	if !isEmpty {
		fmt.Printf("Directory %s already exists and is not empty. Skipping...\n", puzzleDir)
		os.Exit(1)
	}
	// Copy files from template directory
	if err := copyDir("template", puzzleDir); err != nil {
		fmt.Printf("Error copying template files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created puzzle directory at %s\n", puzzleDir)
}

func copyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return err
			}
			if err := copyDir(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
func isDirEmpty(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return false, err
	}
	return len(entries) == 0, nil
}
