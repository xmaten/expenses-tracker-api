package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Expense struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"text;not null;" json:"title"`
	Category  string    `gorm:"size:255;not null" json:"category"`
	Value     uint64    `gorm:"not null" json:"value"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	Date time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (e *Expense) Prepare() {
	e.Title = html.EscapeString(strings.TrimSpace(e.Title))
	e.Category = html.EscapeString(strings.TrimSpace(e.Category))
	e.Author = User{}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
}

func (e *Expense) Validate() map[string]string {
	var err error
	var errorMessages = make(map[string]string)

	if e.Title == "" {
		err = errors.New("required title")
		errorMessages["Required_title"] = err.Error()
	}

	if e.Category == "" {
		err = errors.New("required category")
		errorMessages["Required_category"] = err.Error()
	}

	if e.Value <= 0 {
		err = errors.New("required value")
		errorMessages["Required_value"] = err.Error()
	}

	if e.AuthorID < 1 {
		err = errors.New("required Author")
		errorMessages["Required_author"] = err.Error()
	}

	return errorMessages
}

func (e *Expense) SaveExpense(db *gorm.DB) (*Expense, error) {
	var err error
	err = db.Debug().Model(&Expense{}).Create(&e).Error

	if err != nil {
		return &Expense{}, err
	}

	if e.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", e.AuthorID).Take(&e.Author).Error

		if err != nil {
			return &Expense{}, err
		}
	}

	return e, nil
}

func (e *Expense) FindUserExpenses(db *gorm.DB, uid uint32, dateStart string, dateEnd string) (*[]Expense, error) {
	var err error

	expenses := []Expense{}

	if dateStart != "" && dateEnd != "" {
		err = db.Debug().Model(&Expense{}).Where("author_id = ?", uid).Where("date >= ?", dateStart).Where("date <= ?", dateEnd).Limit(100).Order("created_at desc").Find(&expenses).Error
	} else {
		err = db.Debug().Model(&Expense{}).Where("author_id = ?", uid).Limit(100).Order("created_at desc").Find(&expenses).Error
	}

	if err != nil {
		return &[]Expense{}, err
	}

	if len(expenses) > 0 {
		for i, _ := range expenses {
			err := db.Debug().Model(&User{}).Where("id = ?", expenses[i].AuthorID).Take(&expenses[i].Author).Error

			if err != nil {
				return &[]Expense{}, err
			}
		}
	}

	return &expenses, nil
}






















