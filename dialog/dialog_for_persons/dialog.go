package dialog_for_person

import (
	f "fmt"
	p "games/hero_info/persons"
	os "os"
	t "time"

	"github.com/fatih/color"
)

func Dialog_1(P *p.Person) {
	f.Println("Игрок, Добро пожаловать в игру Adventure time! Здесь тебе придется путешетсвовать, бороться, выживать...")
	t.Sleep(5 * t.Second)
	f.Println("В общем жить еще одну жизнь, но только в виртуальном мире!")
	t.Sleep(4 * t.Second)
	f.Print("Прежде чем начать, давай познакомимся! Как тебя зовут?\t")
	f.Scan(&P.Name)
	f.Printf("Прекрасно %s! Сколько тебе лет? ", P.Name)
	f.Scan(&P.Age)
	if P.Age < 10 {
		f.Printf("%s, изивини, но ты слишком мал, для данной игры\n", P.Name)
		os.Exit(0)

	} else {
		f.Printf("Приятно было познакомиться, %s, аххх, прости...\nЗабыл представиться, совсем вылетело из головы...\nМеня зовут Люцифер, я твой помощник в этом диком мире!\n", P.Name)
	}
}

func Dialog_2(P *p.Person) {
	f.Print("Но для начала тебе нужно выбрать героя, за которого ты будешь играть!\n")
	t.Sleep(3 * t.Second)
	f.Print("У нас есть 4 героя на выбор:\n")
	t.Sleep(2 * t.Second)
	f.Println(p.Hero1.Info())
	t.Sleep(2 * t.Second)
	f.Println(p.Hero2.Info())
	t.Sleep(2 * t.Second)
	f.Println(p.Hero3.Info())
	t.Sleep(2 * t.Second)
	f.Println(p.Hero4.Info())
	t.Sleep(2 * t.Second)
	f.Print("Кого из них ты выбираешь? (введи цифру от 1 до 4): ")
}

func Dialog_3(P *p.Person) {

	green := color.New(color.FgGreen).SprintFunc()
	darkGreen := color.New(color.FgHiGreen).SprintFunc()
	gray := color.New(color.FgHiBlack).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	f.Print("Это трава, по ней можешь спокойно ходить, не переживай ", green("░"), darkGreen("▓"), "\n")
	t.Sleep(2 * t.Second)

	f.Print("А это горы, по ним ходить нельзя, они слишком крутые ", gray("▲"), "\n")
	t.Sleep(2 * t.Second)

	f.Print("А это вода, по ней тоже ходить нельзя, утонешь ", blue("≈"), "\n")
	t.Sleep(2 * t.Second)

	f.Print("Это аптечка, она тебе обязательно понадобиться ", green("A"), "\n")
	t.Sleep(2 * t.Second)

	f.Print("Вот так выглядят твои враги ", red("E"), ", будь осторожен с ними!\n")
	t.Sleep(2 * t.Second)

	f.Print("А это ты ", yellow("H"), ", твоя задача выжить в этом мире и победить всех врагов!\n")
	t.Sleep(2 * t.Second)

	f.Print("Удачи в твоих приключениях!\n")
	t.Sleep(2 * t.Second)
}
