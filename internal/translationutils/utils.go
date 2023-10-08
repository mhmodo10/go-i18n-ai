package translationutils

import (
	"github.com/google/uuid"
	"github.com/mhmodo10/go-i18n-ai/internal/data"
)

// sets translation value to correct attribute in given map and returns a new map
//
// the original map stays unchanged
func AddToMap(t data.Translation, m map[string]interface{}, currentLevel int64) map[string]interface{} {
	mCopy := map[string]interface{}{}
	for k, v := range m {
		value, isMap := m[k].(map[string]interface{})
		if isMap {
			mCopy[k] = AddToMap(t, value, currentLevel+1)
		} else {
			if k != t.AttributeName || t.Level != currentLevel {
				mCopy[k] = v
				continue
			}
			mCopy[k] = t.Translation
		}
	}

	return mCopy
}

// given a map and start level => return an array of single translation objects
func PopulateTranslations(base map[string]interface{}, level int64) []data.Translation {
	translations := []data.Translation{}
	for k, v := range base {
		switch parsedV := v.(type) {
		case map[string]interface{}:
			translations = append(translations, PopulateTranslations(parsedV, level+1)...)
		case string:
			t := data.NewTranslation(uuid.NewString(), parsedV, "", k, level)
			translations = append(translations, *t)

		default:
			continue
		}
	}
	return translations
}

// helper function to find a translation based on given filter function
//
// returns a pointer to the found object and its index
func FindTranslation(translations []data.Translation, filter func(t data.Translation) bool) (*data.Translation, int64) {
	for index, t := range translations {
		match := filter(t)
		if match {
			return &t, int64(index)
		}
	}
	return nil, -1
}

// Copies the base Spec and sets the data of translationSpec in the copy
//
// returns the copy after all changes have been applied
func MergeTranslation(baseSpec data.LanguageSpec, translationSpec data.LanguageSpec) data.LanguageSpec {
	result := baseSpec.Copy()
	for k, v := range translationSpec.Data {
		result.Data[k] = v
	}
	return result
}
