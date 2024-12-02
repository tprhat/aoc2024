package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getInput(day string) {
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
}

func main() {

	flag.Parse()
	args := flag.Args()
	day := args[0]
	getInput(day)
}
