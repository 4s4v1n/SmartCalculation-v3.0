#ifndef CALCULATOR_MODEL_INTERFACE_H
#define CALCULATOR_MODEL_INTERFACE_H

#ifdef __cplusplus
extern "C" {
#endif

typedef void* CoreCalculatorInterface;

typedef struct {
  double monthPay;
  double fullPay;
  double overPay;
  double lastPay;
} CreditResult;

CoreCalculatorInterface CoreCalculatorInit();
void CoreCalculatorFree(void* self);
char* CoreCalculatorPolishNotation(void* self, char* input);
double CoreCalculatorCalculate(void* self, char* input, double x);
int CoreCalculatorValidate(void* self, char* input);
double* CoreCalculatorCalculateAbscissa(void* self, double begin, double end);
double* CoreCalculatorCalculateOrdinate(void* self, char* input, double begin,
                                        double end);
int CoreCalculatorSteps();

int CoreCreditorValidate(int credit, int term, double time, double percent,
                         double sum);
CreditResult CoreCreditorCalculate(int credit, int term, double time,
                                   double percent, double sum);

#ifdef __cplusplus
}
#endif

#endif  // CALCULATOR_MODEL_INTERFACE_H
