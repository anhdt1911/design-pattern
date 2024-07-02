**Abstract factory is nearly the same as Factory method**
**The only diff is abstract factory works with many products and each product have many variants, so a abstract factory have many method for creating types of product, and concrete factory specify the product variants**
```
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
```

## Application code also need to works with abstract type
```
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
```