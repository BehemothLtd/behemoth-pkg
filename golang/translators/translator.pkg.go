package translator

import (
	"fmt"
	"reflect"
)

var ClientLanguage string

func Translate(language *string, constantName string, args ...interface{}) string {
	if ClientLanguage == "" {
		ClientLanguage = "vi"
	}

	lang := ClientLanguage

	if language != nil {
		lang = *language
	}

	translations, found := supportedLanguages[lang]
	if !found {
		return "unsupported language"
	}

	// Using reflection to dynamically access the fields of the struct
	value := reflect.ValueOf(translations).FieldByName(constantName)
	if !value.IsValid() {
		return constantName
	}

	message := value.String()
	return fmt.Sprintf(message, args...)
}
