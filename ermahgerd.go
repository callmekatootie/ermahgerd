package ermahgerd

import (
    "fmt"
    "regexp"
    "strings"
)

const beginNotWords string = `^\W+`
const endNotWords string = `\W+$`

func replace(regex, replaceWith string, s *string) {
    r := regexp.MustCompile(regex)

    *s = r.ReplaceAllString(*s, replaceWith)
}

func removeDuplicates(word string) string {
    var last rune

    return strings.Map(func(r rune) rune {
        if r != last {
            last = r

            return r
        }

        return -1
    }, word)
}

func parse(word string) string {
    // Word is too short to translate
    if len(word) < 2 {
        return word
    }

    // Common words that already have a direct translation
    switch word {
    case "AWESOME":
        return "ERSUM"
    case "BANANA":
        return "BERNERNER"
    case "BAYOU":
        return "BERU"
    case "FAVORITE", "FAVOURITE":
        return "FRAVRIT"
    case "GOOSEBUMPS":
        return "GERSBERMS"
    case "LONG":
        return "LERNG"
    case "MY":
        return "MAH"
    case "THE":
        return "DA"
    case "THEY":
        return "DEY"
    case "WE'RE":
        return "WER"
    case "YOU":
        return "U"
    case "YOU'RE":
        return "YER"
    }

    original := word

    // Remove vowels that occur at the end of the word
    // Only for words whose length is greater than 2 to prevent single character words
    if len(original) > 2 {
        replace(`[AEIOU]$`, "", &word)
    }

    // Reduce duplicate letters
    word = removeDuplicates(word)

    // Reduce consecutive vowels (and Y) to just one
    replace(`[AEIOUY]{2,}`, "E", &word)

    // Retain a single 'Y'
    replace(`Y{2,}`, "Y", &word)

    // DOWN -> DERN
    replace(`OW`, "ER", &word)

    // PANCAKES -> PERNKERKS
    replace(`AKES`, "ERKS", &word)

    // Replace vowels (and Y) with ER
    replace(`[AEIOUY]`, "ER", &word)

    // Other conversions that are direct
    replace(`ERH`, "ER", &word)

    replace(`MER`, "MAH", &word)

    replace(`ERNG`, "IN", &word)

    replace(`ERPERD`, "ERPED", &word)

    replace(`MAHM`, "MERM", &word)

    // If the word begins with Y, retain it
    if original[0] == 89 {
        word = "Y" + word
    }

    // Reduce any duplicate letters (again)
    word = removeDuplicates(word)

    r := regexp.MustCompile(`LOW$`)
    l := regexp.MustCompile(`LER$`)

    if r.MatchString(original) == true && l.MatchString(word) {
        replace(`LER`, "LO", &word)
    }

    return word
}

func Gert(sentence string) string {
    var translatedWords []string

    prefix := regexp.MustCompile(beginNotWords)
    suffix := regexp.MustCompile(endNotWords)

    sentence = strings.ToUpper(sentence)

    words := strings.Split(sentence, " ")

    for _, word := range words {
        wordCopy := word

        replace(beginNotWords, "", &word)
        replace(endNotWords, "", &word)

        if len(word) > 0 {
            beginString := prefix.FindAllString(wordCopy, 1)
            endString := suffix.FindAllString(wordCopy, 1)

            word = parse(word)

            if beginString != nil {
                word = beginString[0] + word
            }

            if endString != nil {
                word = word + endString[0]
            }
        } else {
            word = parse(word)
        }

        translatedWords = append(translatedWords, word)
    }

    return strings.Join(translatedWords, " ")
}
