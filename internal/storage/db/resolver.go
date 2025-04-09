package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (r *Storage) AddPVZ(city string) (*model.PVZResp, error) {
	const op = "internal.storage.db.AddPVZ()"

	log.Debug().Msgf("%s start", op)

	pvz := &model.PVZResp{
		ID:               uuid.New(),
		RegistrationDate: time.Now(),
		City:             city,
	}

	queryAddPVZ := `INSERT INTO pvzs (id, registration_date, city)
					VALUES ($1, $2, $3)`

	_, err := r.db.Exec(queryAddPVZ, pvz.ID, pvz.RegistrationDate, pvz.City)
	if err != nil {
		err = fmt.Errorf("%s:%w", op, err)
		return nil, err
	}

	log.Debug().Msgf("%s end", op)

	return pvz, nil

}

func (r *Storage) AddReception(pvzID uuid.UUID) (*model.ReceptionsResp, error) {
	const op = "internal.storage.db.AddReception()"

	log.Debug().Msgf("%s start", op)
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	defer func() {
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				log.Error().Err(errRB).Msg("roll back transaction failed")
			}
		}
	}()

	receptions := &model.ReceptionsResp{
		ID:       uuid.New(),
		DateTime: time.Now(),
		Status:   "in_progress",
	}

	queryOpenReceptions := `SELECT id 
								FROM receptions 
								WHERE pvz_id = $1 	
								AND status = 'in_progress'`

	queryPVZExist := `SELECT id
						FROM pvzs
						WHERE id = $1`

	queryAddReceptions := `INSERT INTO receptions (id, pvz_id, status, registration_date)
						VALUES ($1, $2, $3, $4)`

	err = tx.QueryRow(queryPVZExist, pvzID).Scan(&receptions.PVZID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrPVZNotExist
		}
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	var previousReceptionID uuid.UUID
	err = tx.QueryRow(queryOpenReceptions, pvzID).Scan(&previousReceptionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return nil, fmt.Errorf("%s:%w", op, err)
		}
	}

	if previousReceptionID != uuid.Nil {
		return nil, storage.ErrPreviousReceptionNotClosed
	}

	_, err = tx.Exec(queryAddReceptions, receptions.ID, receptions.PVZID, receptions.Status, receptions.DateTime)

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("%s:%w", op, err)
		return nil, err
	}

	log.Debug().Msgf("%s end", op)

	return receptions, nil

}

func (r *Storage) AddProduct(pvzID uuid.UUID, productType string) (*model.ProductsResp, error) {
	const op = "internal.storage.db.AddProduct()"

	log.Debug().Msgf("%s start", op)
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	defer func() {
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				log.Error().Err(errRB).Msg("roll back transaction failed")
			}
		}
	}()

	product := &model.ProductsResp{
		ID:       uuid.New(),
		DateTime: time.Now(),
		Type:     productType,
	}

	queryOpenReceptions := `SELECT id 
								FROM receptions 
								WHERE pvz_id = $1 	
								AND status = 'in_progress'`

	queryAddReceptions := `INSERT INTO products (id, reception_id, type, registration_date)
						VALUES ($1, $2, $3, $4)`

	err = tx.QueryRow(queryOpenReceptions, pvzID).Scan(&product.ReceptionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.ErrOpenReceptionNotExist
			return nil, err
		} else {
			return nil, fmt.Errorf("%s:%w", op, err)
		}
	}

	_, err = tx.Exec(queryAddReceptions, product.ID, product.ReceptionID, product.Type, product.DateTime)

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("%s:%w", op, err)
		return nil, err
	}

	log.Debug().Msgf("%s end", op)

	return product, nil

}

