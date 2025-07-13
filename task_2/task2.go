package main

import (
	"strings"
	"unicode"
)
func clean_word(word string) string{
	result:=[]rune{}
	for _, char := range word {
        if !unicode.IsPunct(char) && !unicode.IsSpace(char){
            result = append(result, char) // Append non-punctuation characters
        }
    }
    return strings.ToLower(string(result) )
}

func reverse_word(word string) string{
	result:=make([]rune,len(word))
	for i,char:= range word{
		result[len(word)-1-i]=char
	}
	return string(result)
}

func freq_words(s string) map[string]int{
	words:=strings.Split(s," ")
	freq:=make(map[string]int)
	// fmt.Println(words)
	for _, val:= range words{
		val=clean_word(val)
		freq[val]+=1
	}
	// fmt.Print(freq)
	return freq
}

func plaindrome(s string)bool{
	word:=clean_word(s)
	reversed:=reverse_word(word)
	return strings.EqualFold(word,reversed)

}

// func main(){
// 	s:="i am awel abubekar i am awel awel awel!"
// 	plaindrome(s)
// }