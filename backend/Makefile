BIN     = calculator
GOSRC   = main.go
CXXPATH = pkg/model/core
CXXSRC  = $(CXXPATH)/core_calculator.cpp \
		  $(CXXPATH)/core_creditor.cpp   \
		  $(CXXPATH)/core_interface.cpp
CXXHDR  = $(CXXPATH)/core_calculator.hpp \
		  $(CXXPATH)/core_creditor.hpp   \
		  $(CXXPATH)/core_interface.h
LIB     = libcalculator.so
SOFLAGS = -fPIC -shared
CXX     = g++ --std=c++17
GO      = go

.PHONY: all clean lib run

lib:
	$(CXX) -o $(CXXPATH)/$(LIB) $(CXXSRC) $(SOFLAGS)

build:
	$(GO) build -o $(BIN) ./main.go

run:
	LD_LIBRARY_PATH=$(shell pwd)/$(CXXPATH) ./$(BIN)

clean:
	rm -rf $(CXXPATH)/$(LIB) $(BIN)
