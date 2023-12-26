package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
)

const proverbURL = "https://go-proverbs.github.io"

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		proverb, err := getRandomProverb()
		if err != nil {
			return
		}

		_, err = fmt.Fprintln(conn, proverb)
		if err != nil {
			return
		}

		time.Sleep(3 * time.Second)
	}
}

func getRandomProverb() (string, error) {
	resp, err := http.Get(proverbURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var proverbs []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		proverbs = append(proverbs, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(proverbs))
	return proverbs[randomIndex], nil
}
