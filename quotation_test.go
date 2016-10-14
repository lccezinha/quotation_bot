package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetQuotations(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpResponder := httpmock.NewStringResponder(200, `{"status":true,"valores":{"USD":{"nome":"D\u00f3lar","valor":3.1816,"ultima_consulta":1476389101,"fonte":"UOL Economia - http:\/\/economia.uol.com.br\/"},"EUR":{"nome":"Euro","valor":3.5163,"ultima_consulta":1476389101,"fonte":"UOL Economia - http:\/\/economia.uol.com.br\/"},"ARS":{"nome":"Peso Argentino","valor":0.2103,"ultima_consulta":1476389101,"fonte":"UOL Economia - http:\/\/economia.uol.com.br\/"},"GBP":{"nome":"Libra Esterlina","valor":3.8962,"ultima_consulta":1476389101,"fonte":"UOL Economia - http:\/\/economia.uol.com.br\/"},"BTC":{"nome":"Bitcoin","valor":2099.68695,"ultima_consulta":1476403503,"fonte":"Mercado Bitcoin - http:\/\/www.mercadobitcoin.com.br\/"}}}`)
	httpmock.RegisterResponder("GET", "http://api.promasters.net.br/cotacao/v1/valores", httpResponder)

	client := NewQuotationClient()
	quotations := client.GetQuotations()

	quotation := make(map[string]Quotation)
	quotation["USD"] = Quotation{"Dólar", 3.1816}

	assert.Equal(t, quotation["USD"], quotations["USD"])
}
