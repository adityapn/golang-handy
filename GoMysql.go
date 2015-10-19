package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Sku struct{		
	mrp float32
	name string	
}

func main() {
	db, err := sql.Open("mysql", "root:@/ofd_cms")
	db.SetMaxIdleConns(10)

	if err != nil{
		fmt.Println("can not connect to mysql: ",err)
		return
	}

	rows, error := db.Query("select mrp,name from skus limit 10")
	if error != nil{
		fmt.Println("Error while running query ",error)
		return
	}

	skus := []Sku{}
	for rows.Next(){
		var sku Sku
		skuError := rows.Scan(&sku.mrp,&sku.name)
		if skuError != nil{
			fmt.Println("can not parse sku error ",skuError)
			return
		}		
		skus = append(skus,sku)
	}
	for i :=0 ; i< len(skus) ; i++{
		fmt.Println(skus[i].name)
	}
}
