package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	dao UserDAO
}

func (u *UserService) GetByID(id int) (*User, error) {
	return u.dao.GetOne(id)
}

func (u *UserService) GetAll() ([]*User, error) {
	return u.dao.GetAll()
}

func (u *UserService) DeleteByID(id int) error {
	return u.dao.DeleteByID(id)
}

func (u *UserService) Insert(user User) (User, error) {
	hashPassword, err := u.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = string(hashPassword)
	log.Printf("Inserting user %+v\n", user)
	return u.dao.Insert(user)
}

func (u *UserService) ResetPassword(password string, id int) error {
	user, err := u.dao.GetOne(id)
	if err != nil {
		return err
	}

	hashedPassword, err := u.HashPassword(password)
	if err != nil {
		return err
	}

	return u.dao.UpdatePasswordByID(hashedPassword, user)
}

func (*UserService) HashPassword(password string) ([]byte, error) {
	log.Printf("Hashing password %s\n", password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return hashedPassword, err
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *UserService) GetByEmailAndCheckPassword(email string, password string) (*User, error) {
	user, err := u.dao.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	log.Printf("Got user %+v\n", user)

	log.Printf("Comparing password %s with hash %s", password, user.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
		// // switch {
		// // case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
		// // 	return false, nil
		// // default:
		// // 	return false, err
		// }
	}

	return user, nil
}
