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

func cryecho(x Crier) (string, string) {
	return x.Cry(), x.Echo() 
}

func main() {
	//var c Crier = Duck{"Jack", 1}
	//var c Crier
	//c = Duck{"Jack", 1}
	c, _ := duckInit("Jack",1)
	//fmt.Println(c.Cry())
	//fmt.Println(c.Echo())
	cry1, echo1 := cryecho(c)
	fmt.Printf("%v\n%v\n", cry1, echo1)

	//var p Crier = &Parrot{"Nick", "Tokyo"}
	//var p Crier
	//p = &Parrot{"Nick", "Tokyo"}
	p, _ := parrotInit("Nick", "Tokyo")
	//fmt.Println(p.Cry())
	//fmt.Println(p.Echo())
	cry2, echo2 := cryecho(p)
	fmt.Printf("%v\n%v\n", cry2, echo2)

	//a := &Person{}
	//a.Name = "John"
	//a.Place = "NewYork"
	a := &Person{Parrot{"John", "Newyork"}}
	//fmt.Println(a.Cry())
	//fmt.Println(a.Echo())
	cry3, echo3 := cryecho(a)
	fmt.Printf("%v\n%v\n", cry3, echo3)
}
