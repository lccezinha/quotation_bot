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

	buffer.WriteString("\t\tCotações\t\t\n\n")

	for _, quotation := range formatter.message {
		buffer.WriteString(fmt.Sprintf("1 %s = R$ %.2f\n", quotation.Name, quotation.Value))
	}

	return buffer.String()
}
