package prompts

import (
	"fmt"

	"github.com/mhmodo10/go-i18n-ai/internal/data"
)

// Base prompt to be passed as system prompt to openai api
func TranslationBasePrompt(baseLang string, targetLang string) string {
	basePrompt := `I will provide you with text in %s, I want you to translate it to %s.
	The provided text will be used for configuration files used by react-intl library.
	This means that you should keep all variables and special formats unchanged.
	Each translation object will contains the following properties: 
	Id: an identifier to identify this text
	text: the text to be translated.
	I want you to return an array of json objects with the following attributes: 
	Id : the Id sent in the prompt request
	translation: the translation of the text that was sent in the prompt request.
	Here are the Rules you must follow:
	1- Do not add any text or punctuation that was not in the original text.
	2- It is extremely important to only return the results array of json object with the requested attributes.
	3- Do not add any dots, commas or any other punctuation characters`
	filledPrompt := fmt.Sprintf(basePrompt, baseLang, targetLang)
	return filledPrompt
}

// dynamic prompt to provide the text to be translated
func TranslationTextPrompt(translations []data.Translation) string {
	basePrompt := "Here are the text objects to be translated:\n"
	for _, t := range translations {
		textPrompt := fmt.Sprintf("Id: %s text: %s \n", t.ID, t.Base)
		basePrompt += textPrompt
	}
	return basePrompt
}
