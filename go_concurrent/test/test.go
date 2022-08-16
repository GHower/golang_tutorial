package main

type Animal interface {
	Speak() string
}
type Dog struct {
}
type Cat struct {
}

func (d *Dog) Speak() string {
	return "woof"
}

func (c *Cat) Speak() string {
	return "meow"
}
func main() {
	animals := []Animal{&Dog{}, &Cat{}}
	for _, animal := range animals {
		println(animal.Speak())
	}
}
