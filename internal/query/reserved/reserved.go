package reserved

import "strings"

// Reserved words
const (
	AS     = "AS"
	FROM   = "FROM"
	SELECT = "SELECT"
	WHERE  = "WHERE"
)

var reserved = makeReservedMap([]string{
	AS, FROM, SELECT, WHERE,
})

// IsReserved returns whether or not the provided word is reserved
func IsReserved(word string) (string, bool) {
	upper := strings.ToUpper(word)
	if res, ok := reserved[upper]; ok && res {
		return upper, true
	}
	return word, false
}

func makeReservedMap(words []string) map[string]bool {
	res := make(map[string]bool, len(words))
	for _, word := range words {
		res[strings.ToUpper(word)] = true
	}
	return res
}
