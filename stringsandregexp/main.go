package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func main() {
	product := "Kayak"
	fmt.Println("Product:", product)
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))

	price := "€100"
	fmt.Println("Strings Prefix:", strings.HasPrefix(price, "€"))
	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price), []byte{226, 130}))

	description := "A boat for sailing"
	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.Title(description))
	fmt.Println("Title:", strings.ToTitle(description))

	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))

	for _, char := range product {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
	}

	description = "A boat for one person"
	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))

	splits := strings.Split(description, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}

	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}

	splits = strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split N>>" + x + "<<")
	}

	splitsAfter = strings.SplitAfterN(description, " ", 3)
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter N>>" + x + "<<")
	}

	description = "This  is  double  spaced"
	splits = strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split N double spaced>>" + x + "<<")
	}

	splits = strings.Fields(description)
	for _, x := range splits {
		fmt.Println("Fields double spaced>>" + x + "<<")
	}

	splitter := func(r rune) bool {
		return r == ' '
	}
	splits = strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("FieldsFunc double spaced>>" + x + "<<")
	}

	username := " Alice  "
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed (TrimSpace):", ">>"+trimmed+"<<")

	description = "A boat for one person"
	trimmed = strings.Trim(description, "Asnob ")
	fmt.Println("Trimmed (Trim):", trimmed)

	trimmed = strings.TrimLeft(description, "Asnob ")
	fmt.Println("Trimmed (TrimLeft):", trimmed)

	trimmed = strings.TrimRight(description, "Asnob ")
	fmt.Println("Trimmed (TrimRight):", trimmed)

	prefixTrimmed := strings.TrimPrefix(description, "A boat ")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	fmt.Println("Trimmed (TrimPrefix):", prefixTrimmed)
	fmt.Println("Not trimmed (TrimPrefix):", wrongPrefix)

	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimmed = strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed (TrimFunc):", trimmed)

	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "canoe")
	fmt.Println("Replace:", replace)
	fmt.Println("ReplaceAll:", replaceAll)

	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}
	mapped := strings.Map(mapper, text)
	fmt.Println("Mapped:", mapped)

	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	replaced := replacer.Replace(text)
	fmt.Println("Replaced:", replaced)

	elements := strings.Fields(text)
	joined := strings.Join(elements, "--")
	fmt.Println("Joined:", joined)

	var builder strings.Builder
	for _, sub := range strings.Fields(text) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	fmt.Println("String:", builder.String())

	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}

	pattern, compileErr := regexp.Compile("[A-z]oat")
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}

	pattern = regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description = "Kayak. A boat for one person."
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1], "=", getSubstring(description, firstIndex))
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-", idx[1], "=", getSubstring(description, idx))
	}

	firstMatch := pattern.FindString(description)
	allMatches := pattern.FindAllString(description, -1)
	fmt.Println("First match:", firstMatch)
	for i, m := range allMatches {
		fmt.Println("Match", i, "=", m)
	}

	pattern = regexp.MustCompile(" |boat|one")
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}

	pattern = regexp.MustCompile("A [A-z]* for [A-z]* person")
	str := pattern.FindString(description)
	fmt.Println("Match:", str)

	pattern = regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	subs := pattern.FindStringSubmatch(description)
	for _, s := range subs {
		fmt.Println("FindStringSubmatch:", s)
	}

	pattern = regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	subs = pattern.FindStringSubmatch(description)
	for _, name := range []string{"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}

	template := "(type: ${type}, capacity: ${capacity})"
	replaced = pattern.ReplaceAllString(description, template)
	fmt.Println(replaced)

	replaced = pattern.ReplaceAllStringFunc(description, func(s string) string {
		return "This is the replacement content"
	})
	fmt.Println(replaced)
}
