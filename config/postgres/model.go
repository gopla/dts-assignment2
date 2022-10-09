package postgres

import (
  "fmt"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type Config struct {
  Host         string `json:"host"`
  Port         string `json:"port"`
  DatabaseName string `json:"database_name"`
  User         string `json:"user"`
  Password     string `json:"password"`
}

type PostgresClient interface {
  GetClient() *gorm.DB
  MigrateDB(models ...interface{})
}

type PostgresClientImpl struct {
  cln    *gorm.DB
  config Config
}

func NewPostgresConnection(config Config) PostgresClient {
  connectionString := fmt.Sprintf(`
	host=%s 
	port=%s
	user=%s 
	password=%s 
	dbname=%s`,
    config.Host,
    config.Port,
    config.User,
    config.Password,
    config.DatabaseName)

  db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

  if err != nil {
    panic(err)
  }
  return &PostgresClientImpl{cln: db, config: config}
}

func (p *PostgresClientImpl) GetClient() *gorm.DB {
  return p.cln
}

func (p *PostgresClientImpl) MigrateDB(models ...interface{}) {
  p.cln.AutoMigrate(models...)
}
