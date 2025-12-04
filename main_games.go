package main

import (
	f "fmt"
	d "games/dialog/dialog_for_persons"
	p "games/hero_info/persons"
	a "games/interface/hero"
	t "time"
)

func main() {

	//Блок переменных
	var P p.Person
	//var H p.Hero
	var A a.Hero_Map
	d.Dialog_1(&P)
	f.Println("Покажу тебе нашу карту!")

	A.Generate()

	t.Sleep(2 * t.Second)
	d.Dialog_2(&P)

}
