package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatSingleValue(t *testing.T) {
	quotation := make(map[string]Quotation)
	quotation["USD"] = Quotation{"Dólar", 3.20}

	expected := "Cotação Dólar em relação ao Real: R$ 3.20 = 1 Dólar\n"
	result := Formatter{quotation}.Format()

	assert.Equal(t, expected, result)
}

func TestFormatMultipleValues(t *testing.T) {
	quotation := make(map[string]Quotation)
	quotation["USD"] = Quotation{"Dólar", 3.20}
	quotation["ABC"] = Quotation{"ABC", 10.0}

	expected := "Cotação Dólar em relação ao Real: R$ 3.20 = 1 Dólar\n"
	expected += "Cotação ABC em relação ao Real: R$ 10.00 = 1 ABC\n"

	result := Formatter{quotation}.Format()

	assert.Contains(t, expected, result)
}
