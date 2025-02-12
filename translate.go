package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type Translation struct {
	PublishedDst string  `json:"published_dst"`
	ScreenName   *string `json:"screen_name"`
}

type TranslationItem struct {
	Source       string                   `json:"source"`
	Type         string                   `json:"type"`
	Translations map[string][]Translation `json:"translations"`
}

type Data struct {
	TranslationData []TranslationItem `json:"translation_data"`
}

func loadData() (Data, error) {
	file, err := os.ReadFile("data.json")
	if err != nil {
		return Data{}, err
	}
	var data Data
	err = json.Unmarshal(file, &data)
	return data, err
}

func matchPattern(sourceText string, patternText string) bool {
	pattern := regexp.MustCompile(`\{\{number:(\d+)\}\}`)
	source := pattern.ReplaceAllString(patternText, `(\d+)`)

	pattern = regexp.MustCompile(`\{\{text\[(\d+),(\d+)\]:(\d+)\}\}`)
	source = pattern.ReplaceAllString(source, `(.{${1},${2}})`)

	patternRegex := regexp.MustCompile("^" + source + "$")
	return patternRegex.MatchString(sourceText)
}

func translate(sourceText, targetLanguage, screenName string) string {
	data, err := loadData()
	if err != nil {
		return sourceText
	}

	for _, item := range data.TranslationData {
		if item.Source == sourceText && item.Type == "RdbTextValue" {
			if translations, found := item.Translations[targetLanguage]; found {
				for _, translation := range translations {
					if translation.ScreenName == nil || *translation.ScreenName == screenName {
						return translation.PublishedDst
					}
				}
			}
		}
		if item.Type == "RdbPatternValue" && matchPattern(sourceText, item.Source) {
			if translations, found := item.Translations[targetLanguage]; found {
				for _, translation := range translations {
					if translation.ScreenName == nil || *translation.ScreenName == screenName {
						return translation.PublishedDst
					}
				}
			}
		}
	}

	return sourceText
}

func main() {
	//Examples in English, Vietnamnese and one Non-Existing
	fmt.Println(translate("データセーフティ", "en", "Screen4"))
	fmt.Println(translate("データセーフティ", "vi", "Screen4"))
	fmt.Println(translate("Non-existing text", "en", ""))
}
