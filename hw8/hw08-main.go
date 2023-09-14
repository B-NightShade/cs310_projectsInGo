/* DO NOT MODIFY THIS FILE */
package main

import (
  "log"
  "os"
  "hw08"
  "strconv"
  "math/rand"
)

func main(){
  //check cmd line options
  if len(os.Args) != 3 {
    log.Fatal("usage: ./hw08 N seed")
  }

  //parse cmd line options, seed randomness
  N,_:= strconv.Atoi(os.Args[1])
  seed,_ := strconv.ParseInt(os.Args[2],10,64)
  rand.Seed(seed)

  //generate N random shapes 
  shapes := make([]hw08.Shape, N)
  var s hw08.Shape
  for i := range shapes {
    r := rand.Intn(4)
    switch r {
      case 0: s = genRectangle()
      case 1: s = genCircle()
      case 2: s = genSphere()
      case 3: s = genBox()
    }
    shapes[i] = s
  }

  //print details of each shape
  for _,s := range shapes {
    s.Print()
  }
}

//generate slice of N random decimal numbers 
func decimals(N int) []float64{
  d := make([]float64, N)
  for i := range d {
    d[i] = rand.Float64() * (10.0 - 0.0)
  }
  return d
}

//wrapper function to create Rectangle
func genRectangle() hw08.Rectangle{
  nums := decimals(2)
  return hw08.MakeRectangle(nums[0],nums[1])
}

//wrapper function to create Circle
func genCircle() hw08.Circle{
  nums := decimals(1)
  return hw08.MakeCircle(nums[0])
}

//wrapper function to create Sphere
func genSphere() hw08.Sphere{
  nums := decimals(1)
  return hw08.MakeSphere(nums[0])
}

//wrapper function to create Box
func genBox() hw08.Box{
  nums := decimals(3)
  return hw08.MakeBox(nums[0], nums[1], nums[2])
}
