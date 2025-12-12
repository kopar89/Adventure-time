package main

import (
	f "fmt"
	d "games/dialog/dialog_for_persons"
	p "games/hero_info/persons"
	a "games/interface/hero"
	s "strings"
	t "time"
)

func main() {

	var P p.Person
	var hero p.Hero
	var gameMap a.Map
	var heroID int

	d.Dialog_1(&P)
	d.Dialog_2(&P)
	f.Scan(&heroID)

	switch heroID {
	case 1:
		hero = p.Hero1
	case 2:
		hero = p.Hero2
	case 3:
		hero = p.Hero3
	case 4:
		hero = p.Hero4
	default:
		hero = p.Hero1
	}
	f.Println("Ты выбрал героя:", hero.Name)
	f.Println("Удачи в твоих приключениях!")


	gameMap.Generate()
	gameMap.PlaceHero()

	enemies := []*p.Enemy{
		&p.Enemy_1,
		&p.Enemy_2,
		&p.Enemy_3,
		&p.Enemy_4,
	}

	gameMap.PlaceEnemies(len(enemies)) 
	gameMap.PlaceMedkits(12)           

	f.Println("\nПокажу тебе карту...")
	t.Sleep(1 * t.Second)
	gameMap.Print()

	d.Dialog_3(&P)


	f.Println("Управление: W/A/S/D (exit для выхода)")
	for {
		var cmd string
		f.Print("Ход: ")
		f.Scan(&cmd)
		cmd = s.ToUpper(cmd)

		switch cmd {
		case "W":
			gameMap.MoveHero(-1, 0, &hero, enemies)
		case "S":
			gameMap.MoveHero(1, 0, &hero, enemies)
		case "A":
			gameMap.MoveHero(0, -1, &hero, enemies)
		case "D":
			gameMap.MoveHero(0, 1, &hero, enemies)
		case "EXIT":
			f.Println("Выход из игры.")
			return
		default:
			f.Println("Неизвестная команда")
		}

		gameMap.Print()

		if hero.Health <= 0 {
			f.Println("Игра окончена! Герой погиб...")
			return
		}
	}
}
