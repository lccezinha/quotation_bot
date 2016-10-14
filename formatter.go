package main

import (
	"bytes"
	"fmt"
)

type Formatter struct {
	message map[string]Quotation
}

func (formatter Formatter) Format() string {
	var buffer bytes.Buffer

	for _, quotation := range formatter.message {
		buffer.WriteString(fmt.Sprintf("Cotação %s em relação ao Real: R$ %.2f = 1 %s\n", quotation.Name, quotation.Value, quotation.Name))
	}

	return buffer.String()
}
