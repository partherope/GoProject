package td1

import (
	"bufio"
	"os"
)

func IsPalindrome(word string) bool {
	//check if the word is a palindrome
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

func Palindrome(dict []string) []string {
	//return the palindromes in the dictionary
	var palindromes []string
	for i := 0; i < len(dict); i++ {
		if IsPalindrome(dict[i]) {
			palindromes = append(palindromes, dict[i])
		}
	}
	return palindromes
}

func Footprint(s string) (footprint string) {
	//return the new string sorted in the order of the alphabet
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] < s[j] {
				s = s[:i] + string(s[j]) + s[i+1:j] + string(s[i]) + s[j+1:]
			}
		}
	}
	return s
}

func DicFromFile(filename string) (dic []string) {
	//read the file line by line and store the words in a list
	dic = make([]string, 0)
	file, _ := os.Open(filename)
	defer file.Close()
	if file != nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			dic = append(dic, scanner.Text())
		}
	}
	return dic
}
