package uuid

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

const (
	clean                   = `[[:xdigit:]]{8}[[:xdigit:]]{4}[1-5][[:xdigit:]]{3}[[:xdigit:]]{4}[[:xdigit:]]{12}`
	cleanHexPattern         = `^` + clean + `$`
	curlyHexPattern         = `^\{` + clean + `\}$`
	bracketHexPattern       = `^\(` + clean + `\)$`
	hyphen                  = `[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[[:xdigit:]]{4}-[[:xdigit:]]{12}`
	cleanHyphenHexPattern   = `^` + hyphen + `$`
	curlyHyphenHexPattern   = `^\{` + hyphen + `\}$`
	bracketHyphenHexPattern = `^\(` + hyphen + `\)$`
)

func TestSwitchFormat(t *testing.T) {
	ids := []UUID{NewV4(), NewV4()}
	formats := []Format{CurlyHyphen, Clean, Curly, Bracket, CleanHyphen, BracketHyphen}
	patterns := []string{curlyHyphenHexPattern, cleanHexPattern, curlyHexPattern, bracketHexPattern, cleanHyphenHexPattern, bracketHyphenHexPattern}

	// Reset default
	SwitchFormat(CleanHyphen)

	for _, u := range ids {
		for i := range formats {
			SwitchFormat(formats[i])
			assert.True(t, regexp.MustCompile(patterns[i]).MatchString(u.String()), "Format %s must compile pattern %s", formats[i], patterns[i])
			outputLn(u)
		}
	}

	assert.True(t, didSwitchFormatPanic(), "Switch format should panic when format invalid")

	// Reset default
	SwitchFormat(CleanHyphen)
}

func TestSwitchFormatToUpper(t *testing.T) {
	ids := []UUID{NewV4(), NewV4()}
	formats := []Format{CurlyHyphen, Clean, Curly, Bracket, CleanHyphen, BracketHyphen}
	patterns := []string{curlyHyphenHexPattern, cleanHexPattern, curlyHexPattern, bracketHexPattern, cleanHyphenHexPattern, bracketHyphenHexPattern}

	// Reset default
	SwitchFormat(CleanHyphen)

	for _, u := range ids {
		for i := range formats {
			SwitchFormat(formats[i])
			assert.True(t, regexp.MustCompile(patterns[i]).MatchString(u.String()), "Format %s must compile pattern %s", formats[i], patterns[i])
			outputLn(u)
		}
	}

	assert.True(t, didSwitchFormatPanic(), "Switch format should panic when format invalid")

	// Reset default
	SwitchFormat(CleanHyphen)
}

func didSwitchFormatPanic() bool {
	return func() (didPanic bool) {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()

		SwitchFormat("%%%%%%%%%%%%%")
		return
	}()
}

func TestSprintf(t *testing.T) {
	ids := []UUID{NewV4(), NewV4()}
	formats := []Format{CurlyHyphen, Clean, Curly, Bracket, CleanHyphen, BracketHyphen}
	patterns := []string{curlyHyphenHexPattern, cleanHexPattern, curlyHexPattern, bracketHexPattern, cleanHyphenHexPattern, bracketHyphenHexPattern}

	for _, u := range ids {
		for i := range formats {
			assert.True(t, regexp.MustCompile(patterns[i]).MatchString(Sprintf(formats[i], u)), "Format must compile")
			outputLn(Sprintf(formats[i], u))
		}
	}

	assert.True(t, didSprintfPanic(), "Sprinf should panic when format invalid")
}

func didSprintfPanic() bool {
	return func() (didPanic bool) {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()

		Sprintf("%s*********-------)()()()()(", NameSpaceDNS)
		return
	}()
}
