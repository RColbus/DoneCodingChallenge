// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateRegistration(ctx context.Context, arg CreateRegistrationParams) (Registration, error)
	DeleteRegistrations(ctx context.Context, id int64) error
	GetRegistration(ctx context.Context, id int64) (Registration, error)
	GetRegistrationForUpdate(ctx context.Context, id int64) (Registration, error)
	ListRegistrations(ctx context.Context, arg ListRegistrationsParams) ([]Registration, error)
	UpdateRegistration(ctx context.Context, arg UpdateRegistrationParams) (Registration, error)
}

var _ Querier = (*Queries)(nil)