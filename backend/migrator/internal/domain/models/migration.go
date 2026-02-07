package models

type Migration struct {
	Version string
	UpSQL   string
	DownSQL string
}
