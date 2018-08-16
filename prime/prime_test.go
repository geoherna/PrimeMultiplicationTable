package prime

import (
	"testing"
	)

func TestIsPrime(t *testing.T){
	expected := true
	actual := IsPrime(43)
	if actual != expected {
		t.Error("Test failed")
	}
}

func TestGetPrimeNumbers(t *testing.T) {
	expected := 4
	actual := len(GetPrimeNumbers(10))
	if actual != expected {
		t.Error("Test failed")
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPrime(n)
	}
}

func BenchmarkGetPrimeNumbers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPrimeNumbers(20)
	}
}