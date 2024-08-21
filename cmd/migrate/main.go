package main

import (
	"edot-monorepo/services/user-service/internal/config"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"path/filepath"

	"gorm.io/gorm"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)

	action := "up" // Default action or read from args
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	err := migrate(db, action)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations applied successfully")
}

func migrate(db *gorm.DB, action string) error {
	migrationFiles, err := listMigrationFiles("db/migrations", action)
	if err != nil {
		return fmt.Errorf("failed to list migration files: %w", err)
	}

	for _, file := range migrationFiles {
		sqlContent, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		if err := executeSQL(db, string(sqlContent)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}

		fmt.Printf("Successfully executed migration %s\n", file)
	}

	return nil
}

func listMigrationFiles(dir string, action string) ([]string, error) {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	prefix := fmt.Sprintf("%s.sql", action)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		if strings.HasSuffix(fileName, prefix) {
			files = append(files, filepath.Join(dir, fileName))
		}
	}

	return files, nil
}

func executeSQL(db *gorm.DB, sqlContent string) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}

	_, err = sqlDB.Exec(sqlContent)
	if err != nil {
		return fmt.Errorf("failed to execute SQL: %w", err)
	}

	return nil
}
