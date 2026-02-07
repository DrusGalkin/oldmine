package app

import (
	"database/sql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"migrator/internal/storage"
)

const DUPLICATE_TABLE = "42P07"
const INVALID_SYNTAX = "42601"

func Up(db *sql.DB) {
	for _, m := range storage.AllMigrations {
		if m == nil {
			continue
		}

		_, err := db.Exec(m.UpSQL)
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok {
				switch pgErr.Code {
				case DUPLICATE_TABLE:
					log.Println("Таблица", m.Version, "уже существует")
				case INVALID_SYNTAX:
					log.Println("Невалидный синтаксис в коде миграции", m.Version+", проверьте правильность кода")
				}
			}
			log.Println("Неизвестная ошибка при миграции таблицы", m.Version)
			continue
		}
		log.Println("Таблица", m.Version, "создана")
	}
}
