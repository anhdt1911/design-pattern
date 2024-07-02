# How to implement

## 1. Specify abstract product: define an interfaces with common methods
```
type Button interface {
	onClick()
}
```

## 2. Specify abstract factory: define an interfaces with creation methods
```
type Factory interface {
    // the factory method should have return type as abstract product
	newButton() Button
}
```

## 3. Specify concrete products and factories

Note: **Concrete factories' creation method should also define return type as abstract product**

```
// concrete product
type WindownButton struct{}

func (b *WindownButton) onClick() {
	fmt.Println("Clicks on windows")
}
```
```
// concreate creator
type WindownDialog struct {
}

// return the interface
func (wd WindownDialog) newButton() Button {
	return &WindownButton{}
}
```

## 4. Consume with client 

The client application will need to work with abstract type, only specify the concrete factory type that client needs
```
type Application struct {
	d Dialog
}

func newApp(d Dialog) *Application {
	return &Application{d}
}

func main() {
    // specify concrete type, easily to switch to another factory
	d := WindownDialog{}
	app := newApp(d)
	app.d.newButton().onClick()
}
```