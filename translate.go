package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func translate(sourceText, targetLanguage, screenName string) string {
	data, err := loadData()
	if err != nil {
		fmt.Println("Error loading translation data:", err)
		os.Exit(1)
	}

	for _, item := range data.TranslationData {
		if item.Source == sourceText && item.Type == "RdbTextValue" {
			if translations, found := item.Translations[targetLanguage]; found {
				var defaultTranslation *string
				var screenTranslation *string

				for _, translation := range translations {
					if translation.ScreenName != nil && *translation.ScreenName == screenName {
						screenTranslation = &translation.PublishedDst
					}
					if translation.ScreenName == nil {
						defaultTranslation = &translation.PublishedDst
					}
				}

				if screenTranslation != nil {
					return *screenTranslation
				}
				if defaultTranslation != nil {
					return *defaultTranslation
				}
			}
		}
	}

	return sourceText
}

func main() {
	fmt.Println(translate("WOVN.io", "en", "MainActivity"))
	fmt.Println(translate("WOVN.io", "en", ""))
}
