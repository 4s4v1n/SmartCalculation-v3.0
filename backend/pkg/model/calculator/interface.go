package calculator

type Model interface {
	PolishNotation(input string) string
	Calculate(input string, x float64) float64
	CalculateAbscissa(begin float64, end float64) []float64
	CalculateOrdinate(input string, begin float64, end float64) []float64
	FixAbscissaOrdinate(abscissa []float64, ordinate []float64) ([]float64, []float64)
	Validate(input string) bool
	Release()
}
