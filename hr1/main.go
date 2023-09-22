package main

import "fmt"

func main() {
	words := []string{"the-stealth-warrior", "The_Stealth_Warrior", "the_stealth_warrior"}
	wordsToCipher := []string{"middle-Outz", "Always-Look-on-the-Bright-Side-of-Life", "O	www.hackerrank.com", "159357lcfd", "abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "1593576382", "!@#$%^&*()_+-=", "Hello_World!"}
	for _, word := range words {
		w, c := toCamelCase(word)
		fmt.Println(word + " => " + w + " (" + fmt.Sprint(c) + ")")
	}
	for _, word := range wordsToCipher {
		fmt.Println(word + " => " + caesarCipher(word, 2))
	}
}

func caesarCipher(s string, k int32) string {
	output := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			output += string(((int32(s[i]) - 'a' + k) % 26) + 'a')
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			output += string(((int32(s[i]) - 'A' + k) % 26) + 'A')
		} else {
			output += string(s[i])
		}
	}
	return output
}

func toCamelCase(s string) (string, int) {
	output := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '_' {
			i++
			count++
			output += string(s[i] - 32)
		} else {
			output += string(s[i])
		}
	}
	return output, count
}
