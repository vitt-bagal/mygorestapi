package supplier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// create item(supplier1) struct
type Item struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

// create veg_item(supplier2) struct
type Veg_Item struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	Price       string `json:"price"`
}

// create veg_item(supplier2) struct
type Grain_Item struct {
	ItemId   string `json:"itemId"`
	ItemName string `json:"itemName"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

// Function to return list of items from given supplier
func CallSupplier(apiEnv string) []Item {
	var resp []Item
	var veg []Veg_Item
	var grain []Grain_Item
	apiurl := os.Getenv(apiEnv)
	// Consume Rest api created by supplier
	req, err := http.Get(apiurl)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if apiEnv == "VEG_SUPPLIER" {
		fmt.Println("Called Veg supplier....")
		json.Unmarshal(body, &veg)
		n := len(veg)
		r1 := make([]Item, n)
		for i, v := range veg {
			r1[i].Id = v.ProductId
			r1[i].Name = v.ProductName
			r1[i].Quantity = v.Quantity
			r1[i].Price = v.Price
		}
		resp = append(resp, r1...)
	} else if apiEnv == "GRAIN_SUPPLIER" {
		fmt.Println("Called Grain supplier....")
		json.Unmarshal(body, &grain)
		n := len(grain)
		r1 := make([]Item, n)
		for i, v := range grain {
			r1[i].Id = v.ItemId
			r1[i].Name = v.ItemName
			r1[i].Quantity = v.Quantity
			r1[i].Price = v.Price
		}
		resp = append(resp, r1...)
	} else {
		fmt.Println("Called Fruit supplier....")
		json.Unmarshal(body, &resp)
	}

	return resp
}
