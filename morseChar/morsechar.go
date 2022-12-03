package morsechar

import (
	"errors"
	"strconv"
	"strings"
)

/*
alphabets raw data for morse code in map
*/
var AlphabetMorse = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",

	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
	' ': " ",

	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",
}

/*
numbers raw data in map
*/
var NumMorse = map[int]string{
	1: ".----",
	2: "..---",
	3: "...--",
	4: "....-",
	5: ".....",
	6: "-....",
	7: "--...",
	8: "---..",
	9: "----.",
	0: "-----",
}

/*
NumTOMorseInt function takes an argument in integer type
and it converts into single string.
*/
func NumTOMorseInt(Nums int) (outStr string) {
	var bindInt []string
	numsStr := strconv.Itoa(Nums)
	newNumsStr := strings.TrimLeft(numsStr, "0")
	str := []rune(newNumsStr)
	for i := 0; i < len(str); i++ {
		key, _ := strconv.Atoi(string(str[i]))
		bindInt = append(bindInt, NumMorse[key]+" ")
	}
	out := strings.Join(bindInt, "")
	return out
}

/*
AToMorseStr function takes an argument in rune type
and it converts into single character string.
*/
func AToMorseStr(Char rune) string {
	str := AlphabetMorse[Char]
	return str
}

/*
WordToMorseWord function takes an argument in string as single word  type
and it converts into single word string.
*/
func WordToMorseWord(singleWord string) (outPut string, err error) {
	wordChar := []rune(singleWord)
	var str []string
	//checking space or tab in word
	if strings.Contains(singleWord, " ") || strings.Contains(singleWord, "	") {
		return "", errors.New("word should be without any space")
	}
	for i := 0; i < len(wordChar); i++ {
		key := wordChar[i]
		str = append(str, AlphabetMorse[key]+" ")
	}
	outStr := strings.Join(str, " ")
	return outStr, nil
}

/*
WordToMorseWord function takes an argument in string as single word  type
and it converts into single word string.
*/
func StrToMorseStr(s string) (string, error) {
	str := strings.Split(s, " ")
	var strArr []string
	for i := 0; i < len(str); i++ {
		outSt, err := WordToMorseWord(str[i])
		if err != nil {

			return "", err
		}
		strArr = append(strArr, outSt+"/")
	}
	finalStr := strings.Join(strArr, "")
	return finalStr, nil
}
