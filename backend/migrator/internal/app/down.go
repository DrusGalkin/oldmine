package app

import (
	"database/sql"
	"log"
	"migrator/internal/storage"
)

func Down(db *sql.DB) {
	for _, m := range storage.AllMigrations {
		if m == nil {
			continue
		}

		_, err := db.Exec(m.DownSQL)
		if err != nil {
			log.Println("Ошибка дропа", m.Version)
			continue
		}
		log.Println(m.Version, " - удалена")
	}
}
