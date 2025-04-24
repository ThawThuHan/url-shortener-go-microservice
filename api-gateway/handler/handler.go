package handler

import (
	redirection "api-gateway/proto/redirection-service"
	shortener "api-gateway/proto/shortener-service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	shortener.ShortenerServiceClient
	redirection.RedirectionServiceClient
}

func NewHandler(shortenerClient shortener.ShortenerServiceClient, redirectionClient redirection.RedirectionServiceClient) *Handler {
	return &Handler{ShortenerServiceClient: shortenerClient, RedirectionServiceClient: redirectionClient}
}

func (h *Handler) CreateShortURL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			OriginUrl string `json:"origin_url"`
			SessionId string `json:"session_id"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		grpcReq := &shortener.CreateShortURLRequest{OriginUrl: req.OriginUrl, SessionId: req.SessionId}
		grpcRes, err := h.ShortenerServiceClient.CreateShortURL(c.Context(), grpcReq)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(grpcRes)
	}
}

func (h *Handler) GetOriginURL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortCode := c.Params("short_code")
		if shortCode == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "short_code is required",
			})
		}

		grpcReq := &redirection.GetOriginURLRequest{ShortCode: shortCode}
		grpcRes, err := h.RedirectionServiceClient.GetOriginURL(c.Context(), grpcReq)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(grpcRes)
	}
}

func (h *Handler) GetAccessLog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortCode := c.Params("short_code")
		if shortCode == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "short_code is required",
			})
		}

		grpcReq := &redirection.GetAccessLogRequest{ShortCode: shortCode}
		grpcRes, err := h.RedirectionServiceClient.GetAccessLog(c.Context(), grpcReq)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(grpcRes)
	}
}

func (h *Handler) RedirectHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			ShortCode string `json:"short_code"`
			IPAddress string `json:"ip_address"`
			Location  string `json:"location"`
			City      string `json:"city"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		grpcReq := &redirection.GetOriginURLRequest{ShortCode: req.ShortCode, IpAddress: req.IPAddress, Location: req.Location, City: req.City}
		grpcRes, err := h.RedirectionServiceClient.GetOriginURL(c.Context(), grpcReq)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"origin_url": grpcRes.OriginUrl,
		})
	}
}

func (h *Handler) GetUrls() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionId := c.Params("session_id")
		if sessionId == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "session is required",
			})
		}
		grpcReq := &shortener.GetOriginURLsRequest{SessionId: sessionId}
		grpcRes, err := h.ShortenerServiceClient.GetOriginURLs(c.Context(), grpcReq)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(grpcRes)
	}
}
