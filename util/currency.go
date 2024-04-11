package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	INR = "INR"
)

// IsSupportedCurrency returns if the currency is supported

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, INR:
		return true
	}
	return false
}