func (r *Storage) DeleteLastProducts(pvzID uuid.UUID) error {
	op := "internal.storage.db.DeleteLastProducts()"
	log.Debug().Msgf("%s start", op)
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	defer func() {
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				log.Error().Err(errRB).Msg("roll back transaction failed")
			}
		}
	}()

	queryOpenReceptions := `SELECT id 
								FROM receptions 
								WHERE pvz_id = $1 	
								AND status = 'in_progress'`

	queryDeleteLastProduct := `DELETE FROM products
								WHERE id = (SELECT id
								FROM products		
								WHERE reception_id = $1
								ORDER BY registration_date DESC
								LIMIT 1)
								RETURNING id`

	var receptionID uuid.UUID
	err = tx.QueryRow(queryOpenReceptions, pvzID).Scan(&receptionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.ErrOpenReceptionNotExist
			return err
		} else {
			return fmt.Errorf("%s:%w", op, err)
		}
	}

	var productID uuid.UUID
	err = tx.QueryRow(queryDeleteLastProduct, receptionID).Scan(&productID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.ErrProductsInReceptionNotExist
			return err
		} else {
			return fmt.Errorf("%s:%w", op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("%s:%w", op, err)
		return err
	}

	log.Debug().Msgf("productID %s deleted", productID)
	log.Debug().Msgf("%s end", op)

	return nil
}

// this function can write without transaction, because it is not critical for the system
func (r *Storage) CloseLastReceptions(pvzID uuid.UUID) error {
	op := "internal.storage.db.CloseLastReceptions()"
	log.Debug().Msgf("%s start", op)
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	defer func() {
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				log.Error().Err(errRB).Msg("roll back transaction failed")
			}
		}
	}()

	queryOpenReceptions := `SELECT id 
								FROM receptions 
								WHERE pvz_id = $1 	
								AND status = 'in_progress'`

	queryCloseLastReceptions := `UPDATE receptions
								SET status = 'closed'	
								WHERE id = $1`

	var receptionID uuid.UUID
	err = tx.QueryRow(queryOpenReceptions, pvzID).Scan(&receptionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.ErrOpenReceptionNotExist
			return err
		} else {
			return fmt.Errorf("%s:%w", op, err)
		}
	}

	_, err = tx.Exec(queryCloseLastReceptions, &receptionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.ErrProductsInReceptionNotExist
			return err
		} else {
			return fmt.Errorf("%s:%w", op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("%s:%w", op, err)
		return err
	}

	log.Debug().Msgf("receptionID %s closed", receptionID)
	log.Debug().Msgf("%s end", op)

	return nil
}

func (r *Storage) CreateUser(email string, password string, role string) (*model.RegisterResp, error) {
	const op = "internal.storage.db.CreateUser()"

	log.Debug().Msgf("%s start", op)

	passwordHash, err := createPasswordHash(password)
	if err != nil {
		return nil, err
	}

	user := &model.RegisterResp{
		ID:    uuid.New(),
		Email: email,
		Role:  role,
	}

	queryAddUser := `INSERT INTO users (id, email, password_hash, role, registration_date)
					VALUES ($1, $2, $3, $4, $5)`

	_, err = r.db.Exec(queryAddUser, user.ID, user.Email, passwordHash, user.Role, time.Now())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, storage.ErrUserAlreadyExist
		}
		err = fmt.Errorf("%s:%w", op, err)
		return nil, err
	}

	log.Debug().Msgf("%s end", op)

	return user, nil
}

func createPasswordHash(password string) (string, error) {
	const op = "internal.http_server.hadnlers.auth.createPasswordHash()"

	if len(password) < 1 {
		return "", storage.ErrPasswordIsEmpty
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}
	return string(passwordHash), nil
}

func (r *Storage) Login(email string, password string) (userID uuid.UUID, role string, err error) {
	op := "internal.storage.db.Login()"
	log.Debug().Msgf("%s start", op)

	queryLogin := `SELECT id, password_hash, role
					FROM users
					WHERE email = $1`

	var passwordHash string

	err = r.db.QueryRow(queryLogin, email).Scan(&userID, &passwordHash, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, "", storage.ErrUserNotExist
		}
		return uuid.Nil, "", fmt.Errorf("%s:%w", op, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return uuid.Nil, "", storage.ErrPasswordIsWrong
		}
		return uuid.Nil, "", fmt.Errorf("%s:%w", op, err)
	}

	log.Debug().Msgf("%s end", op)
	return userID, role, nil

}
