package validate

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	id "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func InitTranslator() {
	id := id.New()
	uni = ut.New(id, id)

	trans, _ = uni.GetTranslator("id")
	id_translations.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans)
}

func ExtractValidationErrors(err error) map[string]string {
	errFields := make(map[string]string)

	errVal, ok := err.(validator.ValidationErrors)
	if !ok {
		errFields["error"] = "invalid format validation"
	}

	for _, e := range errVal {
		errFields[strings.ToLower(e.Field())] = e.Translate(trans)
	}

	return errFields
}
