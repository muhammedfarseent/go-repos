package main 
 
import "fmt"

func main(){

	menu:= map[string]float64{

		"masala dossa": 25.50,
		"porotta": 10.55,
		"beaf rost": 250,
	}

	fmt.Println(menu)
	fmt.Println(menu["beaf rost"])
}