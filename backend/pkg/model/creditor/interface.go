package creditor

type CreditResult struct {
	MonthPay float64
	FullPay  float64
	OverPay  float64
	LastPay  float64
}

type Model interface {
	Validate(creditType int, term int, time float64, percent float64, sum float64) bool
	Calculate(creditType int, term int, time float64, percent float64, sum float64) CreditResult
}
