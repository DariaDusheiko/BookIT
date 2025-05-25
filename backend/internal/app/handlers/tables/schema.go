package tables

import (
	"github.com/gin-gonic/gin"
	"github.com/BookIT/backend/internal/app/services"
	"errors"
)

type TableSchemas struct{}

func (s *TableSchemas) ValidateGetTablesRequest(c *gin.Context) (*GetTablesRequest, error) {
	var req GetTablesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	if req.End.Before(req.Start) {
		return nil, errors.New("end time must be after start time")
	}

	return &req, nil
}

func (s *TableSchemas) NewGetTablesResponse(tables []services.TableWithOccupancy) *GetTablesResponse {
	items := make([]TableResponse, 0, len(tables))
	for _, table := range tables {
		items = append(items, TableResponse{
			ID:      table.ID,
			X:           table.X,
			Y:           table.Y,
			Angle:       table.Angle,
			SeatsNumber: table.SeatsNumber,
			Occupied:    table.Occupied,
		})
	}
	return &GetTablesResponse{Tables: items}
}

func (s *TableSchemas) NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Error: message}
}

