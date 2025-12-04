package hero

import (
	f "fmt"
	t "time"
)

type Special_hit_hero_1 interface {
	hit_1(int) int
}

type Special_hit_hero_2 interface {
	hit_2(int) int
}

type Special_hit_hero_3 interface {
	hit_3(int) int
}

type Special_hit_hero_4 interface {
	hit_4(int) int
}

type Hero_Map struct {
	rows   int
	cols   int
	matrix [][]int
}

type GenerateMap interface {
	MapMatrix(rows, cols int) [][]int
}

func (m *Hero_Map) Generate() {
	//var  []string{"", "", ""}
	m.rows = 30
	m.cols = 30
	m.matrix = make([][]int, m.rows)
	for i := range m.matrix {
		m.matrix[i] = make([]int, m.cols)
		for j := range m.matrix[i] {
			m.matrix[i][j] = 0
		}
	}

	for i := 0; i <= 100; i += 10 {
		f.Printf("\rЗагрузка карты: %d%%", i)
		t.Sleep(200 * t.Millisecond)
	}
	f.Println("\nКарта готова!")

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			f.Print(m.matrix[i][j], " ")
		}
		f.Println()
	}

}
