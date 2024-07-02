package main

import "fmt"

type Engine string
type GPS string

const (
	SportEngine Engine = "sport"
	SUVEngine   Engine = "suv"

	GoogleGPS GPS = "good gps"
	AppleGPS  GPS = "suck gps"
)

// product produce by the builder
type Car struct {
	SeatNr int
	Engine Engine
	GPS    GPS
}

// builder abstraction
type Builder interface {
	reset()
	setSeat(seatNr int)
	setEngine(engine Engine)
	setGPS(gps GPS)
}

// concreate builder to build car
type CarBuilder struct {
	Car Car
}

func (cb *CarBuilder) reset() {
	cb.Car = Car{}
}

func (cb *CarBuilder) setSeat(seatNr int) {
	cb.Car.SeatNr = seatNr
}

func (cb *CarBuilder) setEngine(engine Engine) {
	cb.Car.Engine = engine
}

func (cb *CarBuilder) setGPS(gps GPS) {
	cb.Car.GPS = gps
}

// concrete builder is expect to have their own method for getting product
// because various builder may create various products that don't all follow same interface(at least in static type)
func (cb *CarBuilder) getProduct() Car {
	product := cb.Car
	cb.reset()
	return product
}

type CarManual struct {
	Instruction string
}

// concreate builder to build car manual
type CarManualBuilder struct {
	CarManual CarManual
}

func (cb *CarManualBuilder) reset() {
	cb.CarManual = CarManual{
		Instruction: "\n",
	}
}

func (cb *CarManualBuilder) setSeat(seatNr int) {
	cb.CarManual.Instruction = fmt.Sprintf("%sThis car has %d seats\n", cb.CarManual.Instruction, seatNr)
}

func (cb *CarManualBuilder) setEngine(engine Engine) {
	cb.CarManual.Instruction = fmt.Sprintf("%sThis car has %s engine\n", cb.CarManual.Instruction, engine)
}

func (cb *CarManualBuilder) setGPS(gps GPS) {
	cb.CarManual.Instruction = fmt.Sprintf("%sThis car has %s aviation system\n", cb.CarManual.Instruction, gps)
}

// concrete builder is expect to have their own method for getting product
// because various builder may create various products that don't all follow same interface(at least in static type)
func (cb *CarManualBuilder) getProduct() CarManual {
	product := cb.CarManual
	cb.reset()
	return product
}

// Director is optional, it's is for abstract the calling building step process
type Director struct {
}

func NewDirector() *Director {
	return &Director{}
}

// Director works with abstract builder, allow flexibility
func (d *Director) constructSportCar(builder Builder) {
	builder.reset()
	builder.setEngine(SportEngine)
	builder.setSeat(2)
}

func (d *Director) constructSUVCar(builder Builder) {
	builder.reset()
	builder.setEngine(SUVEngine)
	builder.setSeat(5)
	builder.setGPS(GoogleGPS)
}

func main() {
	// director for abstract the step
	director := NewDirector()

	// init the builders
	carBuilder := &CarBuilder{}
	carManualBuilder := &CarManualBuilder{}

	// flexibility in creating product
	//make some car
	director.constructSUVCar(carBuilder)
	fmt.Println(carBuilder.getProduct())
	//make some car manual
	director.constructSUVCar(carManualBuilder)
	fmt.Println(carManualBuilder.getProduct())

	// Or we can call the builder directly for customized  product creation
	carBuilder.setEngine(SportEngine)
	carBuilder.setSeat(10)

}
