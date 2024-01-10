package Room

import (
	"context"
	"fmt"
	"os"
)

type Room struct {
	number int
	square float32
	length float32
	width  float32
	height float32
}

func GetRoomParametersFromDB() (Room, error) {
	conn := getDatabase()
	defer conn.Close(context.Background())

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select number, square, length, width, height from \"RoomParameters\"")
	FindError(err)
	defer rows.Close()

	var myRoom Room
	for rows.Next() {
		var room Room
		if err := rows.Scan(&room.number, &room.square, &room.length, &room.width, &room.height); err != nil {
			fmt.Fprintf(os.Stderr, "Error! %v\n", err)
			os.Exit(1)
		}
		myRoom = room
	}

	return myRoom, nil
}

func PrintRoomParameters() {
	room, err := GetRoomParametersFromDB()
	FindError(err)

	PrintTitle("О НАШЕМ ДОМЕ")
	fmt.Print("\nМы живем в общежитии, в комнате №", room.number, "\nЕе площадь ", room.square, " м")
	fmt.Print("\nЕе размеры:", room.height, "×", room.length, "×", room.width, " м")
}
