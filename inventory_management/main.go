package main



type stockupdater interface{
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

func(i * inventory)update(a int ){
     i.stock=a

} 


func main() {

}
