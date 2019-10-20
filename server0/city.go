package main

import "fmt"
import (
	"html/template"
	"net/http"
	"strconv"
)

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
var users map[string]*city
func main(){
	users = make(map[string]*city)
	ci:=NewCity("Lund")
	ci.output=25
	co:=construction{name: "Library", cost: 300}
	ci.constructing=co
//	ci:=city{name: "Lund", output: 25, constructing: co}
	
	fmt.Println(ci)
	fmt.Println(ci.finishedIn())
	ci.turn()
	fmt.Println(ci.finishedIn())


	tmpl := template.Must(template.ParseFiles("createCity.html"))


	http.HandleFunc("/createCity", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		keys, ok := r.URL.Query()["user"]
		user:=""

		nouser:=false
		c := NewCity(r.FormValue("name"))
		if !ok || len(keys[0]) < 1 {
			fmt.Println("Url Param 'user' is missing")
			nouser=true
		} else {

			v,_:=strconv.ParseFloat(r.FormValue("output"),64)
			c.output=v
			user=keys[0]
			users[user] = c
		}

		replacements := struct{ 
				Success bool
				NoUser bool 
				User string
				City string}{Success: true, NoUser: nouser, User: user, City: c.name}
		//tmpl.Execute(w, struct{ Success bool}{true})
		tmpl.Execute(w, replacements)
	})

	http.ListenAndServe(":8080", nil)
}
