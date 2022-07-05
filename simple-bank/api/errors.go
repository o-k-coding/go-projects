package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// NOTE: this is being used as a general purpose function, if each use case needed some control over different types of errors, this should be moved back out, and that is fine.
func getErrorStatusForDBError(pqErr *pq.Error) int {
	switch pqErr.Code.Name() {
	case "foreign_key_violation", "unique_violation":
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

func handleApiError(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError
	// try to convert to pq Error, if the conversion is ok
	if pqErr, ok := err.(*pq.Error); ok {
		status = getErrorStatusForDBError(pqErr)
	}
	ctx.JSON(status, errorResponse(err))
}
