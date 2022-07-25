package daos

import (
	"context"
	"go-template/models"
)

// FindIncidentReportByID
func FindIncidentReportByID(incidentReportID int) (*models.IncidentReport, error) {
	contextExecutor := getContextExecutor(nil)
	return models.FindIncidentReport(context.Background(), contextExecutor, incidentReportID)
}
