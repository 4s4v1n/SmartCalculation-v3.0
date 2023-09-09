// Copyright telvina 2022

#include "core_calculator.hpp"

#include <algorithm>
#include <utility>

namespace s21 {

auto CalculatorModel::polishNotation(const std::string &str) -> std::string {
  std::string notation{};
  std::string input{str};
  std::cmatch numbers{};
  auto iter{input.begin()};

  while (iter != input.end()) {
    if (std::regex_search(iter.base(), numbers, numberRegex)) {
      auto position{iter - input.begin()};
      auto length{numbers.str().length()};

      notation += input.substr(position, length);
      notation += " ";
      position += static_cast<long>(length);
      iter += static_cast<long>(length);
    } else {
      auto operation{whichOperation(iter)};

      if (operation.first != "(" && operation.first != ")") {
        if (operationStack.empty()) {
          operationStack.push(operation);
        } else if (operation.second > operationStack.top().second) {
          operationStack.push(operation);
        } else {
          while (!operationStack.empty() &&
                 operationStack.top().second >= operation.second) {
            notation += operationStack.top().first + " ";
            operationStack.pop();
          }
          operationStack.push(operation);
        }
      } else if (operation.first == "(") {
        operationStack.push(operation);
      } else {
        while (!operationStack.empty() && operationStack.top().first != "(") {
          notation += operationStack.top().first + " ";
          operationStack.pop();
        }
        if (!operationStack.empty() && operationStack.top().first == "(") {
          operationStack.pop();
        }
      }
    }
  }

  while (!operationStack.empty()) {
    notation += operationStack.top().first + " ";
    operationStack.pop();
  }

  return notation;
}

auto CalculatorModel::whichOperation(std::string::iterator &iter)
    -> CalculatorModel::Operation {
  Operation operation{};
  std::string str{iter.base()};

  if (str.substr(0, 3) == "cos") {
    operation.first = "cos";
    iter += 3;
  } else if (str.substr(0, 3) == "sin") {
    operation.first = "sin";
    iter += 3;
  } else if (str.substr(0, 3) == "tan") {
    operation.first = "tan";
    iter += 3;
  } else if (str.substr(0, 4) == "asin") {
    operation.first = "asin";
    iter += 4;
  } else if (str.substr(0, 4) == "acos") {
    operation.first = "acos";
    iter += 4;
  } else if (str.substr(0, 4) == "atan") {
    operation.first = "atan";
    iter += 4;
  } else if (str.substr(0, 4) == "sqrt") {
    operation.first = "sqrt";
    iter += 4;
  } else if (str.substr(0, 2) == "ln") {
    operation.first = "ln";
    iter += 2;
  } else if (str.substr(0, 2) == "lg") {
    operation.first = "lg";
    iter += 2;
  } else if (*str.data() == '^') {
    operation.first = "^";
    ++iter;
  } else if (*str.data() == '%') {
    operation.first = "%";
    ++iter;
  } else if (*str.data() == '*') {
    operation.first = "*";
    ++iter;
  } else if (*str.data() == '/') {
    operation.first = "/";
    ++iter;
  } else if (*str.data() == '+') {
    --iter;
    if (*iter != '\0') {
      if ((*iter >= 48 && *iter <= 57) || *iter == 'x' || *iter == ')') {
        operation.first = "+";
      } else {
        operation.first = "U+";
      }
    } else {
      operation.first = "U+";
    }
    iter += 2;
  } else if (*str.data() == '-') {
    --iter;
    if (*iter != '\0') {
      if ((*iter >= 48 && *iter <= 57) || *iter == 'x' || *iter == ')') {
        operation.first = "-";
      } else {
        operation.first = "U-";
      }
    } else {
      operation.first = "U-";
    }
    iter += 2;
  } else if (*str.data() == '(') {
    operation.first = "(";
    ++iter;
  } else if (*str.data() == ')') {
    operation.first = ")";
    ++iter;
  }

  priority(operation);

  return operation;
}

auto CalculatorModel::priority(Operation &operation) -> void {
  if (operation.first == "sin" || operation.first == "cos" ||
      operation.first == "tan" || operation.first == "asin" ||
      operation.first == "acos" || operation.first == "atan" ||
      operation.first == "lg" || operation.first == "ln" ||
      operation.first == "sqrt") {
    operation.second = 5;
  } else if (operation.first == "U+" || operation.first == "U-") {
    operation.second = 4;
  } else if (operation.first == "^") {
    operation.second = 3;
  } else if (operation.first == "%" || operation.first == "*" ||
             operation.first == "/") {
    operation.second = 2;
  } else if (operation.first == "+" || operation.first == "-") {
    operation.second = 1;
  } else if (operation.first == "(" || operation.first == ")") {
    operation.second = 0;
  }
}

auto CalculatorModel::calculate(const std::string &str, long double x)
    -> long double {
  auto notation{str};
  char *plug{nullptr};
  auto iter{notation.begin()};
  std::cmatch matcher{};

  while (!numberStack.empty()) {
    numberStack.pop();
  }

  while (iter != notation.end()) {
    if (std::regex_search(iter.base(), matcher, digitRegex)) {
      numberStack.push(std::strtold(iter.base(), &plug));
      iter += static_cast<long>(matcher.str().size());
    } else if (*iter == 'x') {
      numberStack.push(x);
      ++iter;
    } else {
      calculateUsingStack(iter);
    }
    if (*iter == ' ') ++iter;
  }

  return numberStack.top();
}

void CalculatorModel::calculateUsingStack(std::string::iterator &iter) {
  long double first{0};

  std::string str = iter.base();
  if (*str.data() == '+' || *str.data() == '-' || *str.data() == '*' ||
      *str.data() == '/' || *str.data() == '%' || *str.data() == '^') {
    first = numberStack.top();
    numberStack.pop();

    long double second{numberStack.top()};

    numberStack.pop();
    if (*str.data() == '+') {
      numberStack.push(second + first);
      iter += 1;
    } else if (*str.data() == '-') {
      numberStack.push(second - first);
      iter += 1;
    } else if (*str.data() == '*') {
      numberStack.push(second * first);
      iter += 1;
    } else if (*str.data() == '/') {
      numberStack.push(second / first);
      iter += 1;
    } else if (*str.data() == '%') {
      numberStack.push(fmodl(second, first));
      iter += 1;
    } else if (*str.data() == '^') {
      numberStack.push(powl(second, first));
      iter += 1;
    }
  } else {
    first = numberStack.top();
    numberStack.pop();
    if (str.substr(0, 3) == "cos") {
      numberStack.push(cosl(first));
      iter += 3;
    } else if (str.substr(0, 3) == "sin") {
      numberStack.push(sinl(first));
      iter += 3;
    } else if (str.substr(0, 4) == "sqrt") {
      numberStack.push(sqrtl(first));
      iter += 4;
    } else if (str.substr(0, 3) == "tan") {
      numberStack.push(tanl(first));
      iter += 3;
    } else if (str.substr(0, 4) == "asin") {
      numberStack.push(asinl(first));
      iter += 4;
    } else if (str.substr(0, 4) == "acos") {
      numberStack.push(acosl(first));
      iter += 4;
    } else if (str.substr(0, 4) == "atan") {
      numberStack.push(atanl(first));
      iter += 4;
    } else if (str.substr(0, 2) == "ln") {
      numberStack.push(logl(first));
      iter += 2;
    } else if (str.substr(0, 2) == "lg") {
      numberStack.push(log10l(first));
      iter += 2;
    } else if (str.substr(0, 2) == "U-") {
      numberStack.push(-first);
      iter += 2;
    } else if (str.substr(0, 2) == "U+") {
      numberStack.push(first);
      iter += 2;
    }
  }
}

bool CalculatorModel::bracketBalanced(const std::string &str) {
  auto iter{str.begin()};
  int counter{0};

  while (iter != str.end()) {
    if (*iter == '(') counter++;
    if (*iter == ')') counter--;
    ++iter;
  }

  return counter == 0;
}

auto CalculatorModel::actionLexem(std::string::iterator &iter)
    -> CalculatorModel::Lexem {
  std::cmatch matcher{};
  Lexem lexem{WRONG};

  if (std::regex_search(iter.base(), matcher, numberRegex)) {
    lexem = NUMBER;
    iter += static_cast<long>(matcher.str().size());
  } else if (std::regex_search(iter.base(), matcher, binaryUnaryRegex)) {
    lexem = BINARY_UNARY;
    iter += static_cast<long>(matcher.str().size());
  } else if (std::regex_search(iter.base(), matcher, binaryRegex)) {
    lexem = BINARY;
    iter += static_cast<long>(matcher.str().size());
  } else if (std::regex_search(iter.base(), matcher, unaryRegex)) {
    lexem = UNARY;
    iter += static_cast<long>(matcher.str().size());
  } else if (*iter == '(') {
    lexem = OPEN_BRACKET;
    ++iter;
  } else if (*iter == ')') {
    lexem = CLOSE_BRACKET;
    ++iter;
  }
  return lexem;
}

bool CalculatorModel::validateString(const std::string &str) {
  if (str.empty() || !bracketBalanced(str)) {
    return false;
  }

  auto input{str};
  auto iter{input.begin()};
  auto current{actionLexem(iter)};

  if (current == CLOSE_BRACKET || current == BINARY) {
    return false;
  }

  while (iter != input.end()) {
    auto previous{current};

    current = actionLexem(iter);

    if (current == OPEN_BRACKET) {
      if (previous == NUMBER || previous == CLOSE_BRACKET) {
        return false;
      }
      continue;
    } else if (current == NUMBER) {
      if (previous == NUMBER || previous == CLOSE_BRACKET) {
        return false;
      }
      continue;
    } else if (current == BINARY_UNARY) {
      if (previous == BINARY_UNARY) {
        return false;
      }
      continue;
    } else if (current == BINARY) {
      if (previous == OPEN_BRACKET || previous == BINARY_UNARY ||
          previous == UNARY || previous == BINARY) {
        return false;
      }
      continue;
    } else if (current == UNARY) {
      if (previous == NUMBER || previous == CLOSE_BRACKET) {
        return false;
      }
      continue;
    } else if (current == CLOSE_BRACKET) {
      if (previous == OPEN_BRACKET || previous == BINARY_UNARY ||
          previous == UNARY || previous == BINARY) {
        return false;
      }
      continue;
    } else {
      return false;
    }
  }

  if (current == OPEN_BRACKET || current == BINARY || current == UNARY ||
      current == BINARY_UNARY) {
    return false;
  }
  return true;
}

auto CalculatorModel::calculateAbscissa(double begin, double end)
    -> std::vector<long double> {
  std::vector<long double> result{};
  auto step{std::abs(end - begin) / steps};

  for (auto i{0}; i < steps; ++i) {
    result.push_back(begin + i * step);
  }
  result.push_back(end);

  return result;
}

auto CalculatorModel::calculateOrdinate(const std::string &str, double begin,
                                        double end)
    -> std::vector<long double> {
  std::vector<long double> abscissaValues{calculateAbscissa(begin, end)};
  std::vector<long double> result{};

  result.reserve(abscissaValues.size());
  std::for_each(
      abscissaValues.begin(), abscissaValues.end(),
      [&](const double value) { result.push_back(calculate(str, value)); });

  return result;
}

}  // namespace s21
