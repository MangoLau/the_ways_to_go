package even

func IsOdd(i int64) bool {
	mod := i % 2
	return mod == 1
}