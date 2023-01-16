package main

import (
	"fmt"
	"testing"
)

func TestDataRace(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(idx int) {
			fmt.Println(idx)
		}(i)
	}
}

func TestComplexDataRace(t *testing.T) {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
