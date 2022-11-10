package gobyexample
//https://gobyexample.com/interfaces
import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (retangulo rect) area() float64 {
    return retangulo.width * retangulo.height;
}
func (retangulo rect) perim() float64 {
    return 2*retangulo.width + 2*retangulo.height;
}

func (circulo circle) area() float64 {
    return math.Pi * circulo.radius * circulo.radius;
}
func (circulo circle) perim() float64 {
    return 2 * math.Pi * circulo.radius;
}

func measure(g geometry) {
    fmt.Println("Geometria:",g);
    fmt.Println("Área:",g.area());
    fmt.Println("Perímetro:",g.perim());
}

func Medir() {
    retangulo := rect{width: 3, height: 4};
    circulo := circle{radius: 5};

	fmt.Println("https://gobyexample.com/interfaces")
	fmt.Println("[Retângulo]:");
    measure(retangulo);
	fmt.Println("[Círculo]:");
    measure(circulo);
}
