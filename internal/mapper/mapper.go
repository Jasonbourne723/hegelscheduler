package mapper

import "github.com/jinzhu/copier"

func Map(toValue interface{}, fromValue interface{}) error {
	return copier.Copy(toValue, fromValue)
}
