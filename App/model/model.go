package model

// Music struct
type Music struct {
	ID        int    `json:"ID"`
	NamaLagu  string `json:"Namalagu"`
	NamaArtis string `json:"NamaArtis"`
	Link      string `json:"Link"`
	Durasi    int    `json:"Durasi"`
}
