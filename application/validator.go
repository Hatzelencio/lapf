package application

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
	"log"
	"net"
)

const (
	validatorTagCIDRIpv4      = "cidrv4"
	validatorTagOutputFormat  = "output"
	validatorTagCloudProvider = "provider"
)

var validatorFuncMap = map[string]validator.Func{
	validatorTagCIDRIpv4:      isCIDRv4,
	validatorTagOutputFormat:  isOutputFormat,
	validatorTagCloudProvider: isCloudProvider,
}

func init() {
	validate = validator.New()
	for name, fn := range validatorFuncMap {
		if err := validate.RegisterValidation(name, fn); err != nil {
			log.Fatalf("Someting wrong happened: %v", err)
		}
	}
}

func isCIDRv4(fl validator.FieldLevel) bool {
	value, _, _ := fl.ExtractType(fl.Field().Slice(0, fl.Field().Len()))

	for i := 0; i < fl.Field().Len(); i++ {
		ip, _, err := net.ParseCIDR(value.Index(i).String())
		if err != nil || ip.To4() == nil {
			return false
		}
	}
	return true
}

func isOutputFormat(fl validator.FieldLevel) bool {
	return slices.Contains(cliAcceptedOutputFormat, fl.Field().String())
}

func isCloudProvider(fl validator.FieldLevel) bool {
	return slices.Contains(cliAcceptedCloudProvider, fl.Field().String())
}

func newInputValidate(input interface{}) error {
	if err := validate.Struct(input); err != nil {
		return err
	}

	return nil
}
