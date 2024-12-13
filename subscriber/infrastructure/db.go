package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"server/env"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	host         string
	username     string
	password     string
	database     string
	port         string
	timezone     string
	maxIdle      int
	maxOpen      int
	maxLifetime  time.Duration
	maxRetries   int
	retryTimeout time.Duration
}

func newDBConfig() DBConfig {
	e := env.NewEnv()
	return DBConfig{
		host:         e.POSTGRE_HOST,
		username:     e.POSTGRE_USERNAME,
		password:     e.POSTGRE_PASSWORD,
		database:     e.POSTGRE_DATABASE,
		port:         e.POSTGRE_PORT,
		timezone:     e.POSTGRE_TIMEZONE,
		maxIdle:      e.POSTGRE_CONN_MAX_IDLE,
		maxOpen:      e.POSTGRE_CONN_MAX_OPEN,
		maxLifetime:  time.Minute * e.POSTGRE_CONN_MAX_LIFETIME,
		maxRetries:   3,
		retryTimeout: time.Second * 5,
	}
}

func (c *DBConfig) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		c.host, c.username, c.password, c.database, c.port, c.timezone)
}

func connectWithRetry(ctx context.Context, connect func() error) error {
	config := newDBConfig()
	var err error
	for i := 0; i < config.maxRetries; i++ {
		if err = connect(); err == nil {
			return nil
		}
		log.Printf("Failed to connect to database, attempt %d/%d: %v", i+1, config.maxRetries, err)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(config.retryTimeout):
		}
	}
	return fmt.Errorf("failed to connect after %d attempts: %v", config.maxRetries, err)
}

func ConnectSqlDB() (*sql.DB, error) {
	config := newDBConfig()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var db *sql.DB
	connect := func() error {
		var err error
		db, err = sql.Open("postgres", config.getDSN())
		if err != nil {
			return fmt.Errorf("failed to open connection: %w", err)
		}

		db.SetMaxIdleConns(config.maxIdle)
		db.SetMaxOpenConns(config.maxOpen)
		db.SetConnMaxLifetime(config.maxLifetime)

		return db.PingContext(ctx)
	}

	if err := connectWithRetry(ctx, connect); err != nil {
		return nil, fmt.Errorf("sql database connection failed: %w", err)
	}

	log.Println("sql database: connection established")
	return db, nil
}

func ConnectSqlxDB() (*sqlx.DB, error) {
	config := newDBConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var db *sqlx.DB
	connect := func() error {
		var err error
		db, err = sqlx.ConnectContext(ctx, "postgres", config.getDSN())
		if err != nil {
			return fmt.Errorf("failed to connect: %w", err)
		}

		db.SetMaxIdleConns(config.maxIdle)
		db.SetMaxOpenConns(config.maxOpen)
		db.SetConnMaxLifetime(config.maxLifetime)

		return db.PingContext(ctx)
	}

	if err := connectWithRetry(ctx, connect); err != nil {
		return nil, fmt.Errorf("sqlx database connection failed: %w", err)
	}

	log.Println("sqlx database: connection established")
	return db, nil
}
