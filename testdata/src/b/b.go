package b

type MyInterface interface{}

func main(mi interface{}) {}

func test() {
	main(nil) // want "nil passed to interface parameter"
}
