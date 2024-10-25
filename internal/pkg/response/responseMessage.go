package response

type ResponseMessage struct {
	Success            string
	Created            string
	Updated            string
	Deleted            string
	BadRequest         string
	Unauthorized       string
	Forbidden          string
	NotFound           string
	Conflict           string
	InternalError      string
	ServiceUnavailable string
	ValidationError    string
	RateLimited        string
	Maintenance        string
	Timeout            string
	PartialContent     string
}

// Messages provides centralized response messages
var Messages = ResponseMessage{
	Success:            "Operation completed successfully",
	Created:            "Resource created successfully",
	Updated:            "Resource updated successfully",
	Deleted:            "Resource deleted successfully",
	BadRequest:         "Invalid request parameters",
	Unauthorized:       "Authentication required",
	Forbidden:          "Access denied",
	NotFound:           "Resource not found",
	Conflict:           "Resource already exists",
	InternalError:      "An unexpected error occurred",
	ServiceUnavailable: "Service temporarily unavailable",
	ValidationError:    "Data validation failed",
	RateLimited:        "Too many requests, please try again later",
	Maintenance:        "System under maintenance, please try again soon",
	Timeout:            "Operation timed out",
	PartialContent:     "Partial content returned",
}