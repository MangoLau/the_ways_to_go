package even

func IsOdd(i int64) bool {
	mod := i % 2
	return mod == 1
}

func Even(i int) bool {
	return i%2 == 0
}

func Odd(i int) bool {
	return i%2 != 0
}