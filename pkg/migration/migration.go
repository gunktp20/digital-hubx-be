package main

import (
	"context"
	"log"

	"github.com/gunktp20/digital-hubx-be/pkg/config"
	"github.com/gunktp20/digital-hubx-be/pkg/constant"
	"github.com/gunktp20/digital-hubx-be/pkg/database"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

func main() {
	ctx := context.Background()
	conf := config.GetConfig()

	db := database.NewGormPostgresDatabase(ctx, conf).Db

	// ลบ enum class_tiers (ถ้ามี)
	if err := db.Exec("DROP TYPE IF EXISTS class_tiers CASCADE").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to drop old enum class_tiers: "+constant.Reset, err)
	}

	// ลบ enum question_types (ถ้ามี)
	if err := db.Exec("DROP TYPE IF EXISTS question_types CASCADE").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to drop old enum question_types: "+constant.Reset, err)
	}

	// ลบ enum class_session_status (ถ้ามี)
	if err := db.Exec("DROP TYPE IF EXISTS class_session_status CASCADE").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to drop old enum class_session_status: "+constant.Reset, err)
	}

	// ลบ enum reg_status (ถ้ามี)
	if err := db.Exec("DROP TYPE IF EXISTS reg_status CASCADE").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to drop old enum reg_status: "+constant.Reset, err)
	}

	// สร้าง enum ใหม่สำหรับ question_types
	if err := db.Exec("CREATE TYPE question_types AS ENUM ('choice', 'text')").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to create enum question_types: "+constant.Reset, err)
	}

	// สร้าง enum ใหม่สำหรับ class_tiers
	if err := db.Exec("CREATE TYPE class_tiers AS ENUM ('essential', 'literacy', 'mastery')").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to create enum class_tiers: "+constant.Reset, err)
	}

	// สร้าง enum ใหม่สำหรับ class_session_status
	if err := db.Exec("CREATE TYPE class_session_status AS ENUM ('available','fullybook')").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to create enum class_session_status: "+constant.Reset, err)
	}

	// สร้าง enum ใหม่สำหรับ reg_status
	if err := db.Exec("CREATE TYPE reg_status AS ENUM ('registered','completed')").Error; err != nil {
		log.Fatalln(constant.Red+"Failed to create enum reg_status: "+constant.Reset, err)
	}

	// AutoMigrate สำหรับโมเดล
	err := db.AutoMigrate(
		// &models.AppGroup{}, &models.AppPreviewImage{}, &models.AppProcess{}, &models.AppSubProcess{}, &models.AppType{},&models.App{},
		&models.ClassHighLightImage{},
		&models.ClassSession{},
		&models.Class{},
		// &models.Process{},
		// &models.Solution{},
		// &models.SubProcess{},
		//  &models.Type{},
		// &models.UserAppFavorite{},
		&models.UserClassRegistration{},
		&models.UserQuestionAnswer{},
		&models.Choice{},
		&models.Question{},
		&models.User{},
	)

	if err != nil {
		log.Fatalln(constant.Red+"Failed to migrate database: "+constant.Reset, err)
	} else {
		log.Fatalln(constant.Green + "Database migration successful" + constant.Reset)
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()
}
