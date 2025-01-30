package b

func main(a any) {}

func test() {
	main(nil) // nil passed to interface parameter
}
