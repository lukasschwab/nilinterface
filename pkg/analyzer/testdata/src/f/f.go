package f

type Vehicle interface {
	Honk() error
}

type Car struct{}

func (*Car) Honk() error {
	return nil
}

func angryDriver(v Vehicle) {
	for i := 0; i < 10; i++ {
		_ = v.Honk()
	}
}

func simulateTraffic() {
	angryDriver(nil) // want "nil passed to interface parameter"

	// Not a lint error.
	var c *Car
	angryDriver(c)
}
