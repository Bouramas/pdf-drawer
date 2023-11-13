package main

import (
	"flag"
	"fmt"

	"pdf-creator/pkg/certificate"
	"pdf-creator/pkg/invoice"
)

func main() {

	// Parse command-line flags
	name := flag.String("name", "", "the name of the person who completed the course")
	drawCertificate := flag.Bool("certificate", false, "whether to draw a certificate")
	drawInvoice := flag.Bool("invoice", false, "whether to draw a certificate")
	flag.Parse()

	if *drawCertificate {
		certificate.DrawCertificate(name)
	}

	if *drawInvoice {
		lineItems := invoice.SampleLineItems()
		invoice.DrawInvoice(lineItems)
	}

	if !*drawCertificate && !*drawInvoice {
		fmt.Println("Please specify either -certificate or -invoice flag.")
	}

}
