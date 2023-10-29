package model

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Title    string    `json:"title" gorm:"uniqueIndex"`
	Phone    string    `json:"phone"  gorm:"uniqueIndex"`
	CreditTo time.Time `json:"credit_to"`

	Password string `json:"password"`

	// each user can has many team leads
	Users []User `json:"users"`
}

// getGuard implements utility.authenticatable.
func (Company) GetGuard() string {
	return "admin-user"
}

// getID implements utility.authenticatable.
func (company Company) GetID() uint {
	return company.ID
}



//==============================================================================//
//																			    //
//                                  Setters										//
//																			    //
//==============================================================================//

func (company *Company) SetPassword() {

	hash, err := bcrypt.GenerateFromPassword([]byte(company.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err.Error())
	}


	company.Password = string(hash)
}
