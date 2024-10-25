# Inventory Management TUI Application

## Overview

This Go application is a terminal-based inventory management tool designed to help users manage their stock efficiently. Utilizing the `tview` package, it provides a user-friendly Text User Interface (TUI) for adding, viewing, and deleting inventory items. The application supports persistent storage, saving the inventory data in a JSON file.

## Features

- **Add Items**: Easily add new items with names and stock quantities.
- **View Inventory**: Display all current items in stock.
- **Delete Items**: Remove items from the inventory using their item ID.
- **Persistent Storage**: Inventory is saved in `inventory.json`, which is loaded on startup.

## Prerequisites

- Go version 1.16 or higher
- `tview` package for creating the TUI

## Installation

   1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd inventory-management
   ```
   2. Install the required installeents [the only external dependency you'll need!]
   ```bash
   go get github.com/rivo/tview
   ```
   3. Run application
   ```bash
   go run main.go
   ```
## Usage

### Adding an Item
1. Enter the item name in the **Item Name** field.
2. Enter the stock quantity in the **Stock** field.
3. Click the **Add Item** button to save the item.

### Deleting an Item
1. Enter the item ID (index) in the **Item ID to delete** field.
2. Click the **Delete Item** button to remove it from the inventory.

### Exiting the Application
- Click the **Exit** button to close the application.

## Inventory Data Structure
The inventory items are stored in a JSON file (`inventory.json`) with the following structure:

```json
[
    {
        "name": "ItemName",
        "stock": 10
    },
    {
        "name": "AnotherItem",
        "stock": 5
    }
]
```
## Code Structure

### Main Application Logic:
- `loadInventory()`: Loads inventory data from `inventory.json`.
- `saveInventory()`: Saves current inventory state to `inventory.json`.
- `deleteItem(index int)`: Deletes an item from the inventory based on its ID.

### Dependencies:
- `encoding/json`: For JSON encoding/decoding.
- `fmt`: For formatted I/O.
- `log`: For error logging.
- `os`: For operating system interactions.
- `strconv`: For string conversion.
- `github.com/rivo/tview`: For TUI functionality.

## Contributing
Contributions are welcome! If you'd like to contribute, please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.


