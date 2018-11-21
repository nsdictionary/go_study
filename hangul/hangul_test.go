package hangul

import (
	"fmt"
	// "testing"
)

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// true
	// false
	// true
}

// func TestHasConsonantSuffix(t *testing.T) {
// v0 := HasConsonantSuffix("Go 언어")
// v1 := HasConsonantSuffix("그럼")
// v2 := HasConsonantSuffix("우리 밥 먹고 합시다.")

// if v0 != false {
// t.Fatal("v0")
// }

// if v1 != true {
// t.Fatal("v1")
// }

// if v2 != false {
// t.Fatal("v1")
// }
// }

// func TestPrintBytes(t *testing.T) {
// s := "가나다"
// for i := 0; i < len(s); i++ {
// fmt.Printf("%x:", s[i])
// }
// fmt.Println()
// }

// func TestPrintBytes2(t *testing.T) {
// s := "가나다"
// fmt.Printf("%x\n", s)
// fmt.Printf("% x\n", s)
// }

// func TestModifyBytes(t *testing.T) {
// b := []byte("가나다") // byte slice
// b[2]++
// fmt.Println(string(b))

// if string(b) != "각나다" {
// t.Fatal("fail modify byte")
// }
// }
