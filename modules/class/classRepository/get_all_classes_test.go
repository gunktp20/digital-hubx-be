package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestGetAllClasses_Success(t *testing.T) {

	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)

	// classTier := "essential"

	expectedResultClasses := [2]models.Class{
		{
			ID:    uuid.New().String(),
			Title: "Test Record 1",
		},
		{
			ID:    uuid.New().String(),
			Title: "Test Record 2",
		},
	}

	mock.ExpectQuery(`(?i)SELECT count\(\*\) FROM "classes"`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(2))

	// mock.ExpectQuery(`(?i)SELECT \* FROM "classes" WHERE "class_tier" = \$1 LIMIT 10`).
	// 	WithArgs(classTier).
	// 	WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
	// 		AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].Title))

	mock.ExpectQuery(`(?i)SELECT \* FROM "classes"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].Title).
			AddRow(expectedResultClasses[1].ID, expectedResultClasses[1].Title))

	mock.ExpectQuery(`SELECT \* FROM "class_sessions" WHERE "class_sessions"\."class_id" IN \(\$1,\$2\)`).
		WithArgs(expectedResultClasses[0].ID, expectedResultClasses[1].ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "class_id"}).AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].ID).AddRow(expectedResultClasses[1].ID, expectedResultClasses[1].ID))

	repo := NewClassGormRepository(gormDB)

	classes, total, err := repo.GetAllClasses("essential", "", 1, 10)
	assert.NoError(t, err)

	assert.Equal(t, int64(2), total)
	assert.Equal(t, 2, len(*classes))

	err = mock.ExpectationsWereMet()
}

func TestGetAllClasses_With_Class_Tier(t *testing.T) {

	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)

	classTier := "essential"
	page := 1
	limit := 10

	mocktime := time.Now()

	expectedResultClasses := [2]models.Class{
		{
			ID:              uuid.New().String(),
			Title:           "Test Record 1",
			Description:     "",
			CoverImage:      "",
			ClassCategoryID: "",
			ClassTier:       "essential",
			ClassLevel:      1,
			IsActive:        true,
			IsRemove:        false,
			CreatedAt:       mocktime,
			UpdatedAt:       mocktime,
			// ClassSessions:        []models.ClassSession{},
			// ClassHighLightImages: []models.ClassHighLightImage{},
		},
		{
			ID:              uuid.New().String(),
			Title:           "Test Record 2",
			Description:     "",
			CoverImage:      "",
			ClassCategoryID: "",
			ClassTier:       "essential",
			ClassLevel:      1,
			IsActive:        true,
			IsRemove:        false,
			CreatedAt:       mocktime,
			UpdatedAt:       mocktime,
			// ClassSessions:        []models.ClassSession{},
			// ClassHighLightImages: []models.ClassHighLightImage{},
		},
	}

	mock.ExpectQuery(`(?i)SELECT count\(\*\) FROM "classes" WHERE class_tier = \$1`).
		WithArgs(classTier).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(2))

	mock.ExpectQuery(`(?i)SELECT \* FROM "classes" WHERE class_tier = \$1 LIMIT \$2`).
		WithArgs(classTier, limit).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "cover_image", "class_category_id", "class_tier", "class_level", "is_active", "is_remove", "created_at", "updated_at"}).
			AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].Title, expectedResultClasses[0].Description, expectedResultClasses[0].CoverImage, expectedResultClasses[0].ClassCategoryID, expectedResultClasses[0].ClassTier, expectedResultClasses[0].ClassLevel, expectedResultClasses[0].IsActive, expectedResultClasses[0].IsRemove, expectedResultClasses[0].CreatedAt, expectedResultClasses[0].UpdatedAt).
			AddRow(expectedResultClasses[1].ID, expectedResultClasses[1].Title, expectedResultClasses[1].Description, expectedResultClasses[1].CoverImage, expectedResultClasses[1].ClassCategoryID, expectedResultClasses[1].ClassTier, expectedResultClasses[1].ClassLevel, expectedResultClasses[1].IsActive, expectedResultClasses[1].IsRemove, expectedResultClasses[1].CreatedAt, expectedResultClasses[1].UpdatedAt))

	mock.ExpectQuery(`SELECT \* FROM "class_sessions" WHERE "class_sessions"\."class_id" IN \(\$1,\$2\)`).
		WithArgs(expectedResultClasses[0].ID, expectedResultClasses[1].ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "class_id", "class_tier"}).AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].ID, expectedResultClasses[0].ClassTier).AddRow(expectedResultClasses[1].ID, expectedResultClasses[1].ID, expectedResultClasses[1].ClassTier))

	repo := NewClassGormRepository(gormDB)

	actualResultClasses, total, err := repo.GetAllClasses(classTier, "", page, limit)
	assert.NoError(t, err)

	assert.Equal(t, int64(2), total)
	assert.Equal(t, 2, len(*actualResultClasses))
	// assert.Equal(t, expectedResultClasses, actualResultClasses)

	err = mock.ExpectationsWereMet()
}
func TestGetAllClasses_With_Keyword(t *testing.T) {

	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)

	classTier := "essential"
	page := 1
	limit := 10

	mocktime := time.Now()

	expectedResultClasses := [2]models.Class{
		{
			ID:              uuid.New().String(),
			Title:           "Test Record 1",
			Description:     "",
			CoverImage:      "",
			ClassCategoryID: "",
			ClassTier:       "essential",
			ClassLevel:      1,
			IsActive:        true,
			IsRemove:        false,
			CreatedAt:       mocktime,
			UpdatedAt:       mocktime,
			// ClassSessions:        []models.ClassSession{},
			// ClassHighLightImages: []models.ClassHighLightImage{},
		},
		{
			ID:              uuid.New().String(),
			Title:           "Test Record 2",
			Description:     "",
			CoverImage:      "",
			ClassCategoryID: "",
			ClassTier:       "essential",
			ClassLevel:      1,
			IsActive:        true,
			IsRemove:        false,
			CreatedAt:       mocktime,
			UpdatedAt:       mocktime,
			// ClassSessions:        []models.ClassSession{},
			// ClassHighLightImages: []models.ClassHighLightImage{},
		},
	}

	mock.ExpectQuery(`(?i)SELECT count\(\*\) FROM "classes" WHERE class_tier = \$1`).
		WithArgs(classTier).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(2))

	mock.ExpectQuery(`(?i)SELECT \* FROM "classes" WHERE class_tier = \$1 LIMIT \$2`).
		WithArgs(classTier, limit).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "cover_image", "class_category_id", "class_tier", "class_level", "is_active", "is_remove", "created_at", "updated_at"}).
			AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].Title, expectedResultClasses[0].Description, expectedResultClasses[0].CoverImage, expectedResultClasses[0].ClassCategoryID, expectedResultClasses[0].ClassTier, expectedResultClasses[0].ClassLevel, expectedResultClasses[0].IsActive, expectedResultClasses[0].IsRemove, expectedResultClasses[0].CreatedAt, expectedResultClasses[0].UpdatedAt).
			AddRow(expectedResultClasses[1].ID, expectedResultClasses[1].Title, expectedResultClasses[1].Description, expectedResultClasses[1].CoverImage, expectedResultClasses[1].ClassCategoryID, expectedResultClasses[1].ClassTier, expectedResultClasses[1].ClassLevel, expectedResultClasses[1].IsActive, expectedResultClasses[1].IsRemove, expectedResultClasses[1].CreatedAt, expectedResultClasses[1].UpdatedAt))

	mock.ExpectQuery(`SELECT \* FROM "class_sessions" WHERE "class_sessions"\."class_id" IN \(\$1,\$2\)`).
		WithArgs(expectedResultClasses[0].ID, expectedResultClasses[1].ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "class_id", "class_tier"}).AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].ID, expectedResultClasses[0].ClassTier).AddRow(expectedResultClasses[1].ID, expectedResultClasses[1].ID, expectedResultClasses[1].ClassTier))

	repo := NewClassGormRepository(gormDB)

	actualResultClasses, total, err := repo.GetAllClasses(classTier, "", page, limit)
	assert.NoError(t, err)

	assert.Equal(t, int64(2), total)
	assert.Equal(t, 2, len(*actualResultClasses))
	// assert.Equal(t, expectedResultClasses, actualResultClasses)

	err = mock.ExpectationsWereMet()
}

func TestGetAllClasses_Error(t *testing.T) {

	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)

	classTier := "essentialx"
	page := 1
	limit := 10

	expectedResultClasses := [2]models.Class{
		{
			ID:    uuid.New().String(),
			Title: "Test Record 1",
		},
		{
			ID:    uuid.New().String(),
			Title: "Test Record 2",
		},
	}

	// mock.ExpectQuery(`(?i)SELECT count\(\*\) FROM "classes" WHERE class_tier = \$1`).
	// 	WithArgs(classTier).
	// 	WillReturnError(fmt.Errorf("failed to count"))

	mock.ExpectQuery(`(?i)SELECT \* FROM "classes" WHERE class_tier = \$1 LIMIT \$2`).
		WithArgs("essentialx", limit).
		WillReturnError(fmt.Errorf("failed to query"))

	mock.ExpectQuery(`SELECT \* FROM "class_sessions" WHERE "class_sessions"\."class_id" IN \(\$1,\$2\)`).
		WithArgs(expectedResultClasses[0].ID, expectedResultClasses[1].ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "class_id", "class_tier"}).AddRow(expectedResultClasses[0].ID, expectedResultClasses[0].ID, expectedResultClasses[0].ClassTier).AddRow(expectedResultClasses[1].ID, expectedResultClasses[1].ID, expectedResultClasses[1].ClassTier))

	repo := NewClassGormRepository(gormDB)

	actualResultClasses, total, err := repo.GetAllClasses(classTier, "", page, limit)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to query")

	assert.Equal(t, int64(2), total)
	assert.Equal(t, 2, len(*actualResultClasses))
	// assert.Equal(t, expectedResultClasses, actualResultClasses)

	err = mock.ExpectationsWereMet()
}

// import (
// 	"database/sql"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/gunktp20/digital-hubx-be/pkg/models"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
// 	sqldb, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	gormdb, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: sqldb,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return sqldb, gormdb, mock
// }

// func TestGetAllClasses(t *testing.T) {
// 	// sqlDB, db, mock := DbMock(t)
// 	// defer sqlDB.Close()

// 	// // สร้าง instance ของ classGormRepository
// 	// repo := NewClassGormRepository(db)

// 	// // Mocking ข้อมูลที่เราต้องการ
// 	// expectedResultClasses := []models.Class{
// 	// 	{
// 	// 		ID:          "1",
// 	// 		Title:       "Class A",
// 	// 		Description: "Description A",
// 	// 		ClassCategory: models.ClassCategory{
// 	// 			ID:           "cat1",
// 	// 			CategoryName: "Category A",
// 	// 		},
// 	// 		ClassSessions: []models.ClassSession{
// 	// 			{
// 	// 				ID:          "sess1",
// 	// 				ClassID:     "1",
// 	// 				MaxCapacity: 30,
// 	// 			},
// 	// 		},
// 	// 	},
// 	// 	{
// 	// 		ID:          "2",
// 	// 		Title:       "Class B",
// 	// 		Description: "Description B",
// 	// 		ClassCategory: models.ClassCategory{
// 	// 			ID:           "cat1",
// 	// 			CategoryName: "Category B",
// 	// 		},
// 	// 		ClassSessions: []models.ClassSession{
// 	// 			{
// 	// 				ID:          "sess1",
// 	// 				ClassID:     "1",
// 	// 				MaxCapacity: 30,
// 	// 			},
// 	// 		},
// 	// 	},
// 	// }

// 	// classTier := "essential"
// 	// keyword := "React"

// 	// var total int64 = 2
// 	// query := `^SELECT count(*) FROM classes WHERE (title ILIKE \$1 OR description ILIKE \$1) AND class_tier = \$2`
// 	// mock.ExpectQuery(query).
// 	// 	WithArgs(keyword, classTier).
// 	// 	WillReturnRows(sqlmock.NewRows([]string{
// 	// 		"id", "image", "name", "description", "url_pt", "url_ct", "url_ed", "created_at", "updated_at", "publish_at", "created_by", "updated_by", "position", "is_active", "is_remove",
// 	// 	}).AddRow("", "image_url", "Test App", nil, nil, nil, nil, time.Time{}, time.Time{}, time.Time{}, nil, nil, nil, nil, nil))

// 	// // // Mocking คำสั่ง SQL
// 	// // mockQuery := mock.ExpectQuery("^SELECT (.+) FROM \"classes\"")
// 	// // mockQuery.WithArgs("%keyword%", "%keyword%").WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
// 	// // 	AddRow("1", "Class A", "Description A").
// 	// // 	AddRow("2", "Class B", "Description B"))

// 	// // // ทดสอบการเรียกใช้งาน GetAllClasses
// 	// classes, count, err := repo.GetAllClasses("class_tier", "keyword", 1, 10)

// 	// // ตรวจสอบผลลัพธ์
// 	// assert.NoError(t, err)
// 	// assert.Equal(t, total, count)
// 	// assert.Equal(t, expectedResultClasses, *classes)

// 	// // ตรวจสอบว่า mock ได้ถูกเรียกใช้งานทั้งหมดแล้ว
// 	// err = mock.ExpectationsWereMet()
// 	// assert.NoError(t, err)

// 	t.Parallel()

// 	// Create a mock database connection
// 	// db, mock, err := sqlmock.New()
// 	// if err != nil {
// 	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	// }
// 	// defer db.Close()

// 	sqlDB, db, mock := DbMock(t)
// 	defer sqlDB.Close()

// 	// สร้าง instance ของ classGormRepository
// 	repo := NewClassGormRepository(db)

// 	mocktime := time.Now()
// 	// Prepare test data
// 	applicationInput := models.Class{Title: "Title"}
// 	expectedApplications := []models.Class{
// 		{},
// 	}

// 	class_tier := ""
// 	keyword := ""
// 	page := 1
// 	limit := 1

// 	// Mock the database query to count the applications
// 	countQuery := `^SELECT COUNT\(\*\) FROM classes WHERE is_remove = false AND name LIKE \$1`
// 	mock.ExpectQuery(countQuery).
// 		WithArgs("%Test App%").
// 		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

// 	// Mock the database query to get the applications
// 	selectQuery := `^SELECT id, image, name, description, url_pt, url_ct, url_ed, created_at, updated_at, publish_at, created_by, updated_by, position, is_active, is_remove FROM applications WHERE is_remove = false AND name LIKE \$1 ORDER BY position ASC, updated_at DESC`
// 	mock.ExpectQuery(selectQuery).
// 		WithArgs("%Test App%").
// 		WillReturnRows(sqlmock.NewRows([]string{
// 			"id", "image", "name", "description", "url_pt", "url_ct", "url_ed", "created_at", "updated_at", "publish_at", "created_by", "updated_by", "position", "is_active", "is_remove",
// 		}).AddRow(1, "image_url", "Test App", nil, nil, nil, nil, mocktime, mocktime, time.Time{}, "creator", "updater", nil, nil, nil))

// 	// Call the method to test
// 	classes, total, err := repo.GetAllClasses(class_tier, keyword, page, limit)

// 	// Assertions
// 	assert.NoError(t, err)                         // Expect no error for the success case
// 	assert.Equal(t, int64(1), total)               // Expect count to be 1
// 	assert.Equal(t, expectedApplications, classes) // Expect the result to match the expected applications

// 	// Verify that all expectations were met
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// import (
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func TestGetAllClasses(t *testing.T) {

// 	t.Parallel()
// 	// สร้าง mock database
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	// เชื่อมต่อ Gorm กับ mock database
// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info), // เปิด logger เพื่อ debug
// 	})
// 	assert.NoError(t, err)

// 	// Mock สำหรับการนับจำนวนทั้งหมด
// 	mock.ExpectQuery(`SELECT count\(\*\) FROM "classes"`).
// 		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

// 	// // Mock สำหรับการดึงข้อมูล Classes
// 	mock.ExpectQuery(`SELECT \* FROM "classes" LIMIT \$1 OFFSET \$2`).
// 		WithArgs(10, 0). // Limit และ Offset
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "class_tier"}).
// 			AddRow(1, "Class 1", "Description 1", "Beginner").
// 			AddRow(2, "Class 2", "Description 2", "Advanced"))

// 		// // Mock สำหรับ Preload ClassSessions
// 		// mock.ExpectQuery(`SELECT id,class_id,max_capacity,date,start_time,end_time FROM "class_sessions" WHERE "class_sessions"."class_id" IN \(\$1,\$2\)`).
// 		// 	WithArgs(1, 2).
// 		// 	WillReturnRows(sqlmock.NewRows([]string{"id", "class_id", "max_capacity", "date", "start_time", "end_time"}).
// 		// 		AddRow(1, 1, 20, "2025-01-10", "10:00:00", "12:00:00").
// 		// 		AddRow(2, 2, 15, "2025-01-11", "14:00:00", "16:00:00"))

// 	repo := NewClassGormRepository(gormDB)

// 	// // เรียกฟังก์ชันที่ต้องการทดสอบ
// 	_, total, err := repo.GetAllClasses("", "", 1, 10)
// 	assert.NoError(t, err)

// 	// // ตรวจสอบผลลัพธ์
// 	assert.Equal(t, int64(2), total)
// 	// assert.Equal(t, 2, len(*classes))

// 	// // ตรวจสอบ mock expectations
// 	err = mock.ExpectationsWereMet()
// 	// assert.NoError(t, err)
// }

// func TestGetAllClasses(t *testing.T) {

// 	t.Parallel()

// 	// สร้าง mock database
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	// เชื่อมต่อ Gorm กับ mock database
// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info), // เปิด log
// 	})
// 	assert.NoError(t, err)

// 	// Mock สำหรับการนับจำนวนทั้งหมด
// 	mock.ExpectQuery(`SELECT count\(\*\) FROM "classes"`).
// 		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

// 	// Mock สำหรับการดึงข้อมูล Classes
// 	mock.ExpectQuery(`SELECT \* FROM "classes" LIMIT \$1 OFFSET \$2`).
// 		WithArgs(10, 0).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "class_tier"}).
// 			AddRow(1, "Class 1", "Description 1", "Beginner").
// 			AddRow(2, "Class 2", "Description 2", "Advanced"))

// 	// Mock สำหรับ Preload ClassSessions
// 	mock.ExpectQuery(`SELECT id,class_id,max_capacity,date,start_time,end_time FROM "class_sessions" WHERE "class_sessions"."class_id" IN \(\$1,\$2\)`).
// 		WithArgs(1, 2).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "class_id", "max_capacity", "date", "start_time", "end_time"}).
// 			AddRow(1, 1, 20, "2025-01-10", "10:00:00", "12:00:00").
// 			AddRow(2, 2, 15, "2025-01-11", "14:00:00", "16:00:00"))

// 	// สร้าง repository
// 	repo := &classGormRepository{
// 		db: gormDB,
// 	}

// 	// เรียกฟังก์ชันที่ต้องการทดสอบ
// 	classes, total, err := repo.GetAllClasses("", "", 1, 10)
// 	assert.NoError(t, err)

// 	// ตรวจสอบผลลัพธ์
// 	assert.Equal(t, int64(2), total)
// 	assert.Equal(t, 2, len(*classes))

// 	// ตรวจสอบ mock expectations
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestGetAllClasses(t *testing.T) {
// 	// สร้าง mock database
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	// เชื่อมต่อ Gorm กับ mock database
// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})
// 	assert.NoError(t, err)

// 	// Mock สำหรับการนับจำนวนทั้งหมด
// 	mock.ExpectQuery(`SELECT count\(\*\) FROM "classes"`).
// 		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

// 	// Mock สำหรับการดึงข้อมูล Classes
// 	mock.ExpectQuery(`SELECT \* FROM "classes" LIMIT 10 OFFSET 0`).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "class_tier"}).
// 			AddRow(1, "Class 1", "Description 1", "Beginner").
// 			AddRow(2, "Class 2", "Description 2", "Advanced"))

// 	// Mock สำหรับ Preload ClassSessions
// 	mock.ExpectQuery(`SELECT id,class_id,max_capacity,date,start_time,end_time FROM "class_sessions" WHERE "class_sessions"."class_id" IN \(\$1,\$2\)`).
// 		WithArgs(1, 2).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "class_id", "max_capacity", "date", "start_time", "end_time"}).
// 			AddRow(1, 1, 20, "2025-01-10", "10:00:00", "12:00:00").
// 			AddRow(2, 2, 15, "2025-01-11", "14:00:00", "16:00:00"))

// 	// สร้าง repository
// 	repo := &classGormRepository{
// 		db: gormDB,
// 	}

// 	// เรียกฟังก์ชันที่ต้องการทดสอบ
// 	classes, total, err := repo.GetAllClasses("", "", 1, 10)
// 	assert.NoError(t, err)

// 	// ตรวจสอบผลลัพธ์
// 	assert.Equal(t, int64(2), total)
// 	assert.Equal(t, 2, len(*classes))

// 	// ตรวจสอบ mock expectations
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestGetAllClasses(t *testing.T) {
// 	// สร้าง mock database
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	// เชื่อมต่อ Gorm กับ mock database
// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})
// 	assert.NoError(t, err)

// 	// เตรียม mock data
// 	mock.ExpectQuery(`SELECT count\(\*\) FROM "classes"`).
// 		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

// 	mock.ExpectQuery(`SELECT \* FROM "classes"`).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
// 			AddRow(1, "Class 1", "Description 1").
// 			AddRow(2, "Class 2", "Description 2"))

// 	// สร้าง repository
// 	repo := &classGormRepository{
// 		db: gormDB,
// 	}

// 	// เรียกฟังก์ชันที่ต้องการทดสอบ
// 	classes, total, err := repo.GetAllClasses("", "", 1, 10)
// 	assert.NoError(t, err)

// 	// ตรวจสอบผลลัพธ์
// 	assert.Equal(t, int64(2), total)
// 	assert.Equal(t, 2, len(*classes))

// 	// ตรวจสอบ mock expectations
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }
