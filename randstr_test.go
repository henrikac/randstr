package randstr_test

import (
	"testing"

	"github.com/henrikac/randstr"
)

func TestGenerate(t *testing.T) {
	cache := map[string]struct{}{}

	for i := 0; i < 100000; i++ {
		str, err := randstr.Generate()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := cache[str]; ok {
			t.Errorf("%s has already been generated\n", str)
		}
		cache[str] = struct{}{}
	}
}

func TestGenerateLen(t *testing.T) {
	for i := 0; i < 20; i++ {
		str, err := randstr.GenerateLen(i)
		if err != nil {
			t.Fatal(err)
		}
		if len(str) != i {
			t.Errorf("Expected string of size: %d\nGot: %d\n", i, len(str))
		}
	}
}

func TestBase64Encoded(t *testing.T) {
	cache := map[string]struct{}{}

	for i := 0; i < 100000; i++ {
		str, err := randstr.Base64Encoded()
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := cache[str]; ok {
			t.Errorf("%s has already been generated\n", str)
		}
		cache[str] = struct{}{}
	}
}

// -- benchmarking

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randstr.Generate()
	}
}

func BenchmarkGenerateLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randstr.GenerateLen(16)
	}
}

func BenchmarkBase64Encoded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randstr.Base64Encoded()
	}
}
