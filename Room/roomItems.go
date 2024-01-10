package Room

import (
	"context"
	"fmt"
)

type Item struct {
	name     string
	size     string
	color    string
	material string
}

func GetRoomItemsFromDB() ([]Item, error) {
	conn := getDatabase()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select name, size, color, material from \"RoomItems\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itemsMass []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.name, &item.size, &item.color, &item.material); err != nil {
			return nil, err
		}
		itemsMass = append(itemsMass, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return itemsMass, nil
}

func PrintRoomItems() {
	items, err := GetRoomItemsFromDB()
	FindError(err)

	PrintTitle("ПРЕДМЕТЫ")
	for i := 0; i < len(items); i++ {
		fmt.Print("\n\nНазвание: ", items[i].name, "\nРазмер: ", items[i].size, " м", "\nЦвет: ", items[i].color, "\nМатериал: ", items[i].material)
	}
}
