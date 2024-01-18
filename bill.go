package main

import (
	"fmt"
	"os"
)

type bill struct {
	Name  string
	Items map[string]float64
	Tip   float64
}

func NewBill(name string) bill {
	b := bill{
		Name:  name,
		Items: map[string]float64{},
		Tip:   0,
	}
	return b
}

// format bill
func (b bill) format() string {
	fs := "Breakdown of bill: \n"
	var total float64 = 0
	for k, v := range b.Items {
		fs += fmt.Sprintf("%-25v...$%v \n", k+":",v)
		total += v
	}
	fs += fmt.Sprintf("%-25v...$%0.2f \n", "Tip:", b.Tip)

	fs += fmt.Sprintf("%-25v...$%0.2f", "Total:", total+b.Tip)

	return fs
}
// update tip
func (b *bill) UpdateTip(tip float64){
	b.Tip = tip
}
// add items
func (b *bill) AddItems(name string, price float64){
	b.Items[name] = price
}

//  save bill
func (b *bill) save(){
	data := []byte(b.format())
	err:= os.WriteFile("Bills/"+b.Name+".txt",data,0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("Bill was saved")
}