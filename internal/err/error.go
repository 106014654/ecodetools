package err

import "fmt"

func ErrIndexOutRange(length int, index int) error {
	return fmt.Errorf("下标超出范围，长度 %d, 下标 %d", length, index)
}
