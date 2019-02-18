package pagarme

// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BR
// Pagarme says it is ISO_3166-1_alpha-2 but it's actually
// lowercase instead of uppercase

type CountryCode string

func (cc CountryCode) IsValid() bool {
	if len(cc) != 2 {
		return false
	}
	for _, r := range cc {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

const (
	Brazil CountryCode = "br"
)
