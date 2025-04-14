package api

import (
	"banking-service/internal/models"
	"banking-service/internal/service"
	"banking-service/pkg/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	accountService *service.AccountService
}

func NewHandler(accountService *service.AccountService) *Handler {
	return &Handler{
		accountService: accountService,
	}
}

func (h *Handler) HandleDaftar(c echo.Context) error {
	var request models.DaftarRequest

	if err := c.Bind(&request); err != nil {
		logger.Error("Invalid request format", "error", err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Remark: "Format request tidak valid",
		})
	}

	response, err := h.accountService.RegisterNasabah(request)
	if err != nil {
		logger.Error("Registration failed", "error", err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) HandleTabung(c echo.Context) error {
	var request models.TabungRequest

	if err := c.Bind(&request); err != nil {
		logger.Error("Invalid request format", "error", err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Remark: "Format request tidak valid",
		})
	}

	response, err := h.accountService.Deposit(request)
	if err != nil {
		logger.Error("Deposit failed", "error", err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) HandleTarik(c echo.Context) error {
	var request models.TarikRequest

	if err := c.Bind(&request); err != nil {
		logger.Error("Invalid request format", "error", err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Remark: "Format request tidak valid",
		})
	}

	response, err := h.accountService.Withdraw(request)
	if err != nil {
		logger.Error("Withdrawal failed", "error", err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) HandleSaldo(c echo.Context) error {
	noRekening := c.Param("no_rekening")

	response, err := h.accountService.GetBalance(noRekening)
	if err != nil {
		logger.Error("Balance inquiry failed", "error", err.Error())
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Remark: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
