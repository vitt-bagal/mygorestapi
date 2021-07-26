package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/vitt-bagal/mygorestapi/handler/supplier"
)

// List of predefined suppliers
var envSupplier = []string{"FRUIT_SUPPLIER", "VEG_SUPPLIER", "GRAIN_SUPPLIER"}

// Handler function to buy-item-qty endpoint
func buyItemQuantity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("Called buyItemQuantity API...")
	var responseObject, result []supplier.Item
	var foundKey = false
	qty, _ := strconv.ParseInt(params["quantity"], 10, 64)

	for _, env := range envSupplier {
		responseObject = supplier.CallSupplier(env)
		for i, val := range responseObject {
			if strings.EqualFold(responseObject[i].Name, params["name"]) && responseObject[i].Quantity >= int(qty) {
				fmt.Printf("Product value is %v\n", val)
				foundKey = true
				result = append(result, val)
				p := &val
				p.Quantity = p.Quantity - int(qty)
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if !foundKey {
		json.NewEncoder(w).Encode("NOT_FOUND")
		return
	}
	json.NewEncoder(w).Encode(result)

}

// Handler function to buy-item-qty-price endpoint
func buyItemQuantityPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("Called buyItemQuantity API...")
	var responseObject, result []supplier.Item
	var foundKey = false
	qty, _ := strconv.ParseInt(params["quantity"], 10, 64)
	buyPrice, _ := strconv.ParseFloat(params["price"], 10)

	for _, env := range envSupplier {
		responseObject = supplier.CallSupplier(env)
		for i, val := range responseObject {
			sellPrice, _ := strconv.ParseFloat(strings.Split(responseObject[i].Price, "$")[1], 64)
			fmt.Printf("Sell price %v buy prce %v", sellPrice, buyPrice)
			if strings.EqualFold(responseObject[i].Name, params["name"]) && responseObject[i].Quantity >= int(qty) && sellPrice <= buyPrice {
				// Calculate total buying price
				fmt.Printf("Product value is %v\n", val)
				foundKey = true
				result = append(result, val)
				p := &val
				p.Quantity = p.Quantity - int(qty)
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if !foundKey {
		json.NewEncoder(w).Encode("NOT_FOUND")
		return
	}
	json.NewEncoder(w).Encode(result)

}

// Handler function to buy-item endpoint
func buyItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("Called buyitem API...")
	var responseObject, result []supplier.Item
	var foundKey = false
	for _, env := range envSupplier {
		responseObject = supplier.CallSupplier(env)
		for i, val := range responseObject {
			if strings.EqualFold(responseObject[i].Name, params["name"]) {
				fmt.Printf("Product value is %v\n", val)
				foundKey = true
				result = append(result, val)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if !foundKey {
		json.NewEncoder(w).Encode("NOT_FOUND")
		return
	}
	json.NewEncoder(w).Encode(result)
}

// Handler function to fast-buy-item endpoint
func fastBuyItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("Called fastbuyitem API...")
	var resp []supplier.Item
	var veg []supplier.Veg_Item
	var grain []supplier.Grain_Item
	//var wg sync.WaitGroup
	//wg.Add(len(envSupplier))
	var foundKey = false
	w.Header().Set("Content-Type", "application/json")
	for _, env := range envSupplier {
		go func(apiEnv string) {
			//defer wg.Done()
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
				r1 := make([]supplier.Item, n)
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
				r1 := make([]supplier.Item, n)
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
			for i, val := range resp {
				if strings.EqualFold(resp[i].Name, params["name"]) {
					fmt.Printf("Product value is %v\n", val)
					foundKey = true
					json.NewEncoder(w).Encode(val)
					return
				}
			}

			//wg.Wait()
		}(env)
		time.Sleep(100 * time.Millisecond)
		//wg.Wait()
	}
	if !foundKey {
		json.NewEncoder(w).Encode("NOT_FOUND")
		return
	}
}

// Home handler function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Food Aggregator....")
}

// Handler function  to show-summary
func showsummary(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called showsummary API...")
	var responseObject, result []supplier.Item

	for _, env := range envSupplier {
		responseObject = supplier.CallSupplier(env)
		result = append(result, responseObject...)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
