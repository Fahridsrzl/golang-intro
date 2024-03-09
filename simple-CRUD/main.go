package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Title       string
	Author      string
	Releaseyear string
	Pages       int
}

var books []Book

var fileName string = "data.csv"

func main() {
	// createFile(fileName)
	// addNewBook()
	loadDataFromCSV(fileName)
	viewAllBooks()

}

// func createFile(fileName string) {
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("file", fileName, "berhasil dibuat.")
// 	}
// 	defer file.Close()
// }

func addNewBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Book Details :")
	fmt.Print("Book Id :")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title :")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author:")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Release Year :")
	scanner.Scan()
	newBook.Releaseyear = scanner.Text()

	fmt.Print("Book Pages :")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	_, err := findBookById(newBook.Id)
	if err != nil {
		books = append(books, newBook)
	} else {
		return fmt.Errorf("Book with id: %d already exits", newBook.Id)
	}

	err = saveDataToCSV(fileName)
	if err != nil {
		return err
	}
	fmt.Println("Book added succesfully")

	return nil
}

func saveDataToCSV(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.Releaseyear + "," + strconv.Itoa(book.Pages) + "\n"
		file.WriteString(row)
	}
	return nil
}

func loadDataFromCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		// fmt.Println(record)
		id, _ := strconv.Atoi(record[0])
		pages, _ := strconv.Atoi(record[4])

		book := Book{
			Id:          id,
			Title:       record[1],
			Author:      record[2],
			Releaseyear: record[3],
			Pages:       pages,
		}
		books = append(books, book)

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}

	return nil
}

func viewAllBooks() error {
	if len(books) == 0 {
		return fmt.Errorf("no books available")
	}
	for i, book := range books {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("Book -", i+1)
		fmt.Println("Book Id :", book.Id)
		fmt.Println("Book Title :", book.Title)
		fmt.Println("Book Author :", book.Author)
		fmt.Println("Release Year :", book.Releaseyear)
		fmt.Println("Pages :", book.Pages)
		fmt.Println(strings.Repeat("=", 50))
	}
	return nil
}

func findBookById(id int) (Book, error) {
	for _, book := range books {
		if book.Id == id {
			return book, nil
		}
	}
	return Book{}, fmt.Errorf("id : %d not found", id)
}
