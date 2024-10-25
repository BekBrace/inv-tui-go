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
   
2.Install the required installeents 
```bash
go get github.com/rivo/tview
```
3. Run application
```bash
go run main.go
```
