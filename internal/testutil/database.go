package testutil

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func InitDB() (*sqlx.DB, func()) {
	cfg := loadEnv()

	dbName := uuid.NewString()
	dbCleanup, err := createCleanDB(cfg, dbName)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		dbName,
	)

	sqlxDB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		if err := dbCleanup(); err != nil {
			panic(err)
		}

		if err := sqlxDB.Close(); err != nil {
			panic(err)
		}
	}

	return sqlxDB, cleanup
}

func createCleanDB(cfg *testEnv, dbName string) (func() error, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbName))
	if err != nil {
		return nil, err
	}

	if _, err := sqlDB.Exec(fmt.Sprintf("USE `%s`;", dbName)); err != nil {
		return nil, err
	}

	_, thisFilePath, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("failed to get this file path")
	}
	dumpPath := filepath.Join(filepath.Dir(thisFilePath), "..", "..", "db", "sql", "schema.sql")

	dumpBytes, err := os.ReadFile(dumpPath)
	if err != nil {
		return nil, err
	}

	var regexpNewline = regexp.MustCompile(`\r\n|\r|\n`)
	dumpStr := regexpNewline.ReplaceAllString(string(dumpBytes), "")

	for _, stmt := range strings.Split(dumpStr, ";") {
		if stmt == "" {
			continue
		}
		if _, err := sqlDB.Exec(stmt); err != nil {
			return nil, err
		}
	}

	cleanup := func() error {
		_, err = sqlDB.Exec(fmt.Sprintf("DROP DATABASE `%s`;", dbName))
		if err != nil {
			return err
		}

		return sqlDB.Close()
	}

	return cleanup, nil
}
