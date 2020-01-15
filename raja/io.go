package raja

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Io : struct being used for this scanner
type Io struct {
	scanner *bufio.Scanner
}

//NewIo : creates new instance of the Io struct
func NewIo() *Io {
	return &Io{
		scanner: getScanner(),
	}
}

func getScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

	return scanner
}

//ReadLineWithMessage : reads line from standard input stream
func (io *Io) ReadLineWithMessage(message string) string {
	fmt.Print(message)
	return io.ReadLine()
}

//ReadLine : reads line from standart input stream
func (io *Io) ReadLine() string {
	io.scanner.Scan()
	line := io.scanner.Text()
	line = strings.Trim(line, " ")
	return line
}

//ReadArray : Reads 2 diamentional array from standard input stream
func (io *Io) ReadArray(message string) ([][]int, int, int) {
	fmt.Println(message)
	var lines []string
	line := io.ReadLine()
	for len(line) != 0 {
		lines = append(lines, line)
		line = io.ReadLine()
	}
	rows := len(lines)
	cols := -1
	array := make([][]int, rows)
	for i := 0; i < rows; i++ {
		line = strings.ReplaceAll(lines[i], "  ", " ")
		nums := strings.Split(line, " ")
		ccols := len(nums)
		if cols != -1 && cols != ccols {
			println("Error : column '", line, "' size is not equal to", cols)
			os.Exit(1000001)
		}
		cols = ccols
		array[i] = make([]int, ccols)
		for j := 0; j < cols; j++ {
			array[i][j], _ = strconv.Atoi(nums[j])
		}
	}
	return array, rows, cols
}

// CreatEmptyTwoDiamentionalArray : creates two diamentional array with all value zero
func CreatEmptyTwoDiamentionalArray(rows int, cols int) [][]int {
	array := make([][]int, rows)
	for i := 0; i < rows; i++ {
		array[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			array[i][j] = 0
		}
	}
	return array
}

//PrintArray : prints array in butified way
func PrintArray(array [][]int) {
	max := 1
	for i, line := range array {
		for j := range line {
			len := len(strconv.Itoa(array[i][j]))
			if len > max {
				max = len
			}
		}
	}
	max++
	padding := "%" + strconv.Itoa(max) + "v"
	for i, line := range array {
		for j := range line {
			fmt.Printf(padding, array[i][j])
		}
		fmt.Println()
	}
}
