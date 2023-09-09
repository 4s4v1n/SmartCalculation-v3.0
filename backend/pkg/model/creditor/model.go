package creditor

//#cgo LDFLAGS: -L../core -lcalculator
//#include <stdlib.h>
//#include "../core/core_interface.h"
import "C"

type CreditModel struct {
}

func New() *CreditModel {
	return &CreditModel{}
}

func (c CreditModel) Validate(creditType int, term int, time float64, percent float64, sum float64) bool {
	if int(C.int(C.CoreCreditorValidate(C.int(creditType), C.int(term), C.double(time),
		C.double(percent), C.double(sum)))) == 1 {
		return true
	} else {
		return false
	}
}

func (c CreditModel) Calculate(creditType int, term int, time float64, percent float64, sum float64) CreditResult {
	result := C.CoreCreditorCalculate(C.int(creditType), C.int(term),
		C.double(time), C.double(percent), C.double(sum))

	return CreditResult{
		MonthPay: float64(result.monthPay),
		FullPay:  float64(result.fullPay),
		OverPay:  float64(result.overPay),
		LastPay:  float64(result.lastPay),
	}
}
