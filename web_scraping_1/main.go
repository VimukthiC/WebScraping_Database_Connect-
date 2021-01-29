package main

import (
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3307)/car")

	var location string
	fmt.Print("Enter Location : ")
	fmt.Scan(&location)

	var category string
	fmt.Print("Enter Catagory : ")
	fmt.Scan(&category)

	c := colly.NewCollector(
		colly.AllowedDomains("ikman.lk"),
	)

	c.OnHTML(".gtm-normal-ad", func(e *colly.HTMLElement) {

		var title string = e.ChildText(".heading--2eONR")
		var price string = e.ChildText(".price--3SnqI")

		fmt.Println(title)
		fmt.Println(price)

		insert, err := db.Query("INSERT INTO carDetail(title,price,location,category) VALUES (?,?,?,?)", title, price, location, category)

		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	})

	c.Visit("https://ikman.lk/en/ads/" + location + "/" + category)

}
