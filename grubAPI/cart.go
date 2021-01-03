package main

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cart : struct to manage orders
type Cart struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name,omitempty"`
	Serialized string             `bson:"serialized,omitempty"`
	Orders     []Order            `bson:"orders,omitempty"`
	Expires    time.Time          `bson:"expires,omitempty"`
	Admin      string             `bson:"admin,omitempty"`
	CoAdmins   []string           `bson:"coadmins,omitempty"`
	MenuLinks  []string           `bson:"menulinks,omitempty"`
}

// Serialize : function to generate serialized cart name
func (c *Cart) Serialize() {
	// generate random string
	c.Serialized = RandomString(5)
}

// IsExpired : function to check if a cart is epxired
func (c *Cart) IsExpired() bool {
	// get deadline
	deadline := GetRemaining(c.Expires)
	// check if deadline has passed
	if deadline.t <= 0 {
		return true
	}
	return false
}

// AddOrder : function to add order to a cart
func (c *Cart) AddOrder(o Order) (Order, error) {
	// get deadline
	deadline := GetRemaining(c.Expires)
	// check if deadline has passed
	if deadline.t <= 0 {
		// raise error
		return Order{}, errors.New("cart has expired")
	}
	// append to cart items
	c.Orders = append(c.Orders, o)
	// return order
	return o, nil
}

// RemoveOrder : function to remove order from a cart
func (c *Cart) RemoveOrder(person string) error {
	// iterate through items until found
	for idx, item := range c.Orders {
		// check if person name matches
		if item.Person == person {
			// remove item
			c.Orders[len(c.Orders)-1], c.Orders[idx] = c.Orders[idx], c.Orders[len(c.Orders)-1]
			c.Orders = c.Orders[:len(c.Orders)-1]
			// return success
			return nil
		}
	}
	return errors.New("order not found")
}

// GetOrder : function to get order from a cart
func (c *Cart) GetOrder(person string) (Order, error) {
	// iterate through items
	for _, item := range c.Orders {
		// check if person matches
		if item.Person == person {
			// return order
			return item, nil
		}
	}
	return Order{}, errors.New("order not found")
}

// AllOrders : function to get all orders from a cart
func (c *Cart) AllOrders() []Order {
	return c.Orders
}

// NumOrders : function to get num orders in a cart
func (c *Cart) NumOrders() int {
	return len(c.Orders)
}

// AllMenuLinks : function to get all menu links from a cart
func (c *Cart) AllMenuLinks() []string {
	return c.MenuLinks
}

// AddMenuLink : function to add menuLink to a cart
func (c *Cart) AddMenuLink(menuLink string) {
	c.MenuLinks = append(c.MenuLinks, menuLink)
}

// RemoveMenuLink : function to remove menuLink from a cart
func (c *Cart) RemoveMenuLink(menuLink string) {
	// check if menu link matches
	for idx, item := range c.MenuLinks {
		// check if menu link matches
		if item == menuLink {
			// remove item
			c.MenuLinks[len(c.MenuLinks)-1], c.MenuLinks[idx] = c.MenuLinks[idx], c.MenuLinks[len(c.MenuLinks)-1]
			c.MenuLinks = c.MenuLinks[:len(c.MenuLinks)-1]
		}
	}
}

// AllCoAdmins : function to get all CoAdmins from a cart
func (c *Cart) AllCoAdmins() []string {
	return c.CoAdmins
}

// AddCoAdmin : function to add CoAdmin to a cart
func (c *Cart) AddCoAdmin(coAdmin string) {
	c.CoAdmins = append(c.CoAdmins, coAdmin)
}

// RemoveCoAdmin : function to remove CoAdmin to a cart
func (c *Cart) RemoveCoAdmin(coAdmin string) {
	// iterate through co admins
	for idx, item := range c.CoAdmins {
		// check if co admin matches
		if item == coAdmin {
			// remove item
			c.CoAdmins[len(c.CoAdmins)-1], c.CoAdmins[idx] = c.CoAdmins[idx], c.CoAdmins[len(c.CoAdmins)-1]
			c.CoAdmins = c.CoAdmins[:len(c.CoAdmins)-1]
		}
	}
}
