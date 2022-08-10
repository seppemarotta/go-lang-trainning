package main

func main() {
	order,_ := NewOrder(1,"Giuseppe")
	order.AddOrderLine(1,3.5,2,"Item1") // 7
	order.AddOrderLine(2,1,2,"Item2") // 2

	order.PrintTotal() // 9 

	order.UpdateOrderLine(1,3.5,3,"Item1") // 10.5
	order.PrintTotal() //12.5

	order.RemoveOrderLine(0) // -10.5
	order.PrintTotal() // 2
}
