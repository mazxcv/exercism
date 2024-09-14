package collatzconjecture

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	count := 0
	if n == 0 {
		return 0, errors.New("zero value")
	}
	if n < 0 {
		return 0, errors.New("negative value")
	}
	for n != 1 {
		fmt.Println(count, n)
		count++
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return count, nil
}
