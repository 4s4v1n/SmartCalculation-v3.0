#ifndef CREDIT_MODEL_H
#define CREDIT_MODEL_H

#include <map>
#include <string>

namespace s21 {
class CreditModel {
  enum TermType {
    Month = 0,
    Year = 1,
  };

  enum CreditType {
    Annuity = 0,
    Different = 1,
  };

 public:
  CreditModel() = default;
  ~CreditModel() = default;

 public:
  auto static validate(int type, int term, double time, double percent,
                       double sum) -> bool;
  auto static calculate(int credit, int term, double time, double percent,
                        double sum) -> std::map<std::string, double>;
};
}  // namespace s21

#endif  // CREDIT_MODEL_H
