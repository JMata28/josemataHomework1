/*
This is homework assignment #1 by José Amilcar Mata Calidonio for the COMP 510 course.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Welcome Jose Mata's word counter program.\nPlease enter the name of the text file you want to use (followed by .txt):")
	//Create string called nameoffile and save the user's input to it. This will be the name of the txt file to open.
	var nameoffile string
	fmt.Scanln(&nameoffile)
	fmt.Println("The name of the file is:", nameoffile)
	/*
		Open up the file and store the contents of the file into a string called content, else it will print out the byte data.
		The if err loop checks for errors.
		The string is then split into slices and stored as "allwords". I was originally going to use strings.Split but instead I decided to use
		strings.Fields so that a new line would also be used as a delimeter.
	*/
	content, err := os.ReadFile(nameoffile)
	if err != nil {
		log.Fatal("Help file error", err)
	}
	str := string(content)
	allwords := strings.Fields(str)
	countWords(allwords) //Call the function countWords with allwords as the paramenter slice
}

func countWords(sliceofwords []string) {
	wordmap := make(map[string]int)      //create an empty map
	for _, value := range sliceofwords { //for each word in the sliceofwords slice
		wordmap[value]++ //increase by one the count of each word (increase the value of the key by one)
	}
	reportResults(wordmap)
}

func reportResults(wordmap map[string]int) {
	fmt.Println("The following is a list of each word followed by the amount of times it appears:")
	for k, v := range wordmap {
		fmt.Println(k, ":", v)
	}
}

/*
	loop to print out the number of words in the slice allwords
	for index, value := range allwords {
		fmt.Println(index, value)
	}
*/

/*
	In case the professor doesn't let me use strings.Fields:
	There are two separators that I want to use: a space and a new line, so that every word (each substring) is its own element in the slice.
	To achieve this spliting with two separators, I used the "iterative splitting" method, which is based on the code found in this website:
	https://marketsplash.com/tutorials/go/golang-split-string/

	delimiters := []string{" ", "\n"} //I'm creating a slice that contains the delimiters that I want to use to split the string
	for _, delimiter := range delimiters { //For each delimiter in the delimiters slice
		segments := strings.Split(str, delimiter)   //In the segments slice, include the substring separated by the delimiter
		str = strings.Join(segments, " ")
	}
	for i := 0; i <len(Allwords);
*/
