package main

import (
	pdfgen "github.com/NICKNAME-wengreen/BigDemo/pdfGenerator"
	"fmt"
)

func main() {
	r := pdfgen.NewRequestPdf("")

	templatePath := "public/main.html"

	outputPath   := "public/main.pdf"

	templateData := struct {
		Title 	    string
		Description string
	}{
		Title:	     "TitleAAA",
		Description: "DescriptionBBB",
	}

	if err := r.ParseTemplate(templatePath,templateData); err == nil {
		args := []string{"no-pdf-compression"}

		ok,_ := r.GeneratePDF(outputPath,args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
