package Room

import (
	"context"
	"fmt"
)

type Equipment struct {
	name     string
	color    string
	material string
}

func GetRoomEquipmentFromDB() ([]Equipment, error) {
	conn := getDatabase()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select name, color, material from \"RoomEquipment\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []Equipment
	for rows.Next() {
		var item Equipment
		if err := rows.Scan(&item.name, &item.color, &item.material); err != nil {
			return nil, err
		}
		equipments = append(equipments, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return equipments, nil
}

func PrintRoomEquipment() {
	equipments, err := GetRoomEquipmentFromDB()
	FindError(err)

	PrintTitle("ТЕХНИКА")
	for i := 0; i < len(equipments); i++ {
		fmt.Print("\n\nНазвание: ", equipments[i].name, "\nЦвет: ", equipments[i].color, "\nМатериал: ", equipments[i].material)
	}
}
