package main

import "fmt"

// abstract product 1: Chair
type Chair interface {
	sitOn()
}

// concrete product: Sofa
type Sofa struct {
}

func (so *Sofa) sitOn() {
	fmt.Println("so comfy")
}

// concrete product: Stool
type Stool struct {
}

func (st *Stool) sitOn() {
	fmt.Println("hard on my ass")
}

// abstract product 2: Table
type Table interface {
	putThingsOn()
}

// Concrete product 2: Coffe Table
type CoffeeTable struct{}

func (ct *CoffeeTable) putThingsOn() {
	fmt.Println("put your coffee here")
}

// Concrete product 2: Desk
type Desk struct{}

func (d *Desk) putThingsOn() {
	fmt.Println("put your books here")
}

// Abstract factory with creation method for each product
type FurnitureFactory interface {
	newChair() Chair
	newTable() Table
}

// concrete factory for create concrete product variants: Casual
type CasualFurnitureFactory struct{}

func (mf *CasualFurnitureFactory) newChair() Chair {
	return &Stool{}
}
func (mf *CasualFurnitureFactory) newTable() Table {
	return &Desk{}
}

// concrete factory for create concrete product variants: Fancy
type FancyFurnitureFactory struct{}

func (ff *FancyFurnitureFactory) newChair() Chair {
	return &Sofa{}
}
func (ff *FancyFurnitureFactory) newTable() Table {
	return &CoffeeTable{}
}

// Client who uses the factories
type Application struct {
	//references to the factories is all abstract
	f FurnitureFactory
}

func newApp(f FurnitureFactory) *Application {
	return &Application{
		f,
	}
}

func (a *Application) useTheFurniture() {
	c := a.f.newChair()
	t := a.f.newTable()
	c.sitOn()
	t.putThingsOn()
}

func main() {
	app := newApp(&FancyFurnitureFactory{})
	app.useTheFurniture()
}
