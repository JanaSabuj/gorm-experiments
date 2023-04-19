package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	dsn := "root:root@tcp(127.0.0.1:3308)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D41", Price: 100})
	db.Create(&Product{Code: "D4a", Price: 100})
	db.Create(&Product{Code: "D4b", Price: 100})
	db.Create(&Product{Code: "D4c", Price: 100})
	db.Create(&Product{Code: "D4d", Price: 100})
	db.Create(&Product{Code: "D42", Price: 200})

	db.Create(&Product{Code: "D43", Price: 200})
	db.Create(&Product{Code: "D44", Price: 200})
	db.Create(&Product{Code: "D45", Price: 200})

	db.Create(&Product{Code: "D46", Price: 200})
	db.Create(&Product{Code: "D47", Price: 200})

	// select id from products where price = 200 order by id desc limit 2
	// delete from products where price = 200 AND id NOT IN (select id from products where price = 200 order by id desc limit 2)
	// subquery := db.Model(&Product{}).Select("id").Where("price = ?", 200).Order("id DESC").Limit(2)
	// subquery := db.Select("id").Table("(SELECT id FROM products WHERE price = ? ORDER BY id DESC LIMIT 3) t", 200)

	// err = db.Model(&Product{}).Where("price = ? AND id NOT IN (?)", 200, subquery).Delete(&Product{}).Error
	// err = db.Where("price = ?", 200).Not("id IN (?)", subquery).Delete(&Product{}).Error

	price := 200
	db.Exec("DELETE FROM products WHERE price = ? AND id NOT IN (SELECT id FROM (SELECT id FROM products WHERE price = ? ORDER BY id DESC LIMIT 3) t)", price, price)

	if err != nil {
		fmt.Println("Error deleting...")
	}

	// // Read
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)

	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
