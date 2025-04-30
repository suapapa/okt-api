package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/suapapa/okt-api/pkg/okt"
)

const (
	// oktBaseURL = "http://127.0.0.1:8080"
	oktBaseURL = "https://homin.dev/okt"
	oktToken   = ""
)

func main() {
	okt, err := okt.New(oktBaseURL, oktToken)
	if err != nil {
		log.Fatal(err)
	}

	meta, body, err := parseHugoPage("sample.md")
	if err != nil {
		log.Fatal(err)
	}

	normalized, err := okt.NormalizeText(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("normalized : ", len(normalized))

	chunks := simpleKoreanSplitter(normalized)
	fmt.Println("chunks : ", len(chunks))

	phrases, err := okt.ExtractPhrases(normalized)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("phrases : ", strings.Join(phrases, ","), len(phrases))

	fmt.Println("meta : ", meta)
}

func parseHugoPage(fp string) (map[string]interface{}, string, error) {
	data, err := os.ReadFile(fp)
	if err != nil {
		panic(err)
	}
	content := string(data)

	// --- 으로 구분
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, "", fmt.Errorf("could not find front matter")
	}

	yamlPart := parts[1]
	markdownPart := strings.TrimSpace(parts[2])

	// YAML 파싱
	var frontMatter map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlPart), &frontMatter)
	if err != nil {
		return nil, "", fmt.Errorf("YAML parsing error: %w", err)
	}

	// 결과 출력
	// fmt.Println("Front Matter (map):")
	// fmt.Println(frontMatter)

	// fmt.Println("\nMarkdown 본문 (string):")
	// fmt.Println(markdownPart)

	return frontMatter, markdownPart, nil
}

func simpleKoreanSplitter(text string) []string {
	// 마침표, 느낌표, 물음표, "다 ", "요 " 기준으로 나눈다
	re := regexp.MustCompile(`([^\n.!?다요]+[다요.!?])`)
	matches := re.FindAllString(text, -1)

	var chunks []string
	for _, match := range matches {
		clean := strings.TrimSpace(match)
		if len(clean) > 0 {
			chunks = append(chunks, clean)
		}
	}
	return chunks
}
