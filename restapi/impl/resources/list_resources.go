package resources

import (
	"database/sql"

	"github.com/cyverse-de/permissions/logger"
	"github.com/cyverse-de/permissions/models"
	permsdb "github.com/cyverse-de/permissions/restapi/impl/db"
	"github.com/cyverse-de/permissions/restapi/operations/resources"

	"github.com/go-openapi/runtime/middleware"
)

// BuildListResourcesHandler builds the request handler for the list resources endpoint.
func BuildListResourcesHandler(db *sql.DB) func(resources.ListResourcesParams) middleware.Responder {

	// Return the handler function.
	return func(params resources.ListResourcesParams) middleware.Responder {

		// Start a transaction for this request.
		tx, err := db.Begin()
		if err != nil {
			logger.Log.Error(err)
			reason := err.Error()
			return resources.NewListResourcesInternalServerError().WithPayload(
				&models.ErrorOut{Reason: &reason},
			)
		}
		defer tx.Commit()

		// List all resources.
		result, err := permsdb.ListResources(tx, params.ResourceTypeName, params.ResourceName)
		if err != nil {
			logger.Log.Error(err)
			reason := err.Error()
			return resources.NewListResourcesInternalServerError().WithPayload(
				&models.ErrorOut{Reason: &reason},
			)
		}

		// Return the results.
		return resources.NewListResourcesOK().WithPayload(&models.ResourcesOut{Resources: result})
	}
}
