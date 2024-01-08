package Room

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
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
	urlExample := "postgres://userKatya:superSecret@localhost:5436/katya"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select number, square, length, width, height from \"RoomParameters\"")
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting people from the database: %v\n", err)
		os.Exit(1)
	}

	PrintTitle("О НАШЕМ ДОМЕ")
	fmt.Print("\nМы живем в общежитии, в комнате №", room.number, "\nЕе площадь ", room.square, " м")
	fmt.Print("\nЕе размеры:", room.height, "×", room.length, "×", room.width, " м")
}
