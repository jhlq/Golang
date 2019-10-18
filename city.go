package main

import "fmt"

type construction struct {
	name string
	cost float64
}
type city struct {
	name string
	output  float64
	constructing construction
	progress float64
}
func NewCity(name string) *city {
	ci:=city{name: name}
	ci.output=0
	co:=construction{"nothing",0}
	ci.constructing=co
	ci.progress=0
	return &ci
}
func (ci *city) turn() {
    ci.progress+=ci.output
}
func (ci *city) finishedIn() float64 {
	if ci.output>0 {
		return (ci.constructing.cost-ci.progress)/ci.output
	} else {
		return -1
	}
}
func main(){
	ci:=NewCity("Lund")
	ci.output=25
	co:=construction{name: "Library", cost: 300}
	ci.constructing=co
//	ci:=city{name: "Lund", output: 25, constructing: co}
	
	fmt.Println(ci)
	fmt.Println(ci.finishedIn())
	ci.turn()
	fmt.Println(ci.finishedIn())
}
