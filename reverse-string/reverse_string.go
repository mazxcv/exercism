package reverse

func Reverse(input string) (reverse string) {
	for _, v := range input {
		reverse = string(v) + reverse
	}

	return
}
