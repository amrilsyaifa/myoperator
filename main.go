package myoperator

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	constant "github.com/amrilsyaifa/myoperator/constants"
)

type ResponseOperator struct {
	Card     string `json:"card,omitempty" binding:"required"`
	Operator string `json:"operator,omitempty" binding:"required"`
	Code     string `json:"code,omitempty" binding:"required"`
	Message  string `json:"message,omitempty" binding:"required"`
	IsValid  bool   `json:"is_valid,omitempty" binding:"required"`
}

func GetOperator(phoneNumber string, validate bool) (*ResponseOperator, error) {
	var newPhoneNumber string = phoneNumber
	var card string
	var code int
	var operator string
	var codeOperator string
	var message string = string(constant.INVALID)
	var IsValid bool = false

	// Empty Phone Number Validation
	if len(phoneNumber) <= 0 {
		return nil, fmt.Errorf("please provide phone number")
	}

	// + prepend
	if string(phoneNumber[0]) == "+" {
		newPhoneNumber = phoneNumber[1:]
	}

	// remove char -
	if strings.Contains(newPhoneNumber, "-") {
		newPhoneNumber = strings.ReplaceAll(newPhoneNumber, "-", "")
	}

	// validate only numeric
	re := regexp.MustCompile(`^[0-9]+$`)
	isNumeric := re.MatchString(newPhoneNumber)
	if !isNumeric {
		return nil, fmt.Errorf("please provide correct phone number")
	}

	// country code
	if string(newPhoneNumber[0:2]) == "62" {
		newPhoneNumber = fmt.Sprintf("%s%s", "0", string(newPhoneNumber[2:]))
	}

	if string(newPhoneNumber[0:2]) == "08" {
		num, err := strconv.Atoi(newPhoneNumber[2:4])
		if err != nil {
			return nil, fmt.Errorf("something went wrong")
		}
		code = num
	}

	for _, opr := range constant.Operators {
		if slices.Contains(opr.Code, code) {
			if validate {
				maximum := int(constant.MaximumLength)
				minimum := int(constant.MinimumLength)
				if opr.ValidationConfig != nil {
					if opr.ValidationConfig.MaxLength > int(constant.MaximumLength) {
						maximum = opr.ValidationConfig.MaxLength
					}
				}
				if len(newPhoneNumber) < minimum || len(newPhoneNumber) > maximum {
					return nil, fmt.Errorf("phone number %s length must be between %s and %s", newPhoneNumber, strconv.Itoa(minimum), strconv.Itoa(maximum))
				}
			}

			operator = opr.Operator
			card = opr.Name
			message = string(constant.VALID)
			IsValid = true
			codeOperator = opr.Key
			fmt.Printf("%#v\n", operator)
			break
		}
	}

	result := &ResponseOperator{
		Card:     card,
		Operator: operator,
		Code:     codeOperator,
		Message:  message,
		IsValid:  IsValid,
	}
	return result, nil
}
