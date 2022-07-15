package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/phuslu/log"
	"github.com/rostikts/social_network/domain/model"
	"github.com/rostikts/social_network/pkg/user"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) user.Repository {
	return userRepository{db: db}
}

func (r userRepository) Create(user model.User) (model.User, error) {
	rows, err := r.db.NamedQuery(`INSERT INTO users(username, first_name, last_name, email, password) VALUES (:username, :first_name, :last_name, :email, :password) returning id`, user)
	if err != nil {
		return model.User{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID); err != nil {
			log.DefaultLogger.Error().Err(err).Msg("error during scanning user's id")
		}
	}
	return user, nil
}

func (r userRepository) Update(user model.User) (model.User, error) {
	res, err := r.db.NamedExec(`UPDATE users u SET username=:username, first_name=:first_name, last_name=:last_name, email=:email WHERE u.id=:id`, user)
	if err != nil {
		log.DefaultLogger.Info().Err(err).Msg("User got err during data update")
		return model.User{}, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return model.User{}, err
	}
	if rows == 0 {
		return model.User{}, sql.ErrNoRows
	}

	return user, nil
}

func (r userRepository) UpdatePassword(user model.User) error {
	res, err := r.db.NamedExec(`UPDATE users u SET password=:password WHERE u.id=:id`, user)
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r userRepository) FindByID(id uint) (model.User, error) {
	var usr model.User
	err := r.db.Get(&usr, `SELECT u.id, u.username, u.first_name, u.last_name, u.email FROM users u WHERE u.id=$1`, id)
	if err != nil {
		return model.User{}, err
	}
	return usr, nil
}

func (r userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, `SELECT u.id, u.username, u.first_name, u.last_name, u.email FROM users u`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
