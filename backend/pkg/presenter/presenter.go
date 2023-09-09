package presenter

import (
	"errors"
	"strings"
	"telvina/APG5_WebCalc/pkg/configurator"
	"telvina/APG5_WebCalc/pkg/keeper"
	"telvina/APG5_WebCalc/pkg/logger"
	calculator2 "telvina/APG5_WebCalc/pkg/model/calculator"
	creditor2 "telvina/APG5_WebCalc/pkg/model/creditor"
)

const (
	piFormatted  = "3.141592"
	expFormatted = "2.718281"
)

type Presenter struct {
	calculatorModel  calculator2.Model
	creditModel      creditor2.Model
	expressionKeeper *keeper.Keeper
	actionsLogger    *logger.Logger
	config           configurator.Config
}

func New(config configurator.Config) *Presenter {
	p := &Presenter{
		calculatorModel:  calculator2.New(),
		creditModel:      creditor2.New(),
		expressionKeeper: keeper.New(),
		config:           config,
	}
	p.actionsLogger = logger.New(p.config.LogsLocation)

	p.LoadExpressions()

	return p
}

func (p *Presenter) InvokeCalculator(input string, x float64) float64 {
	formatted := input

	if strings.Contains(input, "e") {
		formatted = strings.Replace(input, "e", expFormatted, -1)
	}

	if strings.Contains(input, "pi") {
		formatted = strings.Replace(input, "pi", piFormatted, -1)
	}

	notation := p.calculatorModel.PolishNotation(formatted)

	return p.calculatorModel.Calculate(notation, x)
}

func (p *Presenter) InvokeCalculatorValidate(input string) error {
	formatted := input

	if strings.Contains(input, "e") {
		formatted = strings.Replace(input, "e", expFormatted, -1)
	}

	if strings.Contains(input, "pi") {
		formatted = strings.Replace(input, "pi", piFormatted, -1)
	}

	if !p.calculatorModel.Validate(formatted) {
		return errors.New("invalid string")
	}

	return nil
}

func (p *Presenter) InvokeAbscissa(begin float64, end float64) []float64 {
	return p.calculatorModel.CalculateAbscissa(begin, end)
}

func (p *Presenter) InvokeOrdinate(input string, begin float64, end float64) []float64 {
	notation := p.calculatorModel.PolishNotation(input)

	return p.calculatorModel.CalculateOrdinate(notation, begin, end)
}

func (p *Presenter) InvokeFixAbscissaOrdinate(abscissa []float64, ordinate []float64) ([]float64, []float64) {
	return p.calculatorModel.FixAbscissaOrdinate(abscissa, ordinate)
}

func (p *Presenter) InvokeCreditorValidate(creditType string, termType string,
	time float64, percent float64, sum float64) bool {

	var credit int
	if creditType == "Annuity" {
		credit = 0
	} else if creditType == "Different" {
		credit = 1
	} else {
		return false
	}

	var term int
	if termType == "Months" {
		term = 0
	} else if termType == "Years" {
		term = 1
	} else {
		return false
	}

	return p.creditModel.Validate(credit, term, time, percent, sum)
}

func (p *Presenter) InvokeCreditorCalculate(creditType string, termType string, time float64,
	percent float64, sum float64) creditor2.CreditResult {

	var credit int
	if creditType == "Different" {
		credit = 1
	}

	var term int
	if termType == "Years" {
		term = 1
	}

	return p.creditModel.Calculate(credit, term, time, percent, sum)
}

func (p *Presenter) ReleaseModel() {
	p.calculatorModel.Release()
}

func (p *Presenter) LoadExpressions() {
	p.expressionKeeper.Load(p.config.HistoryLocation)
}

func (p *Presenter) SaveExpressions() {
	p.expressionKeeper.Save(p.config.HistoryLocation)
}

func (p *Presenter) StoreExpression(expression string) {
	p.expressionKeeper.Add(expression)
}

func (p *Presenter) GetExpression() string {
	return p.expressionKeeper.Get()
}

func (p *Presenter) ClearExpressions() {
	p.expressionKeeper.Clear()
}

func (p *Presenter) ReleaseLogger() {
	p.actionsLogger.Release()
}
