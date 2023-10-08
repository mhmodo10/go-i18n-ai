package main

import (
	"github.com/mhmodo10/go-i18n-ai/internal/data"
)

// given a base language spec and another language spec
//
// finds all missing attributes in `other` spec compared to `base` spec
//
// returns a language spec with the found differences
func Diff(base data.LanguageSpec, other data.LanguageSpec) data.LanguageSpec {

	oDiff := data.LanguageSpec{
		LanguageCode: other.LanguageCode,
		Language:     other.Language,
		Data:         make(map[string]interface{}),
	}
	for k, v := range base.Data {
		_, exists := other.Data[k]

		if !exists {
			oDiff.Data[k] = v
		} else {
			parsedBaseValue, isBaseMap := v.(map[string]interface{})
			parsedOtherValue, isOtherMap := other.Data[k].(map[string]interface{})
			if !(isOtherMap || isBaseMap) {
				continue
			}

			difference := Diff(data.LanguageSpec{LanguageCode: base.LanguageCode, Data: parsedBaseValue}, data.LanguageSpec{
				LanguageCode: other.LanguageCode,
				Data:         parsedOtherValue,
			})

			for dk, dv := range difference.Data {
				oDiff.Data[dk] = dv
			}
		}
	}
	return oDiff
}
