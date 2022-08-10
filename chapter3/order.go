package main

import (
	"time"
	"fmt"
	"errors"
)

type Order struct {
	id         int
	orderLines []OrderLine
	totalCost  float64
	created    time.Time
	user       string
}

func NewOrder(id int,user string) (*Order,error) {
	if(len(user)==0){
		return nil, errors.New("user cannot be empty.")
	}
	if(id<0){
		return nil, errors.New("id should be greater than 0.")
	}
	order := Order{}
	order.id = id
	order.created = time.Now()
	order.user = user
	return &order, nil
}

func (o *Order) PrintTotal(){
	if (len(o.orderLines)==0){
		fmt.Print("0")
	} else{
		total:=0.0
		for _,orderLine := range o.orderLines {
			total+=orderLine.LineTotal()
		}
		fmt.Println(total)
	}
}

func (o *Order) AddOrderLine(id int,unitPrice float64,quantity int,item string)(*OrderLine,error){
	newOrderLine,err := NewOrderLine(id,unitPrice,quantity,item)
	if (err!= nil){
		return nil,err
	}
	o.orderLines=append(o.orderLines,(*newOrderLine))
	return newOrderLine,nil
}

func (o *Order) UpdateOrderLine(id int,unitPrice float64,quantity int,item string)(bool){
	updated := false
	for i,orderLine := range o.orderLines {
		if(orderLine.id == id){
			updated = true
			orderLine.unitPrice = unitPrice
			orderLine.quantity = quantity
			orderLine.item = item
			o.orderLines[i]=orderLine
			break;
		}
	}
	return updated
}

func (o *Order) RemoveOrderLine(i int){
	if(i < 0 || i>len(o.orderLines)) {
		return
	}
    o.orderLines[i] = o.orderLines[len(o.orderLines)-1]
    o.orderLines = o.orderLines[:len(o.orderLines)-1]
}