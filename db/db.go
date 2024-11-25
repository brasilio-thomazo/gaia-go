package db

import (
	"context"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"optimus.dev.br/gaia/model"
	"optimus.dev.br/gaia/repository"
	"optimus.dev.br/gaia/security"
)

type DB struct {
	DB *gorm.DB
}

func NewDB() *DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + name + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return &DB{db}
}

func (db *DB) Migrate() {
	db.DB.AutoMigrate(&model.Group{}, &model.User{}, &model.Customer{}, &model.App{})
}

func (d *DB) InitData() {
	ctx := context.Background()
	gr := repository.NewGroupRepository(d.DB)
	if gr.Count(ctx) == 0 {
		root := &model.Group{Name: "root", Permissions: []string{"root"}, Visible: false, Editable: false, Locked: true}
		nobody := &model.Group{Name: "nobody", Permissions: []string{"nobody"}, Visible: false, Editable: false, Locked: true}
		admin := &model.Group{Name: "admin", Permissions: []string{"admin"}, Visible: true, Editable: false, Locked: true}

		if err := gr.Create(ctx, root); err != nil {
			log.Fatal(err)
		}
		if err := gr.Create(context.Background(), nobody); err != nil {
			log.Fatal(err)
		}
		if err := gr.Create(ctx, admin); err != nil {
			log.Fatal(err)
		}
	}

	ur := repository.NewUserRepository(d.DB)
	if ur.Count(ctx) == 0 {
		rootPwd, err := security.BcryptHashPassword("root")
		if err != nil {
			log.Fatalf("failed to hash password: %v", err)
		}
		rootGrp, err := gr.FindByName(ctx, "root")
		if err != nil {
			log.Fatalf("failed to find root group: %v", err)
		}

		root := &model.User{
			Name:     "root",
			Username: "root",
			Password: rootPwd,
			Email:    "root@change.me",
			Locked:   true,
			Visible:  false,
			Editable: false,
			GroupID:  rootGrp.ID,
		}

		if err := ur.Create(ctx, root); err != nil {
			log.Fatalf("failed to create root user: %v", err)
		}

		adminPwd, err := security.BcryptHashPassword("admin")
		if err != nil {
			log.Fatalf("failed to hash password: %v", err)
		}

		adminGrp, err := gr.FindByName(ctx, "admin")
		if err != nil {
			log.Fatalf("failed to find admin group: %v", err)
		}

		admin := &model.User{
			Name:     "admin",
			Username: "admin",
			Password: adminPwd,
			Email:    "admin@change.me",
			Locked:   true,
			Visible:  true,
			Editable: false,
			GroupID:  adminGrp.ID,
		}
		if err := ur.Create(ctx, admin); err != nil {
			log.Fatalf("failed to create admin user: %v", err)
		}
	}
}
