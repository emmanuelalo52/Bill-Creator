package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getInput(prompt string, r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	input,err := r.ReadString('\n')
	return strings.TrimSpace(input),err
}
func CreateBill() bill{
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name,_ := getInput("Create a new bill: ",reader)
	b := NewBill(name)
	fmt.Println("Created the bill: ",b.Name)

	return b
}
func PromtOptions(b bill){
	reader := bufio.NewReader(os.Stdin)
	opt,_ := getInput("Choose actions (A - Add Items, S - Save Bill, T - Add Tip): ",reader)
	switch opt{
	case "A":
		name,_ := getInput("Item Names:", reader)
		price,_ := getInput("Item Prices:", reader)
		p,err := strconv.ParseFloat(price,64)
		if err != nil{
			fmt.Println("The price is not a float")
			PromtOptions(b)
		}
		b.AddItems(name,p)

		fmt.Println("Items added: ",name,price)
		PromtOptions(b)
	case "T":
		tip,_ := getInput("Tip ($):",reader)
		t,err := strconv.ParseFloat(tip,64)
		if err != nil{
			fmt.Println("The tip is not a float")
			PromtOptions(b)
		}
		b.UpdateTip(t)

		fmt.Println("Tip added:",tip)
		PromtOptions(b)
	case "S":
		b.save()
		fmt.Println("File saved:", b.Name)
	default:
		fmt.Println("Invalid Choice")
		PromtOptions(b)
	}
}
func main(){
	

	mybill := CreateBill()
	PromtOptions(mybill)

}