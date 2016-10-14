package main

import (
	"encoding/json"
	"net/http"
)

type QuotationClient struct {
	BaseURL string
}

type QuotationJSON struct {
	Values map[string]Quotation `json:"valores"`
}

type Quotation struct {
	Name  string  `json:"nome"`
	Value float32 `json:"valor"`
}

func (client QuotationClient) GetQuotations() map[string]Quotation {
	resp, err := http.Get(client.BaseURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var quotationJSON QuotationJSON
	if err := json.NewDecoder(resp.Body).Decode(&quotationJSON); err != nil {
		panic(err)
	}

	return quotationJSON.Values
}

func NewQuotationClient() QuotationClient {
	return QuotationClient{"http://api.promasters.net.br/cotacao/v1/valores"}
}
