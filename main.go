package main

import (
	"encoding/json" // Handles JSON encoding/decoding
	"fmt"           // For formatted I/O like printing to the console
	"log"           // Handles logging errors
	"os"            // Provides operating system functionality like file handling
	"strconv"       // Converts strings to other types, like converting string to int

	// tview page in GO
	// https://pkg.go.dev/github.com/rivo/tview#section-readme
	"github.com/rivo/tview" // This is the TUI package that provides user interface elements
)

// Define an Item structure that will hold the stock information
type Item struct {
	Name  string `json:"name"`  // Name of the item (will be stored as JSON)
	Stock int    `json:"stock"` // Quantity of the item in stock (also stored as JSON)
}

// Initialize an empty slice to store the inventory and set the filename for persistence
var (
	inventory     = []Item{}         // Our inventory list, initially empty
	inventoryFile = "inventory.json" // File where inventory will be saved/loaded from
)

// This function loads the inventory from the JSON file
func loadInventory() {
	// Check if the file exists before attempting to load it
	// // The _ is a blank identifier used in Go to ignore a value that you don't need to use later in your code.
	// if _, err := os.Stat(inventoryFile); err == nil {
	// Here, os.Stat returns two values: the file information (of type os.FileInfo) and an error. Since you only care about whether the file exists (i.e., if there is an error), you can ignore the file information by assigning it to _. This way, you avoid cluttering your code with variables you don't use.
	if _, err := os.Stat(inventoryFile); err == nil {
		// Read the contents of the file
		data, err := os.ReadFile(inventoryFile)
		if err != nil {
			log.Fatal("Error reading inventory file:", err)
		}
		// Parse the JSON data into the inventory slice
		json.Unmarshal(data, &inventory)
	}
}

// This function saves the current inventory to the JSON file
func saveInventory() {
	// Marshal the inventory into JSON with pretty formatting (indentation)
	// The MarshalIndent function ensures that the JSON is "pretty-printed," meaning it adds spaces and newlines for easier reading.
	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		log.Fatal("Error saving inventory:", err)
	}
	// Write the JSON data back to the file, overwriting the old data

	// os.WriteFile writes the JSON data to the file specified by inventoryFile.
	// The 0644 sets the file permissions: the owner can read and write, while others can only read the file.
	// This overwrites the old inventory data in the file with the new data in data.
	os.WriteFile(inventoryFile, data, 0644)
}

// Deletes an item from the inventory based on its index
func deleteItem(index int) {
	// Check if the index is valid
	if index < 0 || index >= len(inventory) {
		fmt.Println("Invalid item index.")
		return
	}
	/*
		The goal here is to remove an item from the inventory slice at a specific index.
		inventory[:index]:
		This takes a slice of all elements before the item you want to remove. It creates a new slice that starts from the beginning (0) up to but not including index.
		Example: If inventory = [A, B, C, D] and index = 2, then inventory[:index] will give you [A, B].*/
	// Remove the item from the inventory using slicing
	// append(inventory[:index], inventory[index+1:]...):

	// The append function joins the two slices: everything before the index and everything after the index, effectively skipping the item at index.append(inventory[:index], inventory[index+1:]...):

	// The append function joins the two slices: everything before the index and everything after the index, effectively skipping the item at index.
	// The ... syntax is used to unpack the slice inventory[index+1:] so that its elements are appended individually.
	// The ... syntax is used to unpack the slice inventory[index+1:] so that its elements are appended individually.
	inventory = append(inventory[:index], inventory[index+1:]...)
	// Save the updated inventory back to the file
	saveInventory()
}

// Main function, where the program execution begins
func main() {
	// Create a new TUI application
	app := tview.NewApplication()

	// Load existing inventory from the JSON file
	loadInventory()

	// Create a TextView that will display the inventory items in the TUI
	inventoryList := tview.NewTextView().
		SetDynamicColors(true). // Enable dynamic coloring of text
		SetRegions(true).       // Allows regions for interaction (not used here)
		SetWordWrap(true)       // Enables word wrapping to fit the TextView size

	inventoryList.SetBorder(true).SetTitle("Inventory Items") // Set border and title

	// This function refreshes the inventory display whenever there are changes
	refreshInventory := func() {
		// Clear the current content of the TextView
		inventoryList.Clear()
		// If inventory is empty, display a message
		if len(inventory) == 0 {
			fmt.Fprintln(inventoryList, "No items in inventory.")
		} else {
			// Iterate through inventory and print each item to the TextView
			for i, item := range inventory {
				fmt.Fprintf(inventoryList, "[%d] %s (Stock: %d)\n", i+1, item.Name, item.Stock)
			}
		}
	}

	// Create input fields for item name and stock quantity
	itemNameInput := tview.NewInputField().SetLabel("Item Name: ")
	itemStockInput := tview.NewInputField().SetLabel("Stock: ")

	// Create an input field for deleting an item by its index (ID)
	itemIDInput := tview.NewInputField().SetLabel("Item ID to delete: ")

	// Create a form that lets the user add or delete items
	form := tview.NewForm().
		AddFormItem(itemNameInput).    // Add the item name input to the form
		AddFormItem(itemStockInput).   // Add the item stock input to the form
		AddFormItem(itemIDInput).      // Add the item ID input for deletion
		AddButton("Add Item", func() { // Button to add a new item
			// Get the text input for name and stock
			name := itemNameInput.GetText()
			stock := itemStockInput.GetText()
			// Check if both fields are filled
			if name != "" && stock != "" {
				// Convert the stock input to an integer
				quantity, err := strconv.Atoi(stock)
				if err != nil {
					fmt.Fprintln(inventoryList, "Invalid stock value.")
					return
				}
				// Add the new item to the inventory slice
				inventory = append(inventory, Item{Name: name, Stock: quantity})
				// Save the updated inventory
				saveInventory()
				// Refresh the inventory display
				refreshInventory()
				// Clear the input fields after adding the item
				itemNameInput.SetText("")
				itemStockInput.SetText("")
			}
		}).
		AddButton("Delete Item", func() { // Button to delete an item
			idStr := itemIDInput.GetText()
			// Ensure the ID field is not empty
			if idStr == "" {
				fmt.Fprintln(inventoryList, "Please enter an item ID to delete.")
				return
			}
			// Convert the ID to an integer and check if it's valid
			id, err := strconv.Atoi(idStr)
			if err != nil || id < 1 || id > len(inventory) {
				fmt.Fprintln(inventoryList, "Invalid item ID.")
				return
			}
			// Delete the item (adjust for zero-based index)
			deleteItem(id - 1)
			fmt.Fprintf(inventoryList, "Item [%d] deleted.\n", id)
			// Refresh the inventory display after deletion
			refreshInventory()
			itemIDInput.SetText("") // Clear the ID input field
		}).
		AddButton("Exit", func() { // Button to exit the application
			app.Stop()
		})

	// Set a border and title for the form
	form.SetBorder(true).SetTitle("Manage Inventory").SetTitleAlign(tview.AlignLeft)

	// Create a layout using Flex to display the inventory list and the form side by side
	flex := tview.NewFlex().
		AddItem(inventoryList, 0, 1, false). // Left side: inventory list
		AddItem(form, 0, 1, true)            // Right side: form for adding/deleting items

	// Initial inventory display
	refreshInventory()

	// Start the TUI application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
