package e

func main(a any) {}

func test() {
	// NOTE: I may want this to *not* be an error; see bc0099.
	main(nil) // want "nil passed to interface parameter"
}
