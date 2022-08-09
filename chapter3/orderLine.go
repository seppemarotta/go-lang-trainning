package main

import (
	"time"
	"errors"
)

type OrderLine struct {
	id           int
	unitPrice   float64
	quantity     int
	item         string
	created  time.Time
}

func NewOrderLine(id int,unitPrice float64,quantity int,item string) (*OrderLine,error){
	if(unitPrice<0){
		return nil, errors.New("age should be greater than 0.")
	}
	if(quantity<0){
		return nil, errors.New("quantity should be greater than 0.")
	}
	if(id<0){
		return nil, errors.New("id should be greater than 0.")
	}
	return &OrderLine{id:id, unitPrice:unitPrice,quantity:quantity,item:item,created: time.Now() }, nil
}

func (orderLine *OrderLine) LineTotal() (float64){
	return (float64(orderLine.quantity) * orderLine.unitPrice)
}