package main

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order : struct to keep track of orders
type Order struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Person string             `bson:"person,omitempty"`
	Items  []string           `bson:"items,omitempty"`
	Notes  string             `bson:"notes,omitempty"`
}

// AddItem : function to add item to order
func (o *Order) AddItem(itemName string) {
	// add item to items slice
	o.Items = append(o.Items, itemName)
}

// GetItem : function to get item from order
func (o *Order) GetItem(itemName string) (string, error) {
	// iterate through items until found
	for _, item := range o.Items {
		// check if item name matches
		if item == itemName {
			return item, nil
		}
	}
	return "", errors.New("item not found")
}

// RemoveItem : function to remove item from order
func (o *Order) RemoveItem(itemName string) {
	// iterate through items until found
	for idx, item := range o.Items {
		// check if item name matches
		if item == itemName {
			o.Items[len(o.Items)-1], o.Items[idx] = o.Items[idx], o.Items[len(o.Items)-1]
			o.Items = o.Items[:len(o.Items)-1]
		}
	}
}

// AddNotes : function to add notes to order
func (o *Order) AddNotes(note string) {
	// add notes
	o.Notes = note
}

// GetNotes : function to get notes for an order
func (o *Order) GetNotes() string {
	return o.Notes
}

// RemoveNotes : function to remove notes from an order
func (o *Order) RemoveNotes() {
	// remove notes
	o.Notes = ""
}
