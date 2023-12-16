package main

type Item struct {
	name     string
	size     string
	color    string
	material string
	amount   string
}

func createItem(name string,
	size string,
	color string,
	material string,
	amount string) Item {
	return Item{
		name:     name,
		size:     size,
		color:    color,
		material: material,
		amount:   amount,
	}
}

func initializeItems() []Item {
	var items []Item

	// TODO брать данные из бд и использовать здесь
	//for i := 0; i < 6; i++ {
	//	items = append(items, createItem(itemsMock[0][i], itemsMock[1][i], itemsMock[2][i], itemsMock[3][i], itemsMock[4][i]))
	//}

	return items
}
