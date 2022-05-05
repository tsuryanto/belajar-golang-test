package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test with assert
func TestHelloWithAssert(t *testing.T) {
	result := Hello("Taufiq")
	assert.Equal(t, "Hello Taufiq", result, "Hasil tidak sesuai")
}

// SubTest
func TestHelloWithSubTest(t *testing.T) {
	t.Run("TestHelloTaufiq", func(t *testing.T) {
		result := Hello("Taufiq")
		require.Equal(t, "Hello Taufiq", result, "Hasil tidak sesuai")
	})
	t.Run("TestHelloSri", func(t *testing.T) {
		result := Hello("Sri")
		assert.Equal(t, "Hello Sri", result, "Hasil tidak sesuai")
	})
}

// Table test
// Untuk sekaligus memasukkan multiple data untuk testing
func TestHelloTableTest(t *testing.T) {
	// Slice of Struct - Cara 1

	// type Bio struct {
	// Name    string
	// Request  string
	// Expected string
	// }

	// tests := []Bio{
	// 	{
	// 		Name:    "TestTaufiq",
	// 		Request:  "Taufiq",
	// 		Expected: "Hello Taufiq",
	// 	},
	// 	{
	// 		Name:    "TestSri",
	// 		Request:  "Sri",
	// 		Expected: "Hello Sri",
	// 	},
	// }

	// Slice of Struct - Cara 2

	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "TestTaufiq",
			Request:  "Taufiq",
			Expected: "Hello Taufiq",
		},
		{
			Name:     "TestSri",
			Request:  "Sri",
			Expected: "Hello Sri",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			require.Equal(t, test.Expected, Hello(test.Request), "Data tidak sesuai")
		})
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Taufiq Suryanto")
	}
}

func BenchmarkHelloTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Taufiq)",
			request: "Taufiq",
		},
		{
			name:    "HelloWorld(Sri)",
			request: "Sri",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(benchmark.request)
			}
		})
	}
}
