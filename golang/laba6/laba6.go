package laba6

import "fmt"

type phone struct {
	model    string
	batery   int
	f_memory int
	emkost   int
	volume   int
}

func (c phone) get_number(r_number string) string {
	switch r_number {
	case "+79345462739":
		return ("Степан Андреевич")
	case "+79145462738":
		return ("Татьяна Николаевна")
	case "+79345462738":
		return ("Юлия Яблокова")
	case "+79355462738":
		return ("Иван Иванович")
	case "+79340462738":
		return ("Дмитрий Степанович")
	case "+79345468738":
		return ("Степан Дмитриевич")
	default:
		return ("Введите номер")
	}
}

func (c phone) time_batery(power int) float64 {
	ost := (100 - c.batery) * c.emkost / 100
	skorost := (power / 5) * 1000
	if skorost == 0 {
		return 0
	}
	tm_bat := float64(ost) / float64(skorost) * 60
	return tm_bat
}

func (c phone) new_volume(nazhatv int, nazhatn int) int {
	kvol := c.volume + nazhatv - nazhatn
	return kvol
}

func Readyylabaa6() {
	Phone := phone{"Xiaomi13", 89, 189, 5000, 2}
	fmt.Println("Вам звонит", Phone.get_number("+79345468738"))
	fmt.Println("у вас", Phone.batery, "%")
	fmt.Printf("Чтобы зарядить батарею вам нужно %.2f минут.\n", Phone.time_batery(5))
	fmt.Println("Звук установлен с", Phone.volume, "на", Phone.new_volume(2, 1))
}
