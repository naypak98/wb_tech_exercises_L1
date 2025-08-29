package l1_1

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (h Human) Student() string {
	return fmt.Sprintf("Студент \"%s %s\" (возраст %d)", h.Surname, h.Name, h.Age)
}

type Action struct {
	Human
	ActionName string
}

func Run() {
	a := Action{
		Human: Human{
			Name:    "Алиса",
			Surname: "Попова",
			Age:     22,
		},
		ActionName: "зачислен",
	}

	fmt.Println("Задание L1.1: " + "\n" + a.Student() + " " + a.ActionName + "\n")
}
