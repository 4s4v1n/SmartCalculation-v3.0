package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Valid       bool   `json:"valid"`
	Description string `json:"description"`
}

type calculatorResponse struct {
	Value string `json:"value"`
}

type plotResponse struct {
	Abscissa []float64 `json:"abscissa"`
	Ordinate []float64 `json:"ordinate"`
}

type creditResponse struct {
	MonthPay float64 `json:"month_pay"`
	FullPay  float64 `json:"full_pay"`
	OverPay  float64 `json:"over_pay"`
	LastPay  float64 `json:"last_pay"`
}

type previousExpressionResponse struct {
	Expression string `json:"expression"`
}

func writeError(w http.ResponseWriter, error error, status int) {
	w.WriteHeader(status)

	e := &errorResponse{
		Valid:       false,
		Description: error.Error(),
	}

	byteMessage, _ := json.Marshal(e)

	if _, err := w.Write(byteMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeResponseJson(w http.ResponseWriter, body interface{}) {
	resp, err := json.Marshal(body)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
}

func missingParameterError(parameter string) error {
	return errors.New(fmt.Sprintf("query missing %s parameter", parameter))
}
