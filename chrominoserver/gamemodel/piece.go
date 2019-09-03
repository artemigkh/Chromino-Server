package gamemodel

type ChrominoPiece struct {
	Colors   [3]ChrominoColor `json:"colors"`
	Rotation int              `json:"rotation"`
	X        int              `json:"x"`
	Y        int              `json:"y"`
}
