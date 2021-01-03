package main

import (
	"testing"
	"time"
)

func TestCartInit(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}

	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
}

func TestCartIsExpired(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 5
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}
	if cart.IsExpired() == true {
		t.Errorf("cart expired on init")
	}
	// wait until cart expires
	time.Sleep(10 * time.Second)
	// check cart is expired
	if cart.IsExpired() == false {
		t.Errorf("cart expired failed")
	}
}

func TestCartSerialize(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}
	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	if cart.Serialized != "" {
		t.Errorf("cart serialized non empty on init")
	}
	// add serialized
	cart.Serialize()
	// check if serialized generated
	if cart.Serialized == "" {
		t.Errorf("cart serialize failed.")
	}

}

func TestCartAddOrder(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}

	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	// init slack routing with cart
	order := Order{Person: "test order"}
	order.AddItem("test item")
	_, _ = cart.AddOrder(order)

	if len(cart.Orders) == 0 {
		t.Errorf("cart add item failed")
	}
}

func TestCartGetOrder(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}

	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	// init slack routing with cart
	order := Order{Person: "test order"}
	order.AddItem("test item")
	_, _ = cart.AddOrder(order)

	if len(cart.Orders) == 0 {
		t.Errorf("cart add item failed")
	}
	newOrder, err := cart.GetOrder("test order")
	if newOrder.Person != order.Person && err != nil {
		t.Errorf("cart get item failed")
	}
}

func TestCartRemoveOrder(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}

	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	// init slack routing with cart
	order := Order{Person: "test order"}
	order.AddItem("test item")
	_, _ = cart.AddOrder(order)

	if len(cart.Orders) == 0 {
		t.Errorf("cart add item failed")
	}
	// remove Orders
	err := cart.RemoveOrder("test order")
	if err != nil && len(cart.Orders) != 0 {
		t.Errorf("cart remove item failed")
	}
}

func TestCartAllOrders(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}

	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	// init slack routing with cart
	order := Order{Person: "test order"}
	order.AddItem("test item")
	_, _ = cart.AddOrder(order)

	if len(cart.Orders) == 0 {
		t.Errorf("cart add item failed")
	}
	allOrders := cart.AllOrders()
	if len(allOrders) == 0 {
		t.Errorf("cart all orders failed")
	}
}

func TestCartAddMenuLink(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}
	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	if len(cart.MenuLinks) != 0 {
		t.Errorf("cart menu links non empty on init")
	}
	// add menu link
	cart.AddMenuLink("test menu link")
	// check if menu link added
	if len(cart.MenuLinks) == 0 {
		t.Errorf("cart add menu link failed.")
	}

}

func TestCartRemoveMenuLink(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}
	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	if len(cart.MenuLinks) != 0 {
		t.Errorf("cart menu links non empty on init")
	}
	// add menu link
	cart.AddMenuLink("test menu link")
	// check if menu link added
	if len(cart.MenuLinks) == 0 {
		t.Errorf("cart add menu link failed.")
	}
	// remove menu link
	cart.RemoveMenuLink("test menu link")
	// check if menu link added
	if len(cart.MenuLinks) != 0 {
		t.Errorf("cart add menu link failed.")
	}
}

func TestCartAddCoAdmin(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}
	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	if len(cart.CoAdmins) != 0 {
		t.Errorf("cart co admins non empty on init")
	}
	// add co admin
	cart.AddCoAdmin("test co admin")
	// check if co admin added
	if len(cart.CoAdmins) == 0 {
		t.Errorf("cart add coadmin failed.")
	}

}

func TestCartRemoveCoAdmin(t *testing.T) {
	// mock values
	hours, mins, sec := 0, 0, 30
	// get current time
	now := time.Now()
	// get future time
	future := now.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins) + time.Second*time.Duration(sec))
	// create cart
	cart := Cart{Name: "test", Orders: []Order{}, Expires: future}
	if cart.Name != "test" {
		t.Errorf("cart init failed")
	}
	if len(cart.CoAdmins) != 0 {
		t.Errorf("cart co admins non empty on init")
	}
	// add co admin
	cart.AddCoAdmin("test co admin")
	// check if co admin added
	if len(cart.CoAdmins) == 0 {
		t.Errorf("cart add co admin failed.")
	}
	// remove co admin
	cart.RemoveCoAdmin("test co admin")
	// check if co admin removed
	if len(cart.CoAdmins) != 0 {
		t.Errorf("cart remove co admin failed.")
	}
}
