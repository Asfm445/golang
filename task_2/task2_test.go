package main

import (
	"testing"
)

func TestRemovePunctuation(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello!", "hello"},
        {"hello, world!", "helloworld"},
        {"goodbye!!!", "goodbye"},
        {"1234?5678.", "12345678"},
        {"no punctuation", "nopunctuation"},
        {"", ""},
        {"awel@#12s","awel12s"},
    }

    for _, test := range tests {
        result := clean_word(test.input)
        if result != test.expected {
            t.Errorf("clean_word(%q) = %q; want %q", test.input, result, test.expected)
        }
    }
}

func mapsEqual(a, b map[string]int) bool {
    if len(a) != len(b) {
        return false
    }
    for k, v := range a {
        if b[k] != v {
            return false
        }
    }
    return true
}




func TestFreqWords(t *testing.T){
    tests := []struct {
        input    string
        expected map[string]int
    }{
        {"i am awel abubekar i am awel awel awel!", map[string]int{"i":2,"am":2,"awel":4,"abubekar":1}},
        
    }
    for _, test := range tests {
        result := freq_words(test.input)
        if !mapsEqual(result, test.expected) {
            t.Errorf("freq_words(%q) = %q; want %q", test.input, result, test.expected)
        }
    }
}

func TestReverseWord(t *testing.T){
    tests := []struct {
        input    string
        expected string
    }{
        {"awel", "lewa"},
        {"did","did"},
        {"me","em"},
        
    }
    for _, test := range tests {
        result := reverse_word(test.input)
        if result!=test.expected {
            t.Errorf("reverse_word(%q) = %q; want %v", test.input, result, test.expected)
        }
    }
}

func TestPalaindrome(t *testing.T){
    tests := []struct {
        input    string
        expected bool
    }{
        {"awel", false},
        {"did",true},
        {"d id",true},
        {"did!",true},
        
    }
    for _, test := range tests {
        result := plaindrome(test.input)
        if result!=test.expected {
            t.Errorf("palindrome(%q) = %v; want %v", test.input, result, test.expected)
        }
    }
}
