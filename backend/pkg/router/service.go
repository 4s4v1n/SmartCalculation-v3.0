package router

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

const (
	expressionParameter = `expression`
	xParameter          = `x`
	beginParameter      = `begin`
	endParameter        = `end`
	typeParameter       = `type`
	termParameter       = `term`
	timeParameter       = `time`
	percentParameter    = `percent`
	sumParameter        = `sum`
)

func (rtr *Router) calculatorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	expression := r.URL.Query().Get(expressionParameter)
	if expression == "" {
		writeError(w, missingParameterError(expressionParameter), http.StatusBadRequest)
		return
	}

	if strings.Contains(expression, "x") && !r.URL.Query().Has(xParameter) {
		writeError(w, missingParameterError(xParameter), http.StatusBadRequest)
		return
	}

	if err := rtr.prs.InvokeCalculatorValidate(expression); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	var err error
	var x float64
	if strings.Contains(expression, "x") {
		if x, err = strconv.ParseFloat(r.URL.Query().Get("x"), 64); err != nil {
			writeError(w, err, http.StatusBadRequest)
		}
	}

	rtr.prs.StoreExpression(expression)
	logrus.Infof("endpoint: %s, value: %s", "/calculator", expression)
	writeResponseJson(w, calculatorResponse{
		Value: strconv.FormatFloat(rtr.prs.InvokeCalculator(expression, x), 'f', 6, 64),
	})
}

func (rtr *Router) creditHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	creditType := r.URL.Query().Get(typeParameter)
	if creditType == "" {
		writeError(w, missingParameterError(typeParameter), http.StatusBadRequest)
		return
	}

	term := r.URL.Query().Get(termParameter)
	if term == "" {
		writeError(w, missingParameterError(termParameter), http.StatusBadRequest)
		return
	}

	if !r.URL.Query().Has(timeParameter) {
		writeError(w, missingParameterError(timeParameter), http.StatusBadRequest)
		return
	}
	time, err := strconv.ParseFloat(r.URL.Query().Get(timeParameter), 64)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if !r.URL.Query().Has(percentParameter) {
		writeError(w, missingParameterError(percentParameter), http.StatusBadRequest)
		return
	}
	percent, err := strconv.ParseFloat(r.URL.Query().Get(percentParameter), 64)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if !r.URL.Query().Has(sumParameter) {
		writeError(w, missingParameterError(sumParameter), http.StatusBadRequest)
		return
	}
	sum, err := strconv.ParseFloat(r.URL.Query().Get(sumParameter), 64)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if !rtr.prs.InvokeCreditorValidate(creditType, term, time, percent, sum) {
		writeError(w, errors.New("data not valid"), http.StatusUnprocessableEntity)
		return
	}

	res := rtr.prs.InvokeCreditorCalculate(creditType, term, time, percent, sum)

	logrus.Infof("endpoint: %s", "/credit")
	writeResponseJson(w, creditResponse{
		MonthPay: res.MonthPay,
		FullPay:  res.FullPay,
		OverPay:  res.OverPay,
		LastPay:  res.LastPay,
	})
}

func (rtr *Router) plotHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	expression := r.URL.Query().Get(expressionParameter)
	if expression == "" {
		writeError(w, missingParameterError(expressionParameter), http.StatusBadRequest)
		return
	}

	if !strings.Contains(expression, "x") {
		writeError(w, errors.New("query parameter expression missing x"), http.StatusBadRequest)
		return
	}

	if !r.URL.Query().Has(beginParameter) {
		writeError(w, errors.New("query parameter begin is missing"), http.StatusBadRequest)
		return
	}
	begin, err := strconv.ParseFloat(r.URL.Query().Get(beginParameter), 64)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if !r.URL.Query().Has(endParameter) {
		writeError(w, missingParameterError(endParameter), http.StatusBadRequest)
		return
	}
	end, err := strconv.ParseFloat(r.URL.Query().Get(endParameter), 64)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if err = rtr.prs.InvokeCalculatorValidate(expression); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	abscissa, ordinate := rtr.prs.InvokeFixAbscissaOrdinate(
		rtr.prs.InvokeAbscissa(begin, end),
		rtr.prs.InvokeOrdinate(expression, begin, end))

	logrus.Infof("endpoint: %s, value: %s, %f, %f", "/plot", expression, begin, end)
	writeResponseJson(w, plotResponse{
		Abscissa: abscissa,
		Ordinate: ordinate,
	})
}

func (rtr *Router) previousExpressionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	expression := rtr.prs.GetExpression()

	logrus.Infof("endpoint: %s, value: %s", "/previous_expression", expression)
	writeResponseJson(w, previousExpressionResponse{
		Expression: expression,
	})
}

func (rtr *Router) clearHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	logrus.Infof("endpoint: %s", "/clear_history")
	rtr.prs.ClearExpressions()
}
