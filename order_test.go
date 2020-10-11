package main

import (
	"testing"
)

func TestOrderInit(t *testing.T) {
	order := Order{Person: "test"}
	if order.Person != "test" {
		t.Errorf("order init failed")
	}
}

func TestOrderAddItem(t *testing.T) {
	order := Order{Person: "test"}
	order.AddItem("test_item")
	if len(order.Items) == 0 {
		t.Errorf("order add item failed.")
	}
}

func TestOrderGetItem(t *testing.T) {
	order := Order{Person: "test"}
	order.AddItem("test_item")
	if len(order.Items) == 0 {
		t.Errorf("order add item failed.")
	}
	item, err := order.GetItem("test_item")
	if err != nil {
		t.Errorf("order get item failed.")
	}
	if item != "test_item" {
		t.Errorf("order get item failed.")
	}
}

func TestOrderRemoveItem(t *testing.T) {
	order := Order{Person: "test"}
	order.AddItem("test_item")
	if len(order.Items) == 0 {
		t.Errorf("order add item failed.")
	}
	order.RemoveItem("test_item")
	if len(order.Items) != 0 {
		t.Errorf("order remove item failed.")
	}
}

func TestOrderAddNote(t *testing.T) {
	order := Order{Person: "test"}
	order.AddNotes("test note")
	if order.Notes == "" {
		t.Errorf("order add note failed.")
	}
}

func TestOrderGetNote(t *testing.T) {
	order := Order{Person: "test"}
	order.AddNotes("test note")
	if order.Notes == "" {
		t.Errorf("order add note failed.")
	}
	if order.GetNotes() != "test note" {
		t.Errorf("order get note failed.")
	}
}

func TestOrderRemoveNote(t *testing.T) {
	order := Order{Person: "test"}
	order.AddNotes("test note")
	if order.Notes == "" {
		t.Errorf("order add note failed.")
	}
	order.RemoveNotes()
	if order.Notes != "" {
		t.Errorf("order remove note failed.")
	}
}
