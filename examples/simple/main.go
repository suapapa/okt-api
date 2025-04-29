package main

import (
	"fmt"
	"log"

	"github.com/suapapa/okt-go/pkg/okt"
)

func main() {
	okt, err := okt.NewOKT("https://homin.dev/okt")
	// okt, err := okt.NewOKT("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	normalized, err := okt.NormalizeText("안녕하세요")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>", normalized)

	phrases, err := okt.ExtractPhrases("키스의 조건은 눈을 감아야 한다.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>", phrases)
}
