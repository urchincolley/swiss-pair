package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
)

// Analyze error to determine status code
// This is a bad way to do this
func HttpStatusFromError(e error) int {
	if strings.Contains(e.Error(), "duplicate key value violates unique constraint") {
		return http.StatusForbidden
	}
	if errors.Is(e, sql.ErrNoRows) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
