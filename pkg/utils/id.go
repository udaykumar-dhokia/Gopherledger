package utils

import "math/rand/v2"

const (
	max = 999999
	min = 111111
)

func GenerateID() int {
	return rand.IntN(max-min+1) + max
}
