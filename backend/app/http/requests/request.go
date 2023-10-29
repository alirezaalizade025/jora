package request

import (
	"errors"
	"jora/app/models/attendance"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request interface {
	bindValue(c *gin.Context)
}

func Validation(c *gin.Context, request request) bool {

	// create new validator
	validation := validator.New()
	RegisterCustomValidator(validation)

	// bind value to struct
	request.bindValue(c)

	// do validation by validator tag
	err := validation.Struct(request)

	// handle error message
	if err != nil {
		hasError := handleErrorMessage(request, err, c)
		if !hasError {
			return false
		}
	}

	return true
}

func handleErrorMessage(request request, err error, c *gin.Context) bool {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make(map[string]([]string), len(ve))
		// out := make([]ApiError, len(ve))
		for _, fe := range ve {

			jsonName := fieldJsonName(request, fe.Field())

			// out[i] = ApiError{jsonName, msgForTag(request, fe)}
			out[jsonName] = []string{msgForTag(request, fe)}
		}

		// todo: log validation errors
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": out,
		})
		return false

	}
	return true
}

func msgForTag(request request, tag validator.FieldError) string {

	field := fieldJsonName(request, tag.Field())

	switch tag.Tag() {
	case "required":
		return field + " is required"
	case "gte":
		return field + " must be greater than or equal to " + tag.Param()
	case "lte":
		return field + " must be less than or equal to " + tag.Param()
	case "oneof":
		return field + " must be one of " + tag.Param()
	case "datetime":
		return field + " must be a valid datetime"
	case "in":
		return field + " is not valid" // todo: dynamic message
	}
	return tag.Error()
}

func fieldJsonName(req request, name string) string {

	field, ok := reflect.TypeOf(req).Elem().FieldByName(name)
	if !ok {
		// todo: log it
		return name
	}

	return field.Tag.Get("json")
}


func RegisterCustomValidator(v *validator.Validate) error {
	return v.RegisterValidation("in", in)
}


//==============================================================================//
//
//                                Custom validator								//
//
//==============================================================================//
func in(fl validator.FieldLevel) bool {

	fieldValue := fl.Field().String()
	// Get the argument passed to the custom validator
	arg := fl.Param()

	// map if arg to list
	var list map[int]string

	switch arg {
	case "attendanceType":
		list = attendance.TYPE_MAP()
	default:
		return false
	}

	// Iterate over the list to find a match
	for _, v := range list {
		if v == fieldValue {
			return true
		}
	}


	// Return the result of the validation
	return false

}
