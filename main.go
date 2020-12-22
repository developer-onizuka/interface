package main
import (
	"fmt"
)

type Crier interface {
	Cry() string
	Echo() string
}
type Duck struct{
	Name string
	Number int
}
type Parrot struct{
	Name string
	Place string
}

func (d Duck) Cry() string {
	return "Quack"
}
func (d Duck) Echo() string {
	return fmt.Sprintf("Name:%s\n", d.Name)
}
func (p *Parrot) Cry() string {
	return "Squawk"
}
func (p *Parrot) Echo() string {
	return fmt.Sprintf("I live in %s.\n", p.Place)
}

func main() {
	var c Crier = Duck{"Jack", 1}
	fmt.Println(c.Cry())
	fmt.Println(c.Echo())

	var p Crier = &Parrot{"xyz", "Tokyo"}
	fmt.Println(p.Cry())
	fmt.Println(p.Echo())
}
