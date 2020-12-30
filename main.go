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
type Person struct {
	Parrot
}
//---
func (d Duck) Cry() string {
	return "Quack"
}
func (d Duck) Echo() string {
	return fmt.Sprintf("Name:%s\n", d.Name)
}
//---
func (p *Parrot) Cry() string {
	return "Squawk"
}
func (p *Parrot) Echo() string {
	return fmt.Sprintf("%v lives in %s.\n", p.Name, p.Place)
}
//---
func (p *Person) Cry() string {
	return "Human!"
}
func (p *Person) Echo() string {
	return fmt.Sprintf("%s used to live in %s.\n", p.Name, p.Place)
}
//---

func duckInit(x string, y int) (Crier, int) {
	var c Crier = Duck{x, y}
	var err int = 0
	return c, err
}

func parrotInit(x string, y string) (Crier, int) {
	var p Crier = &Parrot{x, y}
	var err int = 0
	return p, err
}
//---

func main() {
	//var c Crier = Duck{"Jack", 1}
	//var c Crier
	//c = Duck{"Jack", 1}
	c, _ := duckInit("Jack",1)
	fmt.Println(c.Cry())
	fmt.Println(c.Echo())

	//var p Crier = &Parrot{"Nick", "Tokyo"}
	//var p Crier
	//p = &Parrot{"Nick", "Tokyo"}
	p, _ := parrotInit("Nick", "Tokyo")
	fmt.Println(p.Cry())
	fmt.Println(p.Echo())

	//a := &Person{}
	//a.Name = "John"
	//a.Place = "NewYork"
	a := &Person{Parrot{"John", "Newyork"}}
	fmt.Println(a.Cry())
	fmt.Println(a.Echo())
}
