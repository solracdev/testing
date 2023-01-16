package main

type Foo struct{}

type Bar struct{}

func fooToBar(foo Foo) Bar {
	return Bar{}
}

func convertEmptySlice(foos []Foo) []Bar {
	var bars []Bar
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

func convertGivenCapacity(foos []Foo) []Bar {
	n := len(foos)
	bars := make([]Bar, 0, n)
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

func convertGivenLength(foos []Foo) []Bar {
	n := len(foos)
	bars := make([]Bar, n)
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}
