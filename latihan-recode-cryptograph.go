package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var Alfabet = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
	"I": 8,
	"J": 9,
	"K": 10,
	"L": 11,
	"M": 12,
	"N": 13,
	"O": 14,
	"P": 15,
	"Q": 16,
	"R": 17,
	"S": 18,
	"T": 19,
	"U": 20,
	"V": 21,
	"W": 22,
	"X": 23,
	"Y": 24,
	"Z": 25,
}

type Kriptografi struct {
	EncryptedPesan string
	Key            int
}

func mapkey(m map[string]int, value int) (key string) {
	for k, v := range m {
		if v == value {
			key = k
			return
		}
	}
	return
}

func readIndexToString(slc []int) string {
	var myString []string
	for _, v := range slc {
		indexOfAlphabet := v % 26
		if indexOfAlphabet < 0 {
			myString = append(myString, mapkey(Alfabet, 26-indexOfAlphabet))
		}
		myString = append(myString, mapkey(Alfabet, indexOfAlphabet))
	}
	return strings.Join(myString, "")
}

func (c *Kriptografi) Decrypt() string {
	takeSliceStr := strings.Split(c.EncryptedPesan, "")
	var myString []string
	for _, v := range takeSliceStr {
		penambahanIndex := c.Key % 26
		indexAsli := Alfabet[v] - penambahanIndex
		if indexAsli < 0 {
			myString = append(myString, mapkey(Alfabet, 26+indexAsli))
		}
		myString = append(myString, mapkey(Alfabet, indexAsli))
	}
	return strings.Join(myString, "")
}

func Cryptograph(pesan string, key int) (*Kriptografi, error) {
	var encryptedPesanIndex []int
	for _, v := range pesan {
		if unicode.IsLetter(v) {
			encryptedPesanIndex = append(encryptedPesanIndex, Alfabet[strings.ToUpper(string(v))]+key)
		} else {
			return &Kriptografi{EncryptedPesan: "", Key: key}, errors.New("non alphabetic string")
		}
	}
	var encryptedPesan string = readIndexToString(encryptedPesanIndex)
	return &Kriptografi{EncryptedPesan: encryptedPesan, Key: key}, nil
}

func main() {
	crypt, _ := Cryptograph("FadhliAbrar", 135)

	fmt.Println(crypt)
}
