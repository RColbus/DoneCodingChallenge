package db

import (
	"context"
	"database/sql"
	"doneAPI/util"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomRegistration(t *testing.T) Registration {
	arg := CreateRegistrationParams{
		FirstName:          util.RandomString(8),
		LastName:           util.RandomString(10),
		MiddleName:         util.RandomString(5),
		PhoneNumber:        util.RandomPhoneNumber(),
		EmailAddress:       util.RandomEmail(),
		BirthDate:          util.RandomDate(),
		StreetAddressLine1: util.RandomStreetAddress1(),
		StreetAddressLine2: util.RandomStreetAddress2(),
		City:               util.RandomString(6),
		State:              util.RandomState(),
		ZipCode:            util.RandomIntString(5),
		AppointmentDate:    util.RandomDate(),
		AppointmentTime:    util.RandomAppointmentTime(),
		PhotoPath:          util.RandomPath(),
	}

	reg, err := testQueries.CreateRegistration(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, reg) // verify not an empty object

	// verify saved fields equal args
	require.Equal(t, arg.FirstName, reg.FirstName)
	require.Equal(t, arg.LastName, reg.LastName)
	require.Equal(t, arg.MiddleName, reg.MiddleName)
	require.Equal(t, arg.PhoneNumber, reg.PhoneNumber)
	require.Equal(t, arg.EmailAddress, reg.EmailAddress)
	require.Equal(t, arg.BirthDate, reg.BirthDate)
	require.Equal(t, arg.StreetAddressLine1, reg.StreetAddressLine1)
	require.Equal(t, arg.StreetAddressLine2, reg.StreetAddressLine2)
	require.Equal(t, arg.City, reg.City)
	require.Equal(t, arg.State, reg.State)
	require.Equal(t, arg.ZipCode, reg.ZipCode)
	require.Equal(t, arg.AppointmentDate, reg.AppointmentDate)
	require.Equal(t, arg.AppointmentTime, reg.AppointmentTime)
	require.Equal(t, arg.PhotoPath, reg.PhotoPath)

	require.NotZero(t, reg.ID)
	require.NotZero(t, reg.CreatedAt)

	return reg
}

func TestCreateRegistration(t *testing.T) {
	createRandomRegistration(t)
}

func TestGetRegistration(t *testing.T) {
	account1 := createRandomRegistration(t)
	account2, err := testQueries.GetRegistration(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.FirstName, account2.FirstName)
	require.Equal(t, account1.MiddleName, account2.MiddleName)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.BirthDate, account2.BirthDate)
	require.Equal(t, account1.EmailAddress, account2.EmailAddress)
	require.Equal(t, account1.StreetAddressLine1, account2.StreetAddressLine1)
	require.Equal(t, account1.StreetAddressLine2, account2.StreetAddressLine2)
	require.Equal(t, account1.City, account2.City)
	require.Equal(t, account1.State, account2.State)
	require.Equal(t, account1.ZipCode, account2.ZipCode)
	require.Equal(t, account1.AppointmentDate, account2.AppointmentDate)
	require.Equal(t, account1.AppointmentDate, account2.AppointmentDate)
	require.Equal(t, account1.PhotoPath, account2.PhotoPath)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)

	fmt.Println(account2)

}

func TestUpdateRegistration(t *testing.T) {
	account1 := createRandomRegistration(t)

	arg := UpdateRegistrationParams{
		ID:                 account1.ID,
		FirstName:          util.RandomString(8),
		LastName:           util.RandomString(10),
		MiddleName:         util.RandomString(5),
		PhoneNumber:        util.RandomPhoneNumber(),
		EmailAddress:       util.RandomEmail(),
		BirthDate:          util.RandomDate(),
		StreetAddressLine1: util.RandomStreetAddress1(),
		StreetAddressLine2: util.RandomStreetAddress2(),
		City:               util.RandomString(6),
		State:              util.RandomState(),
		ZipCode:            util.RandomIntString(5),
		AppointmentDate:    util.RandomDate(),
		AppointmentTime:    util.RandomAppointmentTime(),
		PhotoPath:          util.RandomPath(),
	}

	account2, err := testQueries.UpdateRegistration(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.FirstName, account2.FirstName)
	require.Equal(t, arg.MiddleName, account2.MiddleName)
	require.Equal(t, arg.LastName, account2.LastName)
	require.Equal(t, arg.BirthDate, account2.BirthDate)
	require.Equal(t, arg.EmailAddress, account2.EmailAddress)
	require.Equal(t, arg.StreetAddressLine1, account2.StreetAddressLine1)
	require.Equal(t, arg.StreetAddressLine2, account2.StreetAddressLine2)
	require.Equal(t, arg.City, account2.City)
	require.Equal(t, arg.State, account2.State)
	require.Equal(t, arg.ZipCode, account2.ZipCode)
	require.Equal(t, arg.AppointmentDate, account2.AppointmentDate)
	require.Equal(t, arg.AppointmentDate, account2.AppointmentDate)
	require.Equal(t, arg.PhotoPath, account2.PhotoPath)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)

}

func TestDeleteRegistration(t *testing.T) {
	account1 := createRandomRegistration(t)

	err := testQueries.DeleteRegistrations(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetRegistration(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListRegistrations(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRegistration(t)
	}

	arg := ListRegistrationsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListRegistrations(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
