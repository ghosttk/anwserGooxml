package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"baliance.com/gooxml/document"
)

func main() {
	doc, err := document.Open("document.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	regexpNumber := regexp.MustCompile(`\d\.$`)
	regexpAnwser := regexp.MustCompile(`【答案】`)
	regexpJieXi := regexp.MustCompile(`【解析】`)

	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	var anwserSlice []string
    strAnwser := ""
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
		for i = 0; i < len(rs); i++ {
			r := rs[i]
			m := regexpNumber.MatchString(r.Text())
			if m {
				rn := rs[i+1]
				n := regexpAnwser.MatchString(rn.Text())
				if n {
					bstart = true
				}
			}
			bjieXi := regexpJieXi.MatchString(r.Text())
			if bjieXi {
				bstart = false
                strAnwser += "\n"
			}
			if bstart {
				anwserSlice = append(anwserSlice, r.Text())
				strAnwser += r.Text()
			}
		}
		//		fmt.Println(anwserSlice)
	}
	fmt.Println(strAnwser)
	data := []byte(strAnwser)
	err = ioutil.WriteFile("anwser.txt", data, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	doc.SaveToFile("edit-document.docx")
}
