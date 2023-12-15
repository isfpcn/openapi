/**
 * @author   Liang
 * @create   2023/11/30 13:38
 * @version  1.0
 */

package validate

import (
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"openapi/internal/code"

	zh2 "github.com/go-playground/locales/zh"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gogo/protobuf/proto"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	tans     ut.Translator

	ValidateHandler = func(resp proto.Message) error {
		err := validate.Struct(resp)
		if err != nil {
			errs := err.(validator.ValidationErrors)
			for _, v := range errs.Translate(tans) {
				return code.Error(code.InvalidArgument, v)
			}
		}
		return nil
	}
)

func New() error {

	zh := zh2.New()
	uni = ut.New(zh)
	validate = validator.New()
	tans, _ = uni.GetTranslator("zh")
	err := zh_translations.RegisterDefaultTranslations(validate, tans)
	return err
}
