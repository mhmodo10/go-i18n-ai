package data

// translation type to hold single translation text
type Translation struct {
	ID            string `json:"id"`
	Base          string `json:"base"`
	Translation   string `json:"translation"`
	AttributeName string `json:"attributeName"`
	Level         int64  `json:"level"`
}

// creates new translation and returns its pointer
func NewTranslation(id string, base string, translation string, attributeName string, level int64) *Translation {
	return &Translation{
		ID:            id,
		Base:          base,
		Translation:   translation,
		AttributeName: attributeName,
		Level:         level,
	}
}

// translations for a certain locale (e.g. en)
type LocaleTranslations struct {
	Locale       string        `json:"locale"`
	Translations []Translation `json:"translations"`
}

// language spec type holds information on a language config file
type LanguageSpec struct {
	Language     string
	LanguageCode string
	Data         map[string]interface{}
}

// returns a copy of the language spec
func (ls *LanguageSpec) Copy() LanguageSpec {
	lsCopy := LanguageSpec{
		Language:     ls.Language,
		LanguageCode: ls.LanguageCode,
		Data:         map[string]interface{}{},
	}
	for k, v := range ls.Data {
		lsCopy.Data[k] = v
	}
	return lsCopy
}
