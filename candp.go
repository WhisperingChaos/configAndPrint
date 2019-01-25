package candp

import "fmt"

type Config struct {
	C string
}

type Print struct {
	P string
}

//go:noinline
func (c Config) PrintIt() {
	fmt.Printf("config", c.C)
}
