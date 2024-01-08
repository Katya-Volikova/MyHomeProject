package Room

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func FindError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error! %v\n", err)
		os.Exit(1)
	}
}

func getDatabase() *pgx.Conn {
	urlExample := "postgres://userKatya:superSecret@localhost:5436/katya"
	conn, err := pgx.Connect(context.Background(), urlExample)
	FindError(err)

	return conn
}

func PrintTitle(title string) {
	fmt.Print("\n\n//////////  ", title, "  //////////")
}
