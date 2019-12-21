package database

import (
	"time"
)

type Track struct {
	ID             uint64    `db:"id" json:"id,omitempty"`
	DateCreated    time.Time `db:"date_created" json:"date_created,omitempty"`
	DateModified   time.Time `db:"date_modified" json:"date_modified,omitempty"`
	Publisher      string    `db:"publisher" json:"publisher,omitempty"`
	Jacket         string    `db:"jacket" json:"jacket,omitempty"`
	Genre          string    `db:"genre" json:"genre,omitempty"`
	Credit         string    `db:"credit" json:"credit,omitempty"`
	DisplayBPM     uint16    `db:"displaybpm" json:"displaybpm,omitempty"`
	Length         uint64    `db:"length" json:"length,omitempty"`
	Title          string    `db:"title" json:"title,omitempty"`
	TitleRomani    string    `db:"title_romani" json:"title_romani,omitempty"`
	Artists        string    `db:"artists" json:"artists,omitempty"`
	ArtistsRomani  string    `db:"artists_romani,omitempty" json:"artists_romani,omitempty"`
	Subtitle       string    `db:"subtitle" json:"subtitle,omitempty"`
	SubtitleRomani string    `db:"subtitle_romani,omitempty" json:"subtitle_romani,omitempty"`
}
