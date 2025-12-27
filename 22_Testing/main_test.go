package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 4)
	expected := 9

	if total != expected {
		t.Errorf("Sum was incorrect, got %d expected %d", total, expected)
	}

	// Tabla de resultados
	tables := []struct { // struct anonimo
		a int
		b int
		r int
	}{ // Inicializacion literal
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.r {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.r)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		r int
	}{
		{5, 2, 5},
		{10, 5, 10},
		{-10, 5, 5},
	}

	for _, item := range tables {
		maxValue := Max(item.a, item.b)

		if maxValue != item.r {
			t.Errorf("Max was incorrect, got %d expected %d", maxValue, item.r)
		}
	}
}

func TestFibo(t *testing.T) {
	tables := []struct {
		n int
		r int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fibo := Fibonacci(item.n)
		if fibo != item.r {
			t.Errorf("Error Fibo(%d): Got %d Expected %d", item.n, fibo, item.r)
		}
	}
}
