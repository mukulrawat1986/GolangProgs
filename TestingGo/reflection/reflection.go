package _reflection

func walk(x interface{}, fn func(string)) {
	fn("Hello World")
}
