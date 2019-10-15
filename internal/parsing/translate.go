package parsing

import (
	"cloud.google.com/go/translate"
	"context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"strings"
)

const MAX_PHRASES_PER_BATCH = 128

func TranslateToEnglish(spanishPhrases []string,
	googleTranslateApiKey string) []string {

	client, err := translate.NewClient(
		context.TODO(),
		option.WithAPIKey(googleTranslateApiKey))
	if err != nil {
		panic(err)
	}

	translations := []translate.Translation{}
	for len(spanishPhrases) > 0 {
		numToRemove := min(MAX_PHRASES_PER_BATCH, len(spanishPhrases))
		batch := spanishPhrases[0:numToRemove]
		spanishPhrases = spanishPhrases[numToRemove:len(spanishPhrases)]

		batchTranslations, err := client.Translate(
			context.TODO(),
			batch,
			language.English,
			&translate.Options{Source: language.Spanish})
		if err != nil {
			panic(err)
		}
		translations = append(translations, batchTranslations...)
	}

	outputs := []string{}
	for _, translation := range translations {
		text := strings.ReplaceAll(translation.Text, "&#39;", "'")
		outputs = append(outputs, text)
	}
	return outputs
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
