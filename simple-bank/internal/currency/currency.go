package currency

const (
	USD     = "usd"
	JAKATAS = "jakatas"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, JAKATAS:
		return true
	default:
		return false
	}
}
