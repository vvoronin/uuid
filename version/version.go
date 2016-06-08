package version

// Version represents the type of UUID.
type Version uint8

const (
	Unknown Version = iota // Unknown
	One                    // Time based
	Two                  // DCE security via POSIX UIDs
	Three                 // Namespace hash uses MD5
	Four                     // Crypto random
	Five                // Namespace hash uses SHA-1
)

// String returns English description of version.
func (o Version) String() string {
	switch o {
	case One:
		return "Version 1: Based on a 60 Bit Timestamp"
	case Two:
		return "Version 2: Based on DCE security domain and 60 bit timestamp"
	case Three:
		return "Version 3: Namespace UUID and unique names hashed by MD5"
	case Four:
		return "Version 4: Crypto-random"
	case Five:
		return "Version 5: Namespace UUID and unique names hashed by SHA-1"
	default:
		return "Unknown: Not supported"
	}
}