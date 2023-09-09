package calculator

//#cgo LDFLAGS: -lcalculator
//#include <stdlib.h>
//#include "../core/core_interface.h"
import "C"
import (
	"math"
	"unsafe"
)

type Calculator struct {
	calculator C.CoreCalculatorInterface
}

func New() *Calculator {
	return &Calculator{
		calculator: C.CoreCalculatorInit(),
	}
}

func (c Calculator) PolishNotation(input string) string {
	notation := C.CoreCalculatorPolishNotation(unsafe.Pointer(c.calculator), C.CString(input))
	defer C.free(unsafe.Pointer(notation))

	return C.GoString(notation)
}

func (c Calculator) Calculate(input string, x float64) float64 {
	return float64(C.CoreCalculatorCalculate(unsafe.Pointer(c.calculator), C.CString(input), C.double(x)))
}

func (c Calculator) Validate(input string) bool {
	if int(C.int(C.CoreCalculatorValidate(unsafe.Pointer(c.calculator), C.CString(input)))) == 1 {
		return true
	} else {
		return false
	}
}

func (c Calculator) Release() {
	C.CoreCalculatorFree(unsafe.Pointer(c.calculator))
}

func (c Calculator) CalculateOrdinate(input string, begin float64, end float64) []float64 {
	var values *C.double = C.CoreCalculatorCalculateOrdinate(unsafe.Pointer(c.calculator),
		C.CString(input), C.double(begin), C.double(end))
	defer C.free(unsafe.Pointer(values))

	slice := unsafe.Slice(values, int(C.CoreCalculatorSteps())+1)

	var result []float64
	for _, value := range slice {
		result = append(result, float64(value))
	}

	return result
}

func (c Calculator) CalculateAbscissa(begin float64, end float64) []float64 {
	var values *C.double = C.CoreCalculatorCalculateAbscissa(unsafe.Pointer(c.calculator),
		C.double(begin), C.double(end))
	defer C.free(unsafe.Pointer(values))

	slice := unsafe.Slice(values, int(C.CoreCalculatorSteps())+1)

	var result []float64
	for _, value := range slice {
		result = append(result, float64(value))
	}

	return result
}

func (c Calculator) FixAbscissaOrdinate(abscissa []float64, ordinate []float64) ([]float64, []float64) {
	var fixedAbscissa []float64
	var fixedOrdinate []float64

	for i, item := range ordinate {
		if !math.IsNaN(item) && !math.IsInf(item, 1) && !math.IsInf(item, -1) {
			fixedAbscissa = append(fixedAbscissa, abscissa[i])
			fixedOrdinate = append(fixedOrdinate, ordinate[i])
		}
	}

	return fixedAbscissa, fixedOrdinate
}
