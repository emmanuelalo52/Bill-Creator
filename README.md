# Bill-Creator
# Bill Management App

This simple Go application allows users to create bills, add items, specify tips, and save the bill details to a text file.

## How to Use

1. **Create a Bill:**
   - Run the program, and it will prompt you to create a new bill by entering its name.

2. **Add Items:**
   - Choose the "A" option to add items to your bill.
   - Enter the item name and price when prompted.

3. **Add Tip:**
   - Choose the "T" option to add a tip to your bill.
   - Enter the tip amount when prompted.

4. **Save Bill:**
   - Choose the "S" option to save the bill details to a text file in the "Bills" directory.

## Code Structure

The code consists of two main parts:

1. **Bill Package (bill.go):**
   - Defines a `bill` struct to represent a bill.
   - Provides methods to create a new bill, format the bill, update the tip, add items, and save the bill.

2. **Main Program (main.go):**
   - Defines functions to get user input, create a bill, and prompt user options.
   - The `main` function initializes a bill, prompts user options, and performs the corresponding actions.

## How to Run

Ensure you have Go installed on your machine.

```bash
go run main.go
