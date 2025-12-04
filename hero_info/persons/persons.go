package persons

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Hero struct {
	Name    string
	Damage  int
	Mana    int
	Health  int
	Special string
	Power   int
}

type Enemy struct {
	name   string
	damage int
	health int
}

// Блок героев
var Hero1 = Hero{
	Name:    "Артур",
	Damage:  45,
	Mana:    35,
	Health:  100,
	Special: "Shoter",
	Power:   25,
}

var Hero2 = Hero{
	Name:    "Джозеф",
	Damage:  30,
	Mana:    40,
	Health:  120,
	Special: "Medic",
	Power:   17,
}

var Hero3 = Hero{
	Name:    "Питер",
	Damage:  36,
	Mana:    40,
	Health:  100,
	Special: "Engineer",
	Power:   20,
}

var Hero4 = Hero{
	Name:    "Карлос",
	Damage:  25,
	Mana:    50,
	Health:  150,
	Special: "Defender",
	Power:   15,
}

// Блок злодеев
var Enemy_1 = Enemy{
	name:   "Гоблин",
	damage: 20,
	health: 100,
}

var Enemy_2 = Enemy{
	name:   "Лучница",
	damage: 30,
	health: 100,
}

// Инициализация бойцов
type Warrior1 struct {
	Hero
	Power int
}

func (w Warrior1) Special_hit_hero_1(enemyHealth int) int {
	if w.Mana > 60 {
		fmt.Printf("%s наносит спец. удар с уроном %d!\n", w.Name, w.Power)
	}
	return enemyHealth - w.Power
}

type Warrior2 struct {
	Hero
	Power int
}

func (w Warrior2) Special_hit_hero_2(enemyHealth int) int {
	if w.Mana > 50 {
		fmt.Printf("%s наносит спец. удар с уроном %d!\n", w.Name, w.Power)
	}
	return enemyHealth - w.Power
}

type Warrior3 struct {
	Hero
	Power int
}

func (w Warrior3) Special_hit_hero_3(enemyHealth int) int {
	if w.Mana > 50 {
		fmt.Printf("%s наносит спец. удар с уроном %d!\n", w.Name, w.Power)
	}
	return enemyHealth - w.Power
}

type Warrior4 struct {
	Hero
	Power int
}

func (w Warrior4) Special_hit_hero_4(enemyHealth int) int {
	if w.Mana > 40 {
		fmt.Printf("%s наносит спец. удар с уроном %d!\n", w.Name, w.Power)
	}
	return enemyHealth - w.Power
}

func (h Hero) Info() string {
	return fmt.Sprintf("%s | Урон: %d | Мана: %d | Здоровье: %d", h.Name, h.Damage, h.Mana, h.Health)
}
