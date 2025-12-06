package hero

import (
	"fmt"
	p "games/hero_info/persons"
	"math/rand"
	"github.com/fatih/color"
)

type Map struct {
	rows, cols int
	grid       [][]rune
	heroPos    struct{ X, Y int }
	enemies    []struct{ X, Y int }
}

func (m *Map) Generate() {
	m.rows, m.cols = 30, 30
	m.grid = make([][]rune, m.rows)

	grass := '░'
	forest := '▓'
	mountain := '▲'
	water := '≈'

	for i := range m.grid {
		m.grid[i] = make([]rune, m.cols)
		for j := range m.grid[i] {
			r := rand.Float64()
			switch {
			case r < 0.5:
				m.grid[i][j] = grass
			case r < 0.7:
				m.grid[i][j] = forest
			case r < 0.85:
				m.grid[i][j] = mountain
			default:
				m.grid[i][j] = water
			}
		}
	}
}

func (m *Map) PlaceHero() {
	for {
		x := rand.Intn(m.rows)
		y := rand.Intn(m.cols)
		if m.grid[x][y] != '≈' && m.grid[x][y] != '▲' {
			m.heroPos.X, m.heroPos.Y = x, y
			return
		}
	}
}

func (m *Map) PlaceEnemies(count int) {
	m.enemies = []struct{ X, Y int }{}
	for k := 0; k < count; k++ {
		for {
			x := rand.Intn(m.rows)
			y := rand.Intn(m.cols)
			if m.grid[x][y] != '≈' && m.grid[x][y] != '▲' && (x != m.heroPos.X || y != m.heroPos.Y) {
				m.enemies = append(m.enemies, struct{ X, Y int }{x, y})
				break
			}
		}
	}
}

func (m *Map) Print() {
	green := color.New(color.FgGreen).SprintFunc()
	darkGreen := color.New(color.FgHiGreen).SprintFunc()
	gray := color.New(color.FgHiBlack).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {

			if i == m.heroPos.X && j == m.heroPos.Y {
				fmt.Print(yellow("H"))
				continue
			}

			enemyHere := false
			for _, e := range m.enemies {
				if e.X == i && e.Y == j {
					fmt.Print(red("E"))
					enemyHere = true
					break
				}
			}
			if enemyHere {
				continue
			}

			switch m.grid[i][j] {
			case '░':
				fmt.Print(green("░"))
			case '▓':
				fmt.Print(darkGreen("▓"))
			case '▲':
				fmt.Print(gray("▲"))
			case '≈':
				fmt.Print(blue("≈"))
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (m *Map) MoveHero(dx, dy int, hero *p.Hero, enemies []*p.Enemy) {
	nx := m.heroPos.X + dx
	ny := m.heroPos.Y + dy

	if nx < 0 || ny < 0 || nx >= m.rows || ny >= m.cols {
		return
	}

	if m.grid[nx][ny] == '≈' || m.grid[nx][ny] == '▲' {
		return
	}

	for i, e := range m.enemies {
		if e.X == nx && e.Y == ny {
			if i >= 0 && i < len(enemies) && enemies[i] != nil {
				Fight(hero, enemies[i])

				if enemies[i].Health <= 0 {
					m.enemies = append(m.enemies[:i], m.enemies[i+1:]...)
					copy(enemies[i:], enemies[i+1:])
					enemies[len(enemies)-1] = nil
					enemies = enemies[:len(enemies)-1]
				}

			} else {
				fmt.Println("Ошибка: неверный индекс врага при столкновении.")
			}
			return
		}
	}

	m.heroPos.X, m.heroPos.Y = nx, ny
}

func Fight(hero *p.Hero, enemy *p.Enemy) {
	fmt.Println("⚔️ Бой начался:", hero.Name, "против", enemy.Name)

	for hero.Health > 0 && enemy.Health > 0 {

		enemy.Health -= hero.Damage
		if enemy.Health < 0 {
			enemy.Health = 0
		}

		fmt.Printf("%s наносит удар (%d). У %s осталось %d HP\n",
			hero.Name, hero.Damage, enemy.Name, enemy.Health)

		if enemy.Health <= 0 {
			fmt.Println("✨ Враг повержен!")
			fmt.Print(hero.Info())
			fmt.Print("\n")
			return
		}

		hero.Health -= enemy.Damage
		if hero.Health < 0 {
			hero.Health = 0
		}

		fmt.Printf("%s атакует (%d). У героя %s осталось %d HP\n",
			enemy.Name, enemy.Damage, hero.Name, hero.Health)

		if hero.Health <= 0 {
			fmt.Println("💀 Герой погиб...")
			hero.Info()
			return
		}
	}
}
