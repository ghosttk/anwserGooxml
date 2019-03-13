package main

import (
	"fmt"
	"log"
    "regexp"

	"baliance.com/gooxml/document"
)

func main() {
	doc, err := document.Open("document.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

    regexpNumber:= regexp.MustCompile(`\d\.$`)
    //regexpAnswer:= regexp.MustCompile(`【答案】`)

	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	for _, p := range paragraphs {
            /*
		for _, r := range p.Runs() {
                m := regexpNumber.MatchString(r.Text())
                if (m) {
                        fmt.Println(r.Text())
                        //n := regexpAnswer.MatchString(
                }
		}
            */
            rs := p.Runs()
            i := 0
            for i=0; i<len(rs); i++ {
                r := rs[i]
                m := regexpNumber.MatchString(r.Text())
                if (m) {
                    fmt.Println(r.Text())
                }

            }
	}

	doc.SaveToFile("edit-document.docx")
}
