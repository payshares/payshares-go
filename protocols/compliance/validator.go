package compliance

import (
	"github.com/asaskevich/govalidator"
	"github.com/payshares/go/address"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("payshares_address", govalidator.CustomTypeValidator(isPaysharesAddress))
}

func isPaysharesAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)

	if err == nil {
		return true
	}

	return false
}
