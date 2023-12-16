package main

import "fmt"

type Room struct {
	number int
	square float32
	length float32
	width  float32
	height float32
}

func (room Room) infoAboutRoom(items []Item, equipment []Equipment) {
	fmt.Print("\nМы живем в общежитии в комнате №", room.number, "\nЕе площадь ", room.square)
	fmt.Print("\nЕе размеры:", room.height, "×", room.length, "×", room.width, " м")
	fmt.Print("\n\nВ ней есть такие предметы:\n\n")

	for i := 0; i < 6; i++ {
		fmt.Print(items[i].name, "\nРазмеры: ", items[i].size, " м", "\nЦвет: ", items[i].color, "\nМатериал: ", items[i].material, "\nКоличество: ", items[i].amount, " шт", "\n\n")
	}

	fmt.Print("\n\nВ ней есть такая техника:\n\n")
	for i := 0; i < 5; i++ {
		fmt.Print(equipment[i].name, "\nЦвет: ", equipment[i].color, "\nМатериал: ", equipment[i].material, "\n\n")
	}

}
