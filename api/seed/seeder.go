package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/xmaten/expenses-tracker-api/api/models"
	"log"
)

var users = []models.User{
	models.User{
		Username: "mateusz",
		Email:    "mateusz@test.com",
		Password: "password",
	},
	models.User{
		Username: "sylwia",
		Email:    "sylwia@test.com",
		Password: "password123",
	},
}

var incomes = []models.Income{
	models.Income{
		Title: "FVAT 1/1/20",
		Value: 5000,
	},
	models.Income{
		Title: "Income2",
		Value: 1000,
	},
}

var expenses = []models.Expense{
	models.Expense{
		Title: "TV",
		Category: "electronics",
		Value: 3000,
	},
	models.Expense{
		Title: "Food",
		Category: "food",
		Value: 300,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Expense{}, &models.Income{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Income{}, &models.Expense{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Expense{}).AddForeignKey("author_id", "uders(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Income{}).AddForeignKey("author_id", "uders(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		expenses[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Expense{}).Create(&expenses[i]).Error
		if err != nil {
			log.Fatalf("cannot seed expenses table: %v", err)
		}

		incomes[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Income{}).Create(&incomes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed incomes table: %v", err)
		}

	}
}

























