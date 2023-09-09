#include "core_interface.h"

#include <cstring>
#include <map>

#include "core_calculator.hpp"
#include "core_creditor.hpp"

CoreCalculatorInterface CoreCalculatorInit() {
  return static_cast<void*>(new s21::CalculatorModel);
}

void CoreCalculatorFree(CoreCalculatorInterface self) {
  delete static_cast<s21::CalculatorModel*>(self);
}

char* CoreCalculatorPolishNotation(CoreCalculatorInterface self, char* input) {
  s21::CalculatorModel* calculator{static_cast<s21::CalculatorModel*>(self)};

  auto notation{calculator->polishNotation(std::string{input})};
  auto output{new char[notation.length()]};

  std::strcpy(output, notation.c_str());

  return output;
}

double CoreCalculatorCalculate(CoreCalculatorInterface self, char* input,
                               double x) {
  s21::CalculatorModel* calculator{static_cast<s21::CalculatorModel*>(self)};
  std::string notation{input};

  return static_cast<double>(calculator->calculate(notation, x));
}

int CoreCalculatorValidate(CoreCalculatorInterface self, char* input) {
  s21::CalculatorModel* calculator{static_cast<s21::CalculatorModel*>(self)};

  return static_cast<int>(calculator->validateString(std::string{input}));
}

double* CoreCalculatorCalculateAbscissa(CoreCalculatorInterface self,
                                        double begin, double end) {
  s21::CalculatorModel* calculator{static_cast<s21::CalculatorModel*>(self)};
  auto result{calculator->calculateAbscissa(begin, end)};
  double* values{new double[result.size()]{}};

  for (auto i{0}; i < result.size(); ++i) {
    values[i] = static_cast<double>(result[i]);
  }

  return values;
}

double* CoreCalculatorCalculateOrdinate(CoreCalculatorInterface self,
                                        char* input, double begin, double end) {
  s21::CalculatorModel* calculator{static_cast<s21::CalculatorModel*>(self)};
  auto result{calculator->calculateOrdinate(input, begin, end)};
  double* values{new double[result.size()]{}};

  for (auto i{0}; i < result.size(); ++i) {
    values[i] = static_cast<double>(result[i]);
  }

  return values;
}

int CoreCalculatorSteps() {
  return static_cast<int>(s21::CalculatorModel::steps);
}

int CoreCreditorValidate(int credit, int term, double time, double percent,
                         double sum) {
  return static_cast<int>(
      s21::CreditModel::validate(credit, term, time, percent, sum));
}

CreditResult CoreCreditorCalculate(int credit, int term, double time,
                                   double percent, double sum) {
  auto result{s21::CreditModel::calculate(credit, term, time, percent, sum)};

  return CreditResult{
      .monthPay = result["monthPay"],
      .fullPay = result["fullPay"],
      .overPay = result["overPay"],
      .lastPay = result["lastPay"],
  };
}