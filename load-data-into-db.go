package main

import (
	"gorm.io/gorm"
  "gorm.io/driver/postgres"
	"log"
	"os"
	"bufio"
	"strings"
)

type Account struct {
	Telephone string
	CompanyNumber string
	Firstname string
	Lastname string
	Sex string
	City string
  Address string
  Attr1 string
  Attr2 string
  Attr3 string
  Attr4 string
  Attr5 string
  Attr6 string
}

func main() {	
	// Connect to the database
	dsn := "host=localhost user=aiu password=aiu dbname=aiu port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Account{})

	// open file  ./Italy/0.csv
	var path = "./Italy/4.txt"
	var file *os.File
	var line string
	var err1 error
	var i int
	var account Account
	var accounts []Account

	// open file
	file, err1 = os.Open(path)
	if err1 != nil {
		log.Fatal(err1)
	}

	log.Println("prepare data")

	// read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if i == 0 {
			i++
			continue
		}

		// split line
		l := strings.Split(line, ":")

		// loop through line
		for j, v := range l {
			switch j {
			case 0:
				account.Telephone = v
			case 1:
				account.CompanyNumber = v
			case 2:
				account.Firstname = v
			case 3:
				account.Lastname = v
			case 4:
				account.Sex = v
			case 5:
				account.City = v
			case 6:
				account.Address = v
			case 7:
				account.Attr1 = v
			case 8:
				account.Attr2 = v
			case 9:
				account.Attr3 = v
			case 10:
				account.Attr4 = v
			case 11:
				account.Attr5 = v
			case 12:
				account.Attr6 = v
			}
		}

		accounts = append(accounts, account)
		// prignt accounts length
		//log.Println(len(accounts), "accounts")
	}
	
	// close file
	err1 = file.Close()
	if err1 != nil {
		log.Fatal(err1)
	}

	log.Println("Inserting data into database...")
	db.CreateInBatches(accounts, 200)
	log.Println("Done.")
}


