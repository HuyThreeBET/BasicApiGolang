package random

import "simple_shop/db/util"

func RandomUsername() string {
	return util.RandomString(8)
}

func RandomFullName() string {
	return util.RandomString(12)
}

func RandomHashedPassword() string {
	return util.RandomString(12)
}
