package main

import (
  "fmt"
  "math"
)

type figura interface{
   area() float64
}

type quadrado struct{
  lado float64
}
type circulo struct{
  raio float64
}

func (q quadrado) area() float64{
  return q.lado * q.lado;
}
func (c circulo) area() float64{
  return 2 * math.Pi * c.raio;
}
func info(f figura){
  fmt.Println("Área da Figura:",f.area());
}
func print(value string){
  fmt.Println(value);
}
func main(){
   bola_de_basquete := circulo{
     raio:12.0,   
   }
   dado := quadrado{
    lado: 26.0,
   }
   print("Bola de Basquete:");
   info(bola_de_basquete);
   print("Dado:");
   info(dado);
}
