package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pjchender/todo-mvc-backend/pkg/errMsg"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"unicode"
)

// 當 gin Context 中有 Errors 時，會進到此 middleware 處理
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				switch e.Type {
				case gin.ErrorTypeBind:

					var validationErrs validator.ValidationErrors
					if ok := errors.As(e.Err, &validationErrs); !ok {
						writeError(c, e.Error())
						return
					}

					var stringErrors []string
					for _, err := range validationErrs {
						stringErrors = append(stringErrors, validationErrorToText(err))
					}
					writeError(c, strings.Join(stringErrors, "; "))
				default:
					log.Println("unknown error type: ", e.Type)
					writeError(c, e.Err.Error())
				}
			}
		}
	}
}

func validationErrorToText(e validator.FieldError) string {
	runes := []rune(e.Field())
	runes[0] = unicode.ToLower(runes[0])
	fieldName := string(runes)
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("Field '%s' is required", fieldName)
	case "max":
		return fmt.Sprintf("Field '%s' must be less or equal to %s", fieldName, e.Param())
	case "min":
		return fmt.Sprintf("Field '%s' must be more or equal to %s", fieldName, e.Param())
	}
	return fmt.Sprintf("Field '%s' is not valid", fieldName)
}

func writeError(ctx *gin.Context, errString string) {
	statusCode := http.StatusBadRequest
	if ctx.Writer.Status() != http.StatusOK {
		statusCode = ctx.Writer.Status()
	}

	ctx.JSON(statusCode, &errMsg.Error{
		Error:       http.StatusText(statusCode),
		StatusCode:  statusCode,
		Description: errString,
	})
}
