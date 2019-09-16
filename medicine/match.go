package medicine

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func Existance(inputText string) bool {
	pharmacyKnowledge := Pharmacy()
	toLower := strings.ToLower(inputText)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	noAccent, _, _ := transform.String(t, toLower)
	words := strings.Split(noAccent, " ")
	for _, word := range words {
		for _, medicine := range pharmacyKnowledge {
			if word == medicine {
				return true
			}
		}
	}
	return false
}

func GetInfo(med string) Medicine {
	toLower := strings.ToLower(med)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	medNoAccent, _, _ := transform.String(t, toLower)
	words := strings.Split(medNoAccent, " ")
	for _, word := range words {
		if word == "ibuprofeno" {
			return NewIbuprofeno()
		}
		if word == "loperamida" {
			return NewLoperamida()
		}
	}
	var medEmpty Medicine
	return medEmpty
}
