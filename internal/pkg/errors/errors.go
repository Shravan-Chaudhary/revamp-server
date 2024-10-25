package errors

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// HttpError represents the structure of our HTTP errors
type HttpError struct {
    Success    bool   `json:"success"`
    StatusCode int    `json:"statusCode"`
    Request    struct {
        IP     *string `json:"ip,omitempty"`
        Method string  `json:"method"`
        URL    string  `json:"url"`
    } `json:"request"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
    Trace   *struct {
        Time  time.Time `json:"time"`
    } `json:"trace,omitempty"`
}

// Error implements the error interface
func (e *HttpError) Error() string {
    return e.Message
}

// CreateError provides factory functions for common HTTP errors
type CreateError struct{}

var HttpErrors = CreateError{}

// Common error creation methods
func (e CreateError) BadRequest(message string) *HttpError {
    return createError(http.StatusBadRequest, message)
}

func (e CreateError) Conflict(message string) *HttpError {
    return createError(http.StatusConflict, message)
}

func (e CreateError) Unauthorized(message string) *HttpError {
    return createError(http.StatusUnauthorized, message)
}

func (e CreateError) Forbidden(message string) *HttpError {
    return createError(http.StatusForbidden, message)
}

func (e CreateError) NotFound(message string) *HttpError {
    return createError(http.StatusNotFound, message)
}

func (e CreateError) InternalServer(message string) *HttpError {
    return createError(http.StatusInternalServerError, message)
}

func (e CreateError) DatabaseError(message string) *HttpError {
    if message == "" {
        message = "database error occurred"
    }
    return createError(http.StatusInternalServerError, message)
}

// createError is an internal helper to create error instances
func createError(status int, message string) *HttpError {
    err := &HttpError{
        Success:    false,
        StatusCode: status,
        Message:    message,
        Data:       nil,
    }

    // Capture stack trace
    var stack [4096]byte
    runtime.Stack(stack[:], false)

    err.Trace = &struct {
        Time  time.Time `json:"time"`
    }{
        Time:  time.Now(),
    }


    return err
}

// ErrorHandler middleware for Gin
func ErrorHandler(isDevelopment bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        // Check if there are any errors
        if len(c.Errors) > 0 {
            err := c.Errors.Last()

            var httpError *HttpError
            if customErr, ok := err.Err.(*HttpError); ok {
                httpError = customErr
            } else {
                // Convert regular error to HttpError
                httpError = HttpErrors.InternalServer(err.Error())
            }

            // Add request information
            ip := c.ClientIP()
            httpError.Request.Method = c.Request.Method
            httpError.Request.URL = c.Request.URL.String()
            if isDevelopment {
                httpError.Request.IP = &ip
            } else {
                httpError.Trace = nil // Remove trace in production
            }

            c.JSON(httpError.StatusCode, httpError)
            c.Abort()
        }
    }
}