-- name: CreateRegistration :one
INSERT INTO registrations (
  first_name,
  middle_name,
  last_name,
  phone_number,
  email_address,
  street_address_line1,
  street_address_line2,
  city,
  state,
  zip_code,
  photo_path,
  birth_date,
  appointment_date,
  appointment_time
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
         ) RETURNING *;

-- name: GetRegistration :one
SELECT * FROM registrations
WHERE ID = $1 LIMIT 1;

-- name: GetRegistrationForUpdate :one
SELECT * FROM registrations
WHERE id = $1 LIMIT 1
    FOR NO KEY UPDATE;

-- name: ListRegistrations :many
SELECT * FROM registrations
LIMIT $1
    OFFSET $2;

-- name: UpdateRegistration :one
UPDATE registrations
SET  first_name = $2,
     middle_name = $3,
     last_name = $4,
     phone_number = $5,
     email_address = $6,
     street_address_line1 = $7,
     street_address_line2 = $8,
     city = $9,
     state = $10,
     zip_code = $11,
     photo_path = $12,
     birth_date = $13,
     appointment_date = $14,
     appointment_time = $15
WHERE id = $1
    RETURNING *;

-- name: DeleteRegistrations :exec
DELETE FROM Registrations
WHERE id = $1;

