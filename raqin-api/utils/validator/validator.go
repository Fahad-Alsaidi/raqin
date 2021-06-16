package validator

// the code in this file got help from
// https://hackwild.com/article/go-input-validation-and-testing/

import (
	"raqin-api/utils/irror"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)

	var exist bool
	trans, exist = uni.GetTranslator("en")
	if !exist {
		panic("translator does not exist")
	}

	validate = validator.New()
	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

// Validate attempts to validate the structs's values.
func Validate(s interface{}) error {
	if err := validate.Struct(s); err != nil {

		// var er = irror.New("validation failed")
		var er = irror.New("validation failed")
		if failures, ok := err.(validator.ValidationErrors); ok {
			for _, e := range failures {
				// can translate each error one at a time.
				er = er.AppendDetail(e.Translate(trans))
			}
		}
		return er
	}
	return nil
}
