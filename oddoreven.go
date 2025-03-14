
//Write a Go program that asks the user for a number and checks if it's even or odd
package main
 
import "fmt"

func main(){
     var number int 
	fmt.Printf("enter any number to chek wether it is odd or even : ")
	fmt.Scanf("%d",&number)
// check the number weather even or odd 
   if number%2==0{
	fmt.Printf("%d number is even \n",number)
   }else{
	fmt.Printf("%d number is odd \n",number)
		}

   } 


