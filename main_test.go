package main

import "testing"

func Test1(t *testing.T) {
	t.Log("Test1 START")
	var xs, ys []int
	t.Run("xs", func(t *testing.T) {
		t.Log(xs[0]) // panic here
	})
	t.Run("ys", func(t *testing.T) {
		if len(ys) == 0 {
			t.Fatal("len ys = 0")
		}
		t.Log(ys[0])
	})
	t.Log("Test1 END")
}

func Test2(t *testing.T) {
	t.Log("Test2 START")
	t.Log("Test2 END")
}
