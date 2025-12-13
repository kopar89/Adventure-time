package persons

import "fmt"

type HeroInterface interface {
	GetInfo() string
	GetDamage() int
	GetHealth() int
	SetHealth(int)
}

type EnemyInterface interface {
	GetInfo() string
	GetDamage() int
	GetHealth() int
	SetHealth(int)
}

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

func (h *Hero) GetInfo() string {
	return fmt.Sprintf("%s | HP: %d | Damage: %d | Mana: %d | Special: %s | Power: %d",
		h.Name, h.Health, h.Damage, h.Mana, h.Special, h.Power)
}

func (h *Hero) GetDamage() int {
	return h.Damage
}

func (h *Hero) GetHealth() int {
	return h.Health
}

func (h *Hero) SetHealth(hp int) {
	h.Health = hp
}

type Enemy struct {
	Name   string
	Damage int
	Health int
}

func (e *Enemy) GetInfo() string {
	return fmt.Sprintf("%s | HP: %d | Damage: %d", e.Name, e.Health, e.Damage)
}

func (e *Enemy) GetDamage() int {
	return e.Damage
}

func (e *Enemy) GetHealth() int {
	return e.Health
}

func (e *Enemy) SetHealth(hp int) {
	e.Health = hp
}

type CharacterFactory interface {
	CreateHero(heroType int) HeroInterface
	CreateEnemy(enemyType int) EnemyInterface
}

type Factory1 struct{}

func (f *Factory1) CreateHero(heroType int) HeroInterface {
	switch heroType {
	case 1:
		return &Hero{Name: "Артур", Damage: 45, Mana: 35, Health: 100, Special: "Shooter", Power: 25}
	case 2:
		return &Hero{Name: "Джозеф", Damage: 30, Mana: 40, Health: 120, Special: "Medic", Power: 17}
	case 3:
		return &Hero{Name: "Питер", Damage: 36, Mana: 40, Health: 100, Special: "Engineer", Power: 20}
	case 4:
		return &Hero{Name: "Карлос", Damage: 25, Mana: 50, Health: 150, Special: "Defender", Power: 15}
	default:
		return &Hero{Name: "Артур", Damage: 45, Mana: 35, Health: 100, Special: "Shooter", Power: 25}
	}
}

func (f *Factory1) CreateEnemy(enemyType int) EnemyInterface {
	switch enemyType {
	case 1:
		return &Enemy{Name: "Гоблин", Damage: 20, Health: 100}
	case 2:
		return &Enemy{Name: "Лучница", Damage: 30, Health: 100}
	case 3:
		return &Enemy{Name: "Орк", Damage: 25, Health: 120}
	case 4:
		return &Enemy{Name: "Волшебник", Damage: 35, Health: 80}
	default:
		return &Enemy{Name: "Гоблин", Damage: 20, Health: 100}
	}
}
