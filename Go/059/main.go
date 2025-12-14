package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isReadableASCIIText(text string) bool {
	for b := range text {
		if text[b] < 32 || text[b] >= 126 {
			return false
		}

		if text[b] == 36 || text[b] == 64 || text[b] == 94 ||
			text[b] == 95 || text[b] == 96 {
			return false
		}
	}

	return true
}

func decrypt(key, text []byte) []byte {
	decrypted := make([]byte, len(text))

	for i, c := range text {
		decrypted[i] = c ^ key[i%len(key)]
	}

	return decrypted
}

func findValidKey(text []byte) []byte {
	for a := byte(97); a <= 122; a++ {
		for b := byte(97); b <= 122; b++ {
			for c := byte(97); c <= 122; c++ {
				decrypted_text := decrypt([]byte{a, b, c}, text)

				if isReadableASCIIText(string(decrypted_text)) {
					return []byte{a, b, c}
				}
			}
		}
	}

	return []byte{}
}

func main() {
	answer := 0

	file, err := os.Open("0059_cipher.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("File is empty!")
	}

	codes := strings.Split(scanner.Text(), ",")
	encrypted := make([]byte, len(codes))

	for i, p := range codes {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			log.Fatal(err)
		}
		encrypted[i] = byte(n)
	}

	key := findValidKey(encrypted)
	decrypted_text := decrypt(key, encrypted)

	for _, c := range decrypted_text {
		answer += int(c)
	}

	fmt.Println("Answer:", answer)
}
