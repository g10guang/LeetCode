package main

import "fmt"

// 字符串转整形
// 声明：s 仅支持是合法的整型，对于其他的输入返回一个 error
func Atoi(s string) (v int64, err error) {
	originString := s
	if len(s) == 0 {
		err = fmt.Errorf("input string is empty")
		return
	}
	isNagative := false
	// 支持正负数
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			isNagative = true
		}
		s = s[1:]
	}
	if len(s) > 1 && s[0] == '0' {
		err = fmt.Errorf("invalid input string=%s, not allow begin with 0", originString)
		return
	}
	// 接下来的输入仅希望是 [0-9] 输入
	for len(s) > 0 {
		char := s[0]
		n := char - '0'
		if n < 0 || n > 9 {
			err = fmt.Errorf("invalid integer string=%s, unknown char=%v", originString, string([]byte{char}))
			return
		}
		nextVal := v*10 + int64(n)
		if nextVal < v {
			// overflow
			err = fmt.Errorf("overflow")
			return
		}
		v = nextVal
		s = s[1:]
	}
	if isNagative {
		v = v * -1
	}
	return
}

func main() {
	test1()
}

func test1() {
	v, err := Atoi("12340")
	fmt.Printf("v=%v err=%v\n", v, err)

	v, err = Atoi("-12340")
	fmt.Printf("v=%v err=%v\n", v, err)

	v, err = Atoi("1394949949494949949499494949")
	fmt.Printf("v=%v err=%v\n", v, err)

	v, err = Atoi("0")
	fmt.Printf("v=%v err=%v\n", v, err)

	v, err = Atoi("-0")
	fmt.Printf("v=%v err=%v\n", v, err)

	v, err = Atoi("0123")
	fmt.Printf("v=%v err=%v\n", v, err)

	v, err = Atoi("482p0")
	fmt.Printf("v=%v err=%v\n", v, err)
}
