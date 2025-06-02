package tables

import "time"

type GetTablesRequest struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type TableResponse struct {
	ID          uint `json:"number"`
	X           int  `json:"x"`
	Y           int  `json:"y"`
	Angle       int  `json:"angle"`
	SeatsNumber int  `json:"seats_number"`
	Occupied    bool `json:"occupied"`
}

type GetTablesResponse struct {
	Tables []TableResponse `json:"tables"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
