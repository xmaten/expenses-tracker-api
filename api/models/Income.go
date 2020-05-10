package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Income struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"text;not null;" json:"title"`
	Value     uint64    `gorm:"not null" json:"value"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	Date string `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (e *Income) Prepare() {
	e.Title = html.EscapeString(strings.TrimSpace(e.Title))
	e.Author = User{}
	e.Date = html.EscapeString(strings.TrimSpace(e.Date))
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
}

func (e *Income) Validate() map[string]string {
	var err error
	var errorMessages = make(map[string]string)

	if e.Title == "" {
		err = errors.New("required title")
		errorMessages["Required_title"] = err.Error()
	}

	if e.Value <= 0 {
		err = errors.New("required value")
		errorMessages["Required_value"] = err.Error()
	}

	if e.Date == "" {
		err = errors.New("required date")
		errorMessages["Required_date"] = err.Error()
	}

	if e.AuthorID < 1 {
		err = errors.New("required Author")
		errorMessages["Required_author"] = err.Error()
	}

	return errorMessages
}

func (e *Income) SaveIncome(db *gorm.DB) (*Income, error) {
	var err error
	err = db.Debug().Model(&Income{}).Create(&e).Error

	if err != nil {
		return &Income{}, err
	}

	if e.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", e.AuthorID).Take(&e.Author).Error

		if err != nil {
			return &Income{}, err
		}
	}

	return e, nil
}

func (e *Income) FindUserIncomes(db *gorm.DB, uid uint32) (*[]Income, error) {
	var err error
	incomes := []Income{}
	err = db.Debug().Model(&Income{}).Where("author_id = ?", uid).Limit(100).Order("created_at desc").Find(&incomes).Error

	if err != nil {
		return &[]Income{}, err
	}

	if len(incomes) > 0 {
		for i, _ := range incomes {
			err := db.Debug().Model(&User{}).Where("id = ?", incomes[i].AuthorID).Take(&incomes[i].Author).Error

			if err != nil {
				return &[]Income{}, err
			}
		}
	}

	return &incomes, nil
}






















