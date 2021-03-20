package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("abc(defgh)adsa"))
	//fmt.Println(findFirstStringInBracket("ad)dasd(dasdsad"))
	//fmt.Println(findFirstStringInBracket("(adasd(dasdas))"))
	//fmt.Println(findFirstStringInBracketOriginal("(adasd(dasdas))"))
}

//findFirstStringInBracket is refactored code
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexOpeningBracket := strings.Index(str,"(")
		indexClosingBracket := strings.Index(str,")")

		if (indexClosingBracket > indexOpeningBracket) && (indexOpeningBracket >=0 && indexClosingBracket >= 0) {
			runes := []rune(str)
			return string(runes[indexOpeningBracket+1:indexClosingBracket])
		}
	}
	return ""
}







//findFirstStringInBracketOriginal is original code from the test question
func findFirstStringInBracketOriginal(str string) string {
	if (len(str) > 0) {
		indexFirstBracketFound := strings.Index(str,"(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound-1])
			}else{
				return ""
			}
		}else{
			return ""
		}
	}else{
		return ""
	}
	return ""
}