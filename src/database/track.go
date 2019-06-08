package database

import (
	"time"
)

type Track struct {
	ID             uint64    `db:"id" json:"id"`
	DateCreated    time.Time `db:"date_created" json:"date_created"`
	DateModified   time.Time `db:"date_modified" json:"date_modified,omitempty"`
	Jacket         string    `db:"jacket" json:"jacket"`
	Genre          string    `db:"genre" json:"genre"`
	Credit         string    `db:"credit" json:"credit,omitempty"`
	DisplayBPM     uint16    `db:"displaybpm" json:"displaybpm"`
	Length         uint64    `db:"length" json:"length"`
	Title          string    `db:"title" json:"title"`
	TitleRomani    string    `db:"title_romani" json:"title_romani,omitempty"`
	Artists        string    `db:"artists" json:"artists"`
	ArtistsRomani  string    `db:"artists_romani" json:"artists_romani,omitempty"`
	Subtitle       string    `db:"subtitle" json:"subtitle"`
	SubtitleRomani string    `db:"subtitle_romani" json:"subtitle_romani,omitempty"`
}
