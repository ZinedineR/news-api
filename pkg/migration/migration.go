package migration

import (
	"fmt"
	"sort"

	newsDomain "news-api/internal/news/domain"
	userDomain "news-api/internal/user/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Initmigrate(db *gorm.DB) {
	// port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"))
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	// checkError(err)

	sqlDB, err := db.DB()
	if err != nil {
		defer sqlDB.Close()
	}

	executePendingMigrations(db)

	// Migrate rest of the models
	logrus.Info(fmt.Println("AutoMigrate Model [table_name]"))
	db.AutoMigrate(&userDomain.User{})
	logrus.Info(fmt.Println("  TableModel [" + (&userDomain.User{}).TableName() + "]"))
	db.AutoMigrate(&userDomain.Verification{})
	logrus.Info(fmt.Println("  TableModel [" + (&userDomain.Verification{}).TableName() + "]"))
	db.AutoMigrate(&newsDomain.Categories{})
	logrus.Info(fmt.Println("  TableModel [" + (&newsDomain.Categories{}).TableName() + "]"))
	db.Debug().AutoMigrate(&newsDomain.News{})
	logrus.Info(fmt.Println("  TableModel [" + (&newsDomain.News{}).TableName() + "]"))

	// db.AutoMigrate(&entity.Users{})
	// log.Info().Msg("  TableModel [" + (&entity.Users{}).TableName() + "]")

}

func executePendingMigrations(db *gorm.DB) {
	db.AutoMigrate(&MigrationHistoryModel{})
	lastMigration := MigrationHistoryModel{}
	skipMigration := db.Order("migration_id desc").Limit(1).Find(&lastMigration).RowsAffected > 0

	// skip to last migration
	keys := make([]string, 0, len(migrations))
	for k := range migrations {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// run all migrations in one transaction
	if len(migrations) == 0 {
		logrus.Infoln(fmt.Print("No pending migrations"))
	} else {
		db.Transaction(func(tx *gorm.DB) error {
			for _, k := range keys {
				if skipMigration {
					if k == lastMigration.MigrationID {
						skipMigration = false
					}
				} else {
					logrus.Infoln(fmt.Sprintf("  " + k))
					tx.Transaction(func(subTx *gorm.DB) error {
						// run migration update
						checkError(migrations[k](subTx))
						// insert migration id into history
						checkError(subTx.Create(MigrationHistoryModel{MigrationID: k}).Error)
						return nil
					})
				}
			}
			return nil
		})
	}
}

type mFunc func(tx *gorm.DB) error

var migrations = make(map[string]mFunc)

// MigrationHistoryModel model migration
type MigrationHistoryModel struct {
	MigrationID string `gorm:"type:text;primaryKey"`
}

// TableName name of migration table
func (model *MigrationHistoryModel) TableName() string {
	return "migration_history"
}

func checkError(err error) {
	if err != nil {
		logrus.Infoln(fmt.Print(err.Error()))
		panic(err)
	}
}

// func registerMigration(id string, fm mFunc) {
// 	migrations[id] = fm
// }
