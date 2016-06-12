package uuid

import (
	"errors"
	"strings"
)

// These constants represent a pattern used by the package with which to print
// a UUID.
type Format string

const (
	Clean   = "%x%x%x%x%x"
	Curly   = "{%x%x%x%x%x}"
	Bracket = "(%x%x%x%x%x)"

	// This is the canonical format.
	CleanHyphen = "%x-%x-%x-%x-%x"

	CurlyHyphen   = "{%x-%x-%x-%x-%x}"
	BracketHyphen = "(%x-%x-%x-%x-%x)"
)

var printFormat Format = CurlyHyphen

// SwitchFormat switches the default printing format for ALL UUIDs.
//
// The default is the canonical uuid.Format.CleanHyphen which has been
// optimised for use with this package and is twice as fast as using any other
// format; supplied or given. The benchmark for non default formats is still
// very quick and quite usable. The package has moved away from using
// fmt.Sprintf which was up to 5 times slower in comparison to custom formats
// and 10 times slower in comparison to the canonical format.
//
// A valid format will have 5 groups of [%x|%X] or follow the pattern,
// ^*%[xX]*%[xX]*%[xX]*%[xX]*%[xX]*$. If the supplied format does not meet this
// standard the function will panic. Constants have been provided for the most
// likely formats.
func SwitchFormat(pFormat Format) {
	checkFormat(string(pFormat))
	printFormat = pFormat
}

// SwitchFormatToUpper is a convenience function to set the Format to upper case
// versions of the given constants.
func SwitchFormatToUpper(pFormat Format) {
	checkFormat(string(pFormat))
	printFormat = Format(strings.ToUpper(string(pFormat)))
}

// Sprintf will return a string representation of the given UUID.
//
// Use this for one time formatting when setting the default using
// uuid.SwitchFormat
//
// A valid format will have 5 groups of [%x|%X] or follow the pattern,
// ^*%[xX]*%[xX]*%[xX]*%[xX]*%[xX]*$. If the supplied format does not meet this
// standard the function will panic.
func Sprintf(pFormat Format, pId UUID) string {
	checkFormat(string(pFormat))
	return formatPrint(pId.Bytes(), string(pFormat))
}

func checkFormat(pFormat string) {
	s := strings.ToLower(pFormat)
	if strings.Count(s, "%x") != 5 {
		panic(errors.New("uuid.Print: invalid format"))
	}
}

const (
	hexTable      = "0123456789abcdef"
	hexUpperTable = "0123456789ABCDEF"

	canonicalLength      = length*2 + 4
	formatArgCount       = 10
	uuidStringBufferSize = length*2 - formatArgCount
)

var groups = []int{4, 2, 2, 2, 6}

func formatPrint(pSrc []byte, pFormat string) string {
	end := len(pFormat)
	buf := make([]byte, end+uuidStringBufferSize)

	var s, ls, b, e, p int
	var u bool
	for _, v := range groups {
		ls = s
		for s < end && (pFormat[s] != '%') {
			s++
			copy(buf[p:], pFormat[ls:s])
			p += s - ls
		}
		s++
		u = pFormat[s] == 'X'
		s++
		e = b + v
		for i, t := range pSrc[b:e] {
			j := p + i + i
			table := hexTable
			if u {
				table = hexUpperTable
			}
			buf[j] = table[t>>4]
			buf[j+1] = table[t&0x0f]
		}
		b = e
		p += v + v
	}
	ls = s
	for s < end && (pFormat[s] != '%') {
		s++
		copy(buf[p:], pFormat[ls:s])
		p += s - ls
	}
	return string(buf)
}

func canonicalPrint(pSrc []byte) string {
	buf := make([]byte, canonicalLength)
	var b, p, e int
	for h, v := range groups {
		e = b + v
		for i, t := range pSrc[b:e] {
			j := p + i + i
			buf[j] = hexTable[t>>4]
			buf[j+1] = hexTable[t&0x0f]
		}
		b = e
		p += v + v
		if h < 4 {
			buf[p] = '-'
			p += 1
		}
	}
	return string(buf)
}
