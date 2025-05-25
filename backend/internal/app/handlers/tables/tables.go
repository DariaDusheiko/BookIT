package tables

import (
	"github.com/BookIT/backend/internal/app/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type TableHandler struct {
	tableService services.TableService 
	schemas      *TableSchemas
}

func NewTableHandler(tableService services.TableService) *TableHandler { 
	return &TableHandler{
		tableService: tableService,
		schemas:      &TableSchemas{},
	}
}

// GetTables godoc
// @Summary Get tables with occupancy status
// @Description Returns list of tables with occupancy status for given time range
// @Tags tables
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param X-Auth-Token header string true "JWT Token"
// @Param request body GetTablesRequest true "Time range"
// @Success 200 {object} GetTablesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tables [post]
func (h *TableHandler) GetTables(c *gin.Context) {
	_, exists := c.Get("userID") 
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
		return
	}

	var req GetTablesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	startPtr := &req.Start
	endPtr := &req.End

	tables, err := h.tableService.GetTablesWithOccupancy(startPtr, endPtr)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := make([]TableResponse, 0, len(tables))
	for _, table := range tables {
		response = append(response, TableResponse{
			ID:          table.ID,
			X:           table.X,
			Y:           table.Y,
			Angle:       table.Angle,
			SeatsNumber: table.SeatsNumber,
			Occupied:    table.Occupied,
		})
	}

	c.JSON(200, gin.H{"tables": response})
}

