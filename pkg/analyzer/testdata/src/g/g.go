package g

func g(f func()) {
	f()
}

func main() {
	g(nil) // want "nil passed to interface parameter"
}
