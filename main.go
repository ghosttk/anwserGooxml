package main

import (
	"fmt"
	"log"
    "regexp"
    "io/ioutil"

	"baliance.com/gooxml/document"
)

func main() {
	doc, err := document.Open("document.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

    regexpNumber:= regexp.MustCompile(`\d\.$`)
    regexpAnwser:= regexp.MustCompile(`【答案】`)
    regexpJieXi:= regexp.MustCompile(`【解析】`)

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
            bstart := false
            var anwserSlice []string
            for i=0; i<len(rs); i++ {
                r := rs[i]
                m := regexpNumber.MatchString(r.Text())
                if (m) {
                    rn := rs[i+1]
                    n := regexpAnwser.MatchString(rn.Text())
                    if (n) {
                            bstart = true
                    }
                }
                bjieXi := regexpJieXi.MatchString(r.Text())
                if (bjieXi) {
                        bstart = false
                }
                if (bstart) {
                        data := []byte(r.Text())
                        err := ioutil.WriteFile("anwser.txt", data, 0644)
                        anwserSlice = append(anwserSlice, r.Text())
                } 
            }
            fmt.Println(anwserSlice)
	}

	doc.SaveToFile("edit-document.docx")
}
