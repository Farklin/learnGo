package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func main() {

	if len(os.Args) > 3 {
		if os.Args[1] == "add" {
			var record = Record{
				Name:       os.Args[2],
				Surname:    os.Args[3],
				Number:     os.Args[4],
				LastAccess: os.Args[5],
			}

			err := addRowCsvFile("BookPhone.csv", record)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Строка успешно добавлена")
		}
	}

	//if len(os.Args) != 3 {
	//	fmt.Println("csvData input output!")
	//	return
	//}
	//
	//input := os.Args[1]
	//output := os.Args[2]
	//
	//lines, err := readCSVFile(input)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for _, line := range lines {
	//	temp := Record{
	//		Name:       line[0],
	//		Surname:    line[1],
	//		Number:     line[2],
	//		LastAccess: line[3],
	//	}
	//	myData = append(myData, temp)
	//
	//	fmt.Println(temp)
	//}
	//
	//err = saveCsvFile(output)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}

func readCSVFile(filename string) ([][]string, error) {
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error Stat file: " + err.Error())
		return [][]string{}, err
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error Open file: " + err.Error())
		return [][]string{}, err
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error read csv file:" + err.Error())
	}

	return lines, nil
}

func saveCsvFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error create file: " + err.Error())
		return err
	}

	defer file.Close()

	csvwriter := csv.NewWriter(file)

	csvwriter.Comma = '\t'

	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}

	csvwriter.Flush()
	return nil
}

func addRowCsvFile(filename string, record Record) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error open file: " + err.Error())
		return err
	}

	defer file.Close()

	csvwriter := csv.NewWriter(file)

	temp := []string{record.Name, record.Surname, record.Number, record.LastAccess}
	err = csvwriter.Write(temp)

	if err != nil {
		fmt.Println("Error write row file: ", err.Error())
		return err
	}

	csvwriter.Flush()

	return nil
}
