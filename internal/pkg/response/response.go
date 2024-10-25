package response

import (
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/types"
	"github.com/gin-gonic/gin"
)

// Environment type for application environment
type Environment string

const (
	Production  Environment = "production"
	Development Environment = "development"
	Test        Environment = "test"
)

// RequestInfo holds information about the HTTP request
type RequestInfo struct {
	IP     *string `json:"ip,omitempty"`
	Method string  `json:"method"`
	URL    string  `json:"url"`
}

// Response represents the standard API response structure
type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Request    RequestInfo `json:"request"`
	Message    string      `json:"message"`
	Data       any         `json:"data"`
}


// ResponseHandler handles the creation and sending of HTTP responses
type ResponseHandler struct {
	cfg    types.Config
}

//TODO: Inject Logger
func NewResponseHandler(cfg types.Config) *ResponseHandler {
	return &ResponseHandler{
		cfg:    cfg,
	}
}

// Send creates and sends an HTTP response using Gin context
func (h *ResponseHandler) Send(c *gin.Context, statusCode int, message string, data any) {
	// Create the response
	var ipAddr *string
	if Environment(h.cfg.Env) != Production {
		ip := c.ClientIP()
		ipAddr = &ip
	}

	response := Response{
		Success:    true,
		StatusCode: statusCode,
		Request: RequestInfo{
			IP:     ipAddr,
			Method: c.Request.Method,
			URL:    c.Request.URL.String(),
		},
		Message: message,
		Data:    data,
	}

	// Log the response
	c.JSON(statusCode, response)
}

// Implement convenience methods for common responses
func (h *ResponseHandler) Ok(c *gin.Context, message string, data any) {
	h.Send(c, 200, message, data)
}

func (h *ResponseHandler) Created(c *gin.Context, message string, data any) {
	h.Send(c, 201, message, data)
}

func (h *ResponseHandler) BadRequest(c *gin.Context, message string, data any) {
	h.Send(c, 400, message, data)
}

func (h *ResponseHandler) NotFound(c *gin.Context, message string, data any) {
	h.Send(c, 404, message, data)
}

func (h *ResponseHandler) InternalServerError(c *gin.Context, message string, data any) {
	h.Send(c, 500, message, data)
}