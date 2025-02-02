package d

type Stringy interface {
	String() string
}

type StringyInt int

func (i StringyInt) String() string { return "" }

func main(mi1, mi2 Stringy) {}

func test() {
	main(nil, nil)           // want "nil passed to interface parameter" "nil passed to interface parameter"
	main(StringyInt(1), nil) // want "nil passed to interface parameter"
	main(nil, StringyInt(1)) // want "nil passed to interface parameter"
	main(StringyInt(1), StringyInt(1))
}
