package hello

import (
	quotev3 "rsc.io/quote/v3"
)

func Hello() string {
	return quotev3.HelloV3()
}
func Proverb() string {
	return quotev3.Concurrency()
}
