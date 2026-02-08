package storage

import (
	"log"
	"migrator/internal/domain/models"
	"strings"
)

type SQL = string

var AllMigrations = []*models.Migration{
	setMigration(
		"users",
		`create table users (
				id serial primary key,
				name text,
				email text,
				password text,
				created_at date 
				);`,
		`drop table users;`,
	),

	setMigration(
		"admins",
		`create table admins (
				id serial primary key,
				user_id integer,
				foreign key (user_id) references users(id) 
				);`,
		`drop table admins;`,
	),

	setMigration(
		"paid_for",
		`create table paid_for (
				id serial primary key,
				user_id integer,
				foreign key (user_id) references users(id) 
				);`,
		`drop table paid_for;`,
	),

	setMigration(
		"skins",
		`create table skins (
				id serial primary key, 
				user_id integer, 
				path text,
				foreign key (user_id) references users(id)
				);`,
		`drop table skins;`,
	),

	setMigration(
		"cloaks",
		`create table cloaks (
				id serial primary key, 
				user_id integer, 
				path text,
				foreign key (user_id) references users(id)
				);`,
		`drop table cloaks;`,
	),
}

func setMigration(version string, codes ...SQL) *models.Migration {
	var correctedCodes []SQL

	if len(codes) >= 2 {
		if !strings.Contains(strings.ToLower(codes[1]), "drop") {
			for i, c := range codes {
				if strings.Contains(strings.ToLower(c), "drop") {
					tmp := codes[0:i]
					tmp2 := codes[i : len(codes)+1]
					for j := 0; j < (len(tmp)+len(tmp2))-2; j++ {
						if i < len(tmp)-1 {
							correctedCodes = append(correctedCodes, tmp[j])
						} else {
							correctedCodes = append(correctedCodes, tmp2[j])
						}
					}
				}
			}
		}
	} else {
		log.Println("Файл", version, "не содержит кода")
		return nil
	}

	return &models.Migration{
		Version: version,
		UpSQL:   codes[0],
		DownSQL: codes[1],
	}
}
