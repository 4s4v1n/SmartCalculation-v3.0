#ifndef CORE_CALCULATOR_H
#define CORE_CALCULATOR_H

#include <cmath>
#include <iostream>
#include <regex>
#include <stack>
#include <string>
#include <vector>

namespace s21 {
class CalculatorModel {
  using Operation = std::pair<std::string, size_t>;

  enum Lexem {
    WRONG = 0,
    NUMBER = 1,
    OPEN_BRACKET = 2,
    CLOSE_BRACKET = 3,
    BINARY = 4,
    BINARY_UNARY = 5,
    UNARY = 6
  };

 public:
  CalculatorModel() = default;
  ~CalculatorModel() = default;

 public:
  auto polishNotation(const std::string &str) -> std::string;
  auto calculate(const std::string &str, long double x = 0) -> long double;
  auto validateString(const std::string &str) -> bool;
  auto calculateAbscissa(double begin, double end) -> std::vector<long double>;
  auto calculateOrdinate(const std::string &str, double begin, double end)
      -> std::vector<long double>;

 private:
  auto whichOperation(std::string::iterator &iter) -> Operation;
  auto priority(Operation &operation) -> void;
  auto calculateUsingStack(std::string::iterator &iter) -> void;
  auto bracketBalanced(const std::string &str) -> bool;
  auto actionLexem(std::string::iterator &iter) -> Lexem;

 public:
  constexpr static double steps{100};

 private:
  const std::regex numberRegex{R"((^\d+\.\d+|^\d+|^x))"};
  const std::regex unaryRegex{"^(ln|lg|sin|cos|tan|asin|acos|atan|sqrt)\\("};
  const std::regex binaryUnaryRegex{"^(\\+|-)"};
  const std::regex binaryRegex{"^(\\*|/|\\^|%)"};
  const std::regex digitRegex{R"(^(\d+\.\d+|\d+|-\d+|-\d+\.\d+))"};

  std::string _inputLine;
  std::string _outputLine;

  std::stack<Operation> operationStack;
  std::stack<long double> numberStack;
};

}  // namespace s21

#endif  // CORE_CALCULATOR_H
