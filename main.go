package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])

	leftOverLetters := string(runes[0 : len(letter)-key])
	return fmt.Sprintf("%s%s", lastLetterKey, leftOverLetters)

}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString strings.Builder
	findOne := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			//hashedString = hashedString + string(hashLetter[letterPosition])
			hashedString.WriteString(string(hashLetter[letterPosition]))
			return r
		}
		return r

	}
	_ = strings.Map(findOne, plainText)
	return hashedString.String()

}

func decrypt(key int, encryptText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var decryptedString strings.Builder
	findOne := func(r rune) rune {
		pos := strings.Index(hashLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			//hashedString = hashedString + string(originalLetter[letterPosition])
			decryptedString.WriteString(string(originalLetter[letterPosition]))
		}
		return r
	}
	_ = strings.Map(findOne, encryptText)
	//return hashedString
	return decryptedString.String()
}

func main() {

	plainText := "HELLOWORLD"
	fmt.Println("Plain Text", plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted Text", encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted Text", decrypted)

}
