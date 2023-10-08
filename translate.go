package main

import (
	"encoding/json"
	"errors"

	"github.com/mhmodo10/go-i18n-ai/internal/clients"
	"github.com/mhmodo10/go-i18n-ai/internal/data"
	"github.com/mhmodo10/go-i18n-ai/internal/prompts"
	"github.com/mhmodo10/go-i18n-ai/internal/translationutils"
)

// this file will include functionality for generating different translations of
// given json file
// the function will take base map[string]interface and a slice of
// map[string]interface (one for each locale)
// it will call openai api using the client and return a slice with the translations
func Translate(base data.LanguageSpec, target data.LanguageSpec) (data.LanguageSpec, error) {
	targetCopy := target.Copy()

	// skip if there's nothing to translate
	if len(target.Data) == 0 {
		return targetCopy, nil
	}

	openai := clients.NewOpenAIClient(ApiKey)

	localeTranslations := translationutils.PopulateTranslations(targetCopy.Data, 0)
	b := data.ChatCompletionBody{
		Model: Model,
		Messages: []data.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompts.TranslationBasePrompt(base.Language, target.Language),
			},
			{
				Role:    "system",
				Content: "Do not add any text that was not in the original text",
			},
			{
				Role:    "user",
				Content: prompts.TranslationTextPrompt(localeTranslations),
			},
		},
		Temperature: 0,
	}
	res, err := openai.ChatCompletion(b)
	if err != nil {
		return targetCopy, err
	}
	if res.Error != nil {
		return targetCopy, errors.New(res.Error.String())

	}
	translated := []data.Translation{}
	err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &translated)
	if err != nil {
		return targetCopy, err
	}

	for _, translation := range translated {
		targetTranslation, _ := translationutils.FindTranslation(localeTranslations, func(t data.Translation) bool {
			return t.ID == translation.ID
		})
		if targetTranslation == nil {
			continue
		}
		targetCopy.Data = translationutils.AddToMap(data.Translation{
			ID:            targetTranslation.ID,
			Base:          targetTranslation.Base,
			Translation:   translation.Translation,
			Level:         targetTranslation.Level,
			AttributeName: targetTranslation.AttributeName,
		}, targetCopy.Data, 0)
	}
	return targetCopy, nil
}
