package repositories

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/pressly/goose"

	"gorm.io/driver/postgres"
	gormv2 "gorm.io/gorm"
)

const (
	driver       = "postgres"
	migrationDir = "./migrations/sql"
)

func GetDB() (*gormv2.DB, error) {
	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DATABASE")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gormv2.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}

func usage() {
	const (
		usageRun      = `goose [OPTIONS] COMMAND`
		usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version`
	)
	fmt.Println(usageRun)
	flag.PrintDefaults()
	fmt.Println(usageCommands)
}

func Migrate() error {
	flag.Usage = usage

	flag.Parse()
	args := flag.Args()

	dbString := NewConnectionString()
	fmt.Println("dbStrinng....", dbString)
	db, err := sql.Open(driver, dbString)
	if err != nil {
		return err
	}

	defer db.Close()

	if err = goose.SetDialect(driver); err != nil {
		return fmt.Errorf("failed to set goose dialect: %v", err)
	}

	if len(args) == 0 {
		return errors.New("expected at least one arg")
	}

	command := args[0]

	if err = goose.Run(command, db, migrationDir, args[1:]...); err != nil {
		return fmt.Errorf("goose run: %v", err)
	}
	return db.Close()
}

func NewConnectionString() string {
	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DATABASE")
	)
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
}
