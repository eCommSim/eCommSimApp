package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DatabasePort_g int64

func getDBConn() *bun.DB {
	// Open a PostgreSQL connection
	dsn := fmt.Sprintf("postgres://auth:mypass@localhost:%d/authenticate?sslmode=disable", DatabasePort_g)
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun DB on top of it
	return bun.NewDB(pgdb, pgdialect.New())
}

func CreateUserTable() (sql.Result, error) {
	db := getDBConn()
	ctx := context.Background()
	res, err := db.NewCreateTable().Model((*user)(nil)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("CreateUserTable:", res)
	return res, nil
}

func insertUser(Email, Username, PasswordHash, Fullname, Role string) error {
	db := getDBConn()
	defer db.Close()

	ctx := context.Background()

	u, _ := getUser(Email)

	if u != nil {
		cTimeStamp := u.CTimeStamp
		u = &user{
			Email:        Email,
			Username:     Username,
			PasswordHash: PasswordHash,
			Fullname:     Fullname,
			CTimeStamp:   cTimeStamp,
			UTimeStamp:   time.Now().String(),
			Role:         Role,
		}
		_, err := db.NewUpdate().Model(u).Column("email").WherePK().Exec(ctx)
		if err != nil {
			return err
		} else {
			return nil
		}
	}

	u = &user{
		Email:        Email,
		Username:     Username,
		PasswordHash: PasswordHash,
		Fullname:     Fullname,
		CTimeStamp:   time.Now().String(),
		UTimeStamp:   time.Now().String(),
		Role:         Role,
	}
	_, err := db.NewInsert().Model(u).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// func UpdateUser(Email, Username, PasswordHash, Fullname, Role string) error {
// 	db := getDBConn()
// 	ctx := context.Background()

// u, err := getUser(Email)
// if err != nil {
// 	return err
// }

// cTimeStamp := u.CTimeStamp

// 	_, err = db.NewUpdate().Model(&user{
// 		Email:        Email,
// 		Username:     Username,
// 		PasswordHash: PasswordHash,
// 		Fullname:     Fullname,
// 		CTimeStamp:   cTimeStamp,
// 		UTimeStamp:   time.Now().GoString(),
// 		Role:         Role,
// 	}).Column("email").WherePK().Exec(ctx)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func getUser(Email string) (*user, error) {
	db := getDBConn()
	defer db.Close()

	ctx := context.Background()

	u := new(user)
	err := db.NewSelect().Model(u).Where("email = ?", Email).Scan(ctx)

	if err != nil {
		return nil, err
	}
	fmt.Println("Found user:", u.Fullname)
	return u, nil
}
