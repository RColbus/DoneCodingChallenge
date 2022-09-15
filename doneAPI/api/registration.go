package api

import (
	"database/sql"
	db "doneAPI/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createRegistrationRequest creates new registration request object
type createRegistrationRequest struct {
	FirstName          string `json:"first_name" binding:"required"`
	MiddleName         string `json:"middle_name"`
	LastName           string `json:"last_name" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	EmailAddress       string `json:"email_address" binding:"required"`
	StreetAddressLine1 string `json:"street_address_line1" binding:"required"`
	StreetAddressLine2 string `json:"street_address_line2"`
	City               string `json:"city" binding:"required"`
	State              string `json:"state" binding:"required"`
	ZipCode            string `json:"zip_code" binding:"required"`
	PhotoPath          string `json:"photo_path" binding:"required"`
	BirthDate          string `json:"birth_date" binding:"required"`
	AppointmentDate    string `json:"appointment_date" binding:"required"`
	AppointmentTime    string `json:"appointment_time" binding:"required"`
}

// createRegistration creates new registration
func (server *Server) createRegistration(ctx *gin.Context) {
	var req createRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRegistrationParams{
		FirstName:          req.FirstName,
		LastName:           req.LastName,
		MiddleName:         req.MiddleName,
		PhoneNumber:        req.PhoneNumber,
		EmailAddress:       req.EmailAddress,
		BirthDate:          req.BirthDate,
		StreetAddressLine1: req.StreetAddressLine1,
		StreetAddressLine2: req.StreetAddressLine2,
		City:               req.City,
		State:              req.State,
		ZipCode:            req.ZipCode,
		AppointmentDate:    req.AppointmentDate,
		AppointmentTime:    req.AppointmentTime,
		PhotoPath:          req.PhotoPath,
	}

	registration, err := server.store.CreateRegistration(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, registration)
}

// GetRegistrationRequest creates new registration request object
type GetRegistrationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getRegistration gets registration by ID
func (server *Server) getRegistration(ctx *gin.Context) {
	var req GetRegistrationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	registration, err := server.store.GetRegistration(ctx, req.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			// No rows returned
			ctx.JSON(http.StatusNotFound, errorResponse(err))

		}
		// Server Error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, registration)
}

// ListRegistrationRequest creates new registration request object
type ListRegistrationRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

// listRegistration returns a list of registrations by page
func (server *Server) listRegistration(ctx *gin.Context) {
	var req ListRegistrationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRegistrationsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	registrations, err := server.store.ListRegistrations(ctx, arg)
	if err != nil {
		// Server Error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, registrations)
}

// UpdateRegistration updates registration
func (server *Server) UpdateRegistration(ctx *gin.Context) {
	// Not implemented in the code sample
}

// DeleteRegistration deletes registration
func (server *Server) DeleteRegistration(ctx *gin.Context) {
	// Not implemented in the code sample
}
