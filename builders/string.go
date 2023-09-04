package builders

import "strings"

type stringBuilder struct {
	value string
}

func (b *builders) String() stringBuilder {
	return stringBuilder{
		value: "",
	}
}

func (sb *stringBuilder) Append(value string) stringBuilder {
	sb.value = sb.value + value

	return *sb
}

// Appends multiple items to the array and returns the builder.
func (sb *stringBuilder) AppendMultiple(value []string) stringBuilder {
	for _, val := range value {
		sb.value = sb.value + val
	}

	return *sb
}

func (sb *stringBuilder) EveryWord(predicate func(word string, str string) bool) bool {
	var (
		words = strings.Split(sb.value, " ")
		s []string
	)

	for i := 0; i < len(words);
}

func (sb *stringBuilder) ForEachWord(predicate func(word string, str string)) stringBuilder {
	var (
		words = strings.Split(sb.value, " ")
		s []string
	)

	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], ",") {
			s = append(s, strings.Replace(words[i], ",", "", 1))
			s = append(s, ",") 
		} else {
			s = append(s, words[i])
		}
	}

	for i := 0; i < len(s); i++ {
		predicate(s[i], sb.value)
	}
}