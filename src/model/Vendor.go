package model

import (
	"fmt"
	"time"
)

func Test()  {
	user := User{}
	if err := DB().First(&user).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(user.LastName)

	customer := Customer{}
	if err := DB().Where("CustomerID=?",1).Find(&customer).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(customer.CustomerCode)
}

func (Vendor) TableName() string {
	return "TVendor"
}
func (Customer) TableName() string {
	return "TCustomerTest"
}
func (User) TableName() string {
	return "users"
}

type Vendor struct {
	VendorId        int32 `gorm:"column:VendorId"`
	VendorCode     string `gorm:"column:VendorCode"`
	VendorName   string `gorm:"column:VendorName"`
	CreateTime time.Time `gorm:"column:CreateTime"`
	UpdateTime time.Time `gorm:"column:UpdateTime"`
}

type Customer struct {
	ID        int32 `gorm:"primary_key;column:CustomerID"`
	CustomerCode     string `gorm:"column:CustomerCode"`
}

type User struct {
	ID        uint `gorm:"primary_key"`
	FirstName string `gorm:"column:first_name"`
	LastName string `gorm:"column:last_name"`
}