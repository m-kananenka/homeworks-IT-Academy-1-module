package box

import (
	"fmt"
	"time"
)

type Collection struct {
	Figures []Figure
	History []string
}

func (c Collection) Quantity() {
	var sum int
	for i := range c.Figures {
		sum += i
	}
	fmt.Println(sum)

	time := time.Now()
	c.History = append(c.History, fmt.Sprintf("Метод количество фигур был добавлен %v\n", time))
}

func (c *Collection) AddFigures(figure Figure) {
	c.Figures = append(c.Figures, figure)

	time := time.Now()
	c.History = append(c.History, fmt.Sprintf("Метод добавления фигур был добавлен %v\n", time))

}

func (c *Collection) SumSqr() float64 {
	var sum float64
	for _, v := range c.Figures {
		sum += v.Sqr()
	}
	time := time.Now()
	c.History = append(c.History, fmt.Sprintf("Метод сумма площадей был добавлен %v\n", time))

	return sum
}

func (c *Collection) SumPer() float64 {
	var sum float64
	for _, v := range c.Figures {
		sum += v.Per()
	}
	time := time.Now()
	c.History = append(c.History, fmt.Sprintf("Метод сумма периметров был добавлен %v\n", time))
	return sum
}

func (c *Collection) CallHistory() {
	for _, i := range c.History {
		fmt.Println(i)
	}
}

func NewCollection(figures []Figure, history []string) Collection {
	return Collection{
		Figures: figures,
		History: history,
	}
}
