package model

import (
	"fmt"
	"time"
)

func Test()  {
	vendor := GetVendorByID(75)
	fmt.Println(vendor)

	vendor.VendorName="上海斯年信息技术有限公司"

	VendorSave(vendor)
	vendors := GetVendors()
	fmt.Println(len(vendors))
}

func GetVendorByID(id int64) *Vendor {
	vendor := Vendor{}
	if err := DB().Where(VendorId + " = ?",75).First(&vendor).Error; err != nil {
		fmt.Println(err)
		return nil;
	}

	return &vendor
}
func VendorSave(vendor *Vendor)  {
	if err := DB().Save(&vendor).Error; err != nil {
		fmt.Println(err)
	}
}
func GetVendors() []Vendor {
	vendors := []Vendor{}
	if err := DB().Find(&vendors).Error; err != nil {
		fmt.Println(err)
	}

	return vendors
}

func (Vendor) TableName() string {
	return "TVendor"
}

//定义数据模型
type Vendor struct {
	VendorId        int32 `gorm:"primary_key;column:VendorId"`
	VendorCode     string `gorm:"column:VendorCode"`
	VendorName   string `gorm:"column:VendorName"`
	CreateTime time.Time `gorm:"column:CreateTime"`
	UpdateTime time.Time `gorm:"column:UpdateTime"`
	Status   string `gorm:"column:Status"`
}

//定义数据库字段常量
const (
	VendorId = "VendorId"
	VendorCode = "VendorCode"
	VendorName = "VendorName"
	CreateTime = "CreateTime"
	UpdateTime = "UpdateTime"
	Status = "Status"
)

//定义枚举类型常量
const (
	Status_A string = "A"
	Status_P string = "P"
)