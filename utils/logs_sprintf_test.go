package utils

import (
	"fmt"
	"testing"
)

func TestLogSprintf(t *testing.T) {
	fmt.Println(Sprintf("something"))
	fmt.Println(Sprintf("some: ", "thing"))
	fmt.Println(Sprintf("%s: ", "some", "thing"))
	fmt.Println(Sprintf("something%"))
	fmt.Println(Sprintf("%s%o%M%E%t%h%i%n%g%%", "s"))
	fmt.Println(Sprintf("something %s sth", "and", "or"))
	fmt.Println(Sprintf("something %s sth"))
}
