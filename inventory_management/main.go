package main

import "fmt"

type stockupdater interface {
	update(int)
}

type product struct {
	name  string
	price float32
	stock int
}

type inventory struct {
	product
	category string
}

func (i *inventory) update(a int) {
	i.stock = a

}

var inventoryMap = make(map[string]*inventory)
var orderHistory = []string{}

func addProduct(name string, price float32, stock int, category string) {

	if name=="" {
		fmt.Println("name empty")
		return
	}
	if price <=0{
			fmt.Println("price invalid")
			return 
	}
	inventoryMap[name] = &inventory{
		product: product{
			name:  name,
			price: price,
			stock: stock},
		category: category,
	}

}

func update(name string, newstock int) {
	if newstock <=0 {
		fmt.Println("negative stock")
		return 


	}
	if val, found := inventoryMap[name]; found {
		val.update(newstock)
	}else{
		fmt.Println("product not found")
	}
}
func placeorder(name string) {
	orderHistory = append(orderHistory, name)
}

func printInventory() {
	for _, v := range inventoryMap {
		fmt.Println(v)
	}

}
func inventoryStatus()(int,int){
   totalstock:=0
   for _,v:=range inventoryMap{
	totalstock+=v.stock


   }
   return totalstock,len(inventoryMap)


}
func analysorder(){
   ordercount:=make(map[string]int)
   for _,v:=range orderHistory{
	       ordercount[v]++;
   }

   fmt.Println("\n order naalytics")

   for k,v:=range ordercount{
      fmt.Println(k,"->",v)
   }



}

func categoryrouter(pname string){
   if inv,found:=inventoryMap[pname]; found{
	  switch inv.category{
	  case "Electronics":
                       fmt.Println(pname,"belongs to the product name")
        case "Gadgets":
            fmt.Println(pname, "belongs to Gadgets section")
        default:
            fmt.Println(pname, "belongs to Other section")
        }


	  }else{
                      fmt.Println("Product not found, cannot route category")

  } 


   }




func main() {
    var choice int

    for {   
        fmt.Println("\n--- MENU ---")
        fmt.Println("1. Add Product")
        fmt.Println("2. Update Stock")
        fmt.Println("3. Place Order")
        fmt.Println("4. View Inventory")
        fmt.Println("5. Analyze Orders")
        fmt.Println("6. View Stats")
        fmt.Println("7. Category Router")
        fmt.Println("0. Exit")
        fmt.Print("Enter choice: ")

        fmt.Scan(&choice)   
        switch choice {   
        case 1:
            var name, category string
            var price float32
            var stock int

            fmt.Print("Product name: ")
            fmt.Scan(&name)
            fmt.Print("Price: ")
            fmt.Scan(&price)
            fmt.Print("Stock: ")
            fmt.Scan(&stock)
            fmt.Print("Category: ")
            fmt.Scan(&category)

            addProduct(name, price, stock, category)   

        case 2:
            var name string
            var stock int
            fmt.Print("Product name to update: ")
            fmt.Scan(&name)
            fmt.Print("New stock: ")
            fmt.Scan(&stock)

            update(name, stock) 

        case 3:
            var name string
            fmt.Print("Product name to order: ")
            fmt.Scan(&name)
            placeorder(name)  
        case 4:
            printInventory() 
        case 5:
            analysorder() 

        case 6:
            totalStock, totalProducts := inventoryStatus() 
            fmt.Println("Total stock:", totalStock)
            fmt.Println("Different products:", totalProducts)

        case 7:
            var name string
            fmt.Print("Product to route: ")
            fmt.Scan(&name)
            categoryrouter(name) 
        case 0:
            fmt.Println("Exiting program... Bye!")
            break 

        default:
            fmt.Println("Invalid choice, try again!")
        }

       
    }
}
