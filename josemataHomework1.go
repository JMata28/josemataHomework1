/*
This is homework assignment #1 by José Amilcar Mata Calidonio for the COMP 510 course.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("Welcome Jose Mata's word counter program.\nPlease enter the name of the text file you want to use (followed by .txt):")
	//Create string called nameoffile and save the user's input to it. This will be the name of the .txt file to open.
	var nameoffile string
	fmt.Scanln(&nameoffile)
	fmt.Println("The name of the file is:", nameoffile)
	/*
		The following three lines of code open up the .txt file and store the contents of the file into a string called str, so it is no longer handled as byte data.
		The if statement on Line 25 checks for errors when opening the .txt file.
	*/
	content, err := os.ReadFile(nameoffile) //This line of code was taken from the code used in our in-class Go program created during the lecture on January 29, 2024 and modified to be used in this program.
	if err != nil {
		log.Fatal("Help file error", err)
	} //This line of code was taken from the code used in our in-class Go program created during the lecture on January 29, 2024.
	str := string(content)
	/*
		The following three lines of code were taken from the instructions of this assignment and modified to be used in this program.
		Dr. Santore helped me by explaining what these three lines of code do.
	*/
	reg, err := regexp.Compile("[^a-zA-Z0-9]+") //This line creates a regular expression that matches anything that isn't a lower case letter, a capital letter, or a number.
	if err != nil {                             //This if statement checks for errors when create the regular expression
		log.Fatal(err)
	} //This line checks for errors in the creation of the regular expression.
	processedString := reg.ReplaceAllString(str, " ") /*This line goes through the str string and replaces anything
	that matches the reg regular expression with a space and stores the new string in the string called processedString.*/
	/*
		The string is then split into a slice and stored as "allwords". I was originally going to use strings.Split but instead I decided to use
		strings.Fields so that multiple spaces and/or a new line would also be used as a delimeter.
	*/
	allwords := strings.Fields(processedString)
	countWords(allwords) //Call the function countWords with allwords as the paramenter slice
}

func countWords(sliceofwords []string) {
	wordmap := make(map[string]int)      //create an empty map called wordmap
	for _, value := range sliceofwords { //for each word in the slice called sliceofwords
		wordmap[value]++ /*increase by one the count of each word (increase the value of the key by one).
		The keys of the map will be automatically created if there isn't one made for the word yet and will start with a value of 0.*/
	}
	reportResults(wordmap) //call the function reportResults and pass the map wordmap as the parameter
}

func reportResults(results map[string]int) {
	fmt.Println("The following is a list of each word followed by the amount of times it appears in the .txt file:")
	for k, v := range results { //for each key-value pair in the map results
		fmt.Println(k, ":", v) //print first the key (the word), then a colon, then the value (number of word counts) and continue in a new line
	}
}
