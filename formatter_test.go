package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatSingleValue(t *testing.T) {
	quotation := make(map[string]Quotation)
	quotation["USD"] = Quotation{"Dólar", 3.20}

	expected := "\t\tCotações\t\t\n\n"
	expected += "1 Dólar = R$ 3.20\n"

	result := Formatter{quotation}.Format()

	assert.Equal(t, expected, result)
}

func TestFormatMultipleValues(t *testing.T) {
	quotation := make(map[string]Quotation)
	quotation["USD"] = Quotation{"Dólar", 3.20}
	quotation["ABC"] = Quotation{"ABC", 10.0}

	expected := "\t\tCotações\t\t\n\n"
	expected += "1 Dólar = R$ 3.20\n"
	expected += "1 ABC = R$ 10.00\n"

	result := Formatter{quotation}.Format()

	assert.Contains(t, expected, result)
}
