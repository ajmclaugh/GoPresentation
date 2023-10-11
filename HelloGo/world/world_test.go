package HelloGo

import (
	"regexp"
	"testing"
)

func TestAddWorldPointer(t *testing.T) {
	hello_point := "Hello"
	want := regexp.MustCompile("Hello World")
	AddWorldPointer(&hello_point)
	if !want.MatchString(hello_point) {
		t.Fatalf("Expected  %#q, got %s ", want, hello_point)
	}

}

func BenchmarkAddWorld(b *testing.B) {
	hello := "Hello"
	AddWorld(hello)

}
func BenchmarkAddWorldPointer(b *testing.B) {
	hello_point := "Hello"
	AddWorldPointer(&hello_point)

}

func TestAddWorld(t *testing.T) {
	hello := AddWorld("Hello")
	want := regexp.MustCompile("Hello World")
	if !want.MatchString(hello) {
		t.Fatalf("Expected  %#q, got %s ", want, hello)
	}

}
