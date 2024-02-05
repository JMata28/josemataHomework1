# josemataHomework1

This is Jose Amilcar Mata Calidonio's homework assignment #1 for the COMP 510 course. 

The instructions for this assignment are shown below:

Homework 1 : counting words
(in your class/homework grade this time)


Summary:
write a word counting program with go using at least three functions


Due: Sunday  Feb 4th by 11:59pm


get War and Peace as a text file https://www.gutenberg.org/files/2600/2600-0.txt



Details:

Be sure to create a new project with your name somewhere in the project name. (I might use jsantoreHomework1)


Write this in at least three functions



main should ask the user for the file name, and then open the file and convert it to a slice of strings using strings.Split then pass the slice of strings to the next functions
countWords should take a slice of strings as its parameter and then count the number of times each word occurs (using a map), it should return the map containing the word counting
finally reportResults should take a map of string to int (the word count map) and print out how often each word occurs in the file. One word count per line.


Note that since main needs to ask the user for a file, I can (and will) try multiple files when grading.


To get rid of punctuation use something like


// Make a Regex to say we only want letters and numbers
reg, err := regexp.Compile("[^a-zA-Z0-9]+")
if err != nil {
log.Fatal(err)
}
processedString := reg.ReplaceAllString(example, "")


Submitting the project.

submit the project by putting it on github and sending me a collaboration invite. (I'm git user jsantore)


if you need to min/max, here is how I'll grade