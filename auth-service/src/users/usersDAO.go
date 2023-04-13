package users

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

type UserDAO struct {
	dbConn *sql.DB
}

// GetAll returns a slice of all users, sorted by last name
func (u *UserDAO) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, active, created_at, updated_at
	from users order by last_name`

	rows, err := u.dbConn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail returns one user by email
func (u *UserDAO) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, password, first_name, last_name, password, active, created_at, updated_at from users where email = $1`

	var user User
	row := u.dbConn.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *UserDAO) GetOne(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, active, created_at, updated_at from users where id = $1`

	var user User
	row := u.dbConn.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *UserDAO) Update(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update users set
		email = $1,
		first_name = $2,
		last_name = $3,
		active = $4,
		updated_at = $5
		where id = $6
	`

	_, err := u.dbConn.ExecContext(ctx, stmt,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Active,
		time.Now(),
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes one user from the database, by User.ID
func (u *UserDAO) Delete(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := u.dbConn.ExecContext(ctx, stmt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *UserDAO) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := u.dbConn.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *UserDAO) Insert(user User) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into users (email, first_name, last_name, password, active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) returning id`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := u.dbConn.QueryRowContext(ctx, stmt,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Active,
		user.UpdatedAt,
		user.CreatedAt,
	).Scan(&newID)

	if err != nil {
		return user, err
	}
	user.ID = newID

	return user, nil
}

func (u *UserDAO) UpdatePasswordByID(hashedPassword []byte, user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update users set password = $1 where id = $2`
	_, err := u.dbConn.ExecContext(ctx, stmt, hashedPassword, user.ID)
	if err != nil {
		return err
	}

	return nil
}
