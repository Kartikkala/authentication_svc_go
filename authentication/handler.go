package authentication

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *SvcWithHooks
}

func NewHandler(svc *SvcWithHooks) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RegisterHandler(c echo.Context) error {
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	err := h.svc.RegisterUserEntry(c.Request().Context(), req.Email, req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "user registered")
}

func (h *Handler) LoginHandler(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	user, err := h.svc.LoginUserEntry(c.Request().Context(), req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	token, err := h.svc.GenerateTokenEntry(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"token": token,
	})
}

func (h *Handler) TokenVerificationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tok := c.Request().Header.Get("token")
		if tok == "" {
			return echo.NewHTTPError(401, "Missing token header")
		}
		claims, err := h.svc.DecodeTokenEntry(&tok)

		if err != nil {
			fmt.Println("Token verification error: ", err.Error())
			return echo.NewHTTPError(401, "Invalid or expired token noob!")
		}
		c.Set("user", claims)

		return next(c)
	}
}
