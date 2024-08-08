# Library Management Program

This is a simple Library Management program written in Go. It allows users to manage a library system with functionalities such as adding, removing, borrowing, and returning books. The program also enables users to list all available and borrowed books.

## Features

- **Add a Book**: Add a new book to the library.
- **Remove a Book**: Remove a book from the library using its ID.
- **Borrow a Book**: Borrow a book from the library using the book ID and the member ID.
- **Return a Book**: Return a borrowed book to the library using the book ID and the member ID.
- **List Available Books**: List all books that are currently available in the library.
- **List Borrowed Books**: List all books borrowed by a specific member using the member ID.

## How to Run

1. **Clone the Repository**: 
   ```sh
   git clone github.com/Abzaek/library_management
   ```
2. **Navigate to the Project Directory**:
   ```sh
   cd library_management
   ```
3. **Run the Program**:
   ```sh
   go run main.go
   ```

## Commands

- **Add a Book**: 
  - Command: `/Add`
  - Prompts the user to enter the book's Title, ID, and Author.
  
- **Remove a Book**:
  - Command: `/Remove <bookID>`
  - Removes the book with the specified ID from the library.

- **Borrow a Book**:
  - Command: `/Borrow <bookID> <memberID>`
  - Borrows the book with the specified ID to the member with the given ID.

- **Return a Book**:
  - Command: `/Return <bookID> <memberID>`
  - Returns the book with the specified ID from the member with the given ID.

- **List Available Books**:
  - Command: `/List`
  - Lists all books that are currently available in the library.

- **List Borrowed Books**:
  - Command: `/ListBorrowed <memberID>`
  - Lists all books borrowed by the member with the given ID.

- **Exit the Program**:
  - Command: `/exit`
  - Exits the program.

## Example

After running the program, the following command prompt will be displayed:

```
Hello welcome! the following are the available commands:
    /Add                  to add a book
    /Remove bookID        to remove a book with ID "bookID"
    /Borrow bookID memberID  to borrow a book with ID "bookID" and a member ID "memberID"
    /Return bookID memberID  to borrow a book with ID "bookID" and a member ID "memberID"
    /List                 to list all the available books
    /ListBorrowed memberID   to list all the borrowed books by a member with a member ID "memberID"
    /exit                 to exit 
```

Users can input the corresponding commands to interact with the library system.


