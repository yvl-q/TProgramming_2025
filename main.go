package main

import "fmt"

// Структура Phone
type Phone struct {
	brand        string
	model        string
	calledNumber string
}

// Конструктор
func NewPhone(brand string, model string, calledNumber string) Phone {
	return Phone{
		brand:        brand,
		model:        model,
		calledNumber: calledNumber,
	}
}

// Метод 1: получить номер
func (p Phone) GetCalledNumber() string {
	return p.calledNumber
}

// Метод 2: установить номер
func (p *Phone) SetCalledNumber(number string) {
	p.calledNumber = number
}

// Метод 3: вывести информацию
func (p Phone) PrintInfo() {
	fmt.Println("Бренд:", p.brand)
	fmt.Println("Модель:", p.model)
	fmt.Println("Номер вызываемого абонента:", p.calledNumber)
}

func main() {
	phone := NewPhone("Iphone", "17", "+7987654321")

	fmt.Println("Информация о телефоне:")
	phone.PrintInfo()

	fmt.Println()

	fmt.Println("Текущий номер:", phone.GetCalledNumber())

	phone.SetCalledNumber("+7123456789")

	fmt.Println("Новый номер:", phone.GetCalledNumber())
}
