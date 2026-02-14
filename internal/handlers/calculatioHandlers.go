package handlers

import (
	"http_server/internal/calculationService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandler struct {
	service calculationService.CalculationService //calculationService - package name
	//CalculationServic - INTERFACE name // SO packageName.TypeName
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

////

// <<GET>>//////////////////////////////////////////////////////////////////////////////////////////////////////
func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot get calculations"})

	}
	return c.JSON(http.StatusOK, calculations)
}

// //<<PATCH>>/////////////////////////////////////////////////////////////////////////////////////////////////////////
func (h *CalculationHandler) PatchCalculation(c echo.Context) error {
	id := c.Param("id")
	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedCalc, err := h.service.UpdateCalculation(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot update calculation"})
	}
	return c.JSON(http.StatusOK, updatedCalc)
}

// <<POST>>///////////////////////////////////////////////////////////////////////////////////////////////////////////
func (h *CalculationHandler) PostCalculations(c echo.Context) error {
	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	calc, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot create calculation"})
	}
	return c.JSON(http.StatusCreated, calc)

}

// <<Delete>>///////////////////////////////////////////////////////////////////////////////////////////////////////////
func (h *CalculationHandler) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot delete calculation"})
	}
	
	return c.NoContent(http.StatusNoContent)
}
