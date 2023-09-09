// Copyright telvina 2022
#include "core_creditor.hpp"

#include <cmath>

namespace s21 {

auto CreditModel::validate(int credit, int term, double time, double percent,
                           double sum) -> bool {
  if (credit != CreditType::Annuity && credit != CreditType::Different) {
    return false;
  }
  if (term != TermType::Month && term != TermType::Year) {
    return false;
  }
  if (time <= 0) {
    return false;
  }
  if (percent <= 0) {
    return false;
  }
  if (sum <= 0) {
    return false;
  }

  return true;
}

auto CreditModel::calculate(int credit, int term, double time, double percent,
                            double sum) -> std::map<std::string, double> {
  std::map<std::string, double> result{};

  percent /= 100 * 12;

  if (term == TermType::Year) {
    time *= 12;
  }

  if (credit == CreditType::Annuity) {
    result["monthPay"] =
        sum * (percent + (percent / (std::pow(1 + percent, time) - 1)));
    result["fullPay"] = result["monthPay"] * time;
    result["overPay"] = result["fullPay"] - sum;
  } else {
    auto coefficient{sum / time};
    auto start{sum};

    for (auto i{0}; i < time; ++i) {
      result["lastPay"] = coefficient + sum * percent;
      result["fullPay"] += result["lastPay"];
      sum -= coefficient;

      if (i == 0) {
        result["monthPay"] = result["lastPay"];
      }
    }
    result["overPay"] = result["fullPay"] - start;
  }

  return result;
}

}  // namespace s21
