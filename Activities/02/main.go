/*
	1. Declare the package name to be able to execute this file
	2. Import "fmt", "time", and "math/rand"
	3. MAIN ACTIVITY:

		Your goal is to replicate the project taught in the powerpoint
		presentation today.  You will have a main function that calls
		three functions, newDeck(), shuffle(), and print().  Additionally,
		if you would like, you can add a Deal() function.

	Global Variables:
		Add the following global variable to the file:

			type deck []string

	main():
		- Initialize a variable called 'cards' that calls newDeck()
		- Call the shuffle() function on 'cards' to shuffle the deck
		- Call the print() function on 'cards' to print the deck to console

	newDeck():
		- Takes no arguments but returns a 'deck' type
		- Initalizes a variable with an empty deck
		- Create two variables, values and suits that are a slice of strings
		  The values should be the face value of a normal deck of cards
		  (EX. "Ace" or "King") and the suits should be the four suits a card
		  can have (EX. "Spades")
		- Loop through all the suits and values and append them to the variable
		  you initialized to be an empty deck (EXAMPLE OF APPENDED VALUE: "Ace of Spades")
		- Return the deck variable

	shuffle():
		- Receiver function of type 'deck' that takes no arguments and does not
		  return any values
		- Initialize a variable of type Source that takes in a Seed as an argument
		  	-- HINT: You will have to use the 'time' package as the Seed
		- Initalize a variable of type Rand using the Source variable as the seed
		- Loop through the entire deck generating random position values on each
		  iteration of the loop and setting each card in the deck to a new random
		  position.
		    -- HINT: Use the variable of type Rand with an integer generating function
			   You may wish to look at the 'math/rand' package for this.

		- NOTE: You may have to read some of the golang documentation to complete
		  this function as described.
		  	- https://pkg.go.dev/math/rand
			- https://pkg.go.dev/time

	print():
		- Receiver function of type 'deck' that takes no arguments and returns no
		  values
		- Loops through the deck of cards and prints each card to the console.

	deal():
		- A method that takes two arguments, the first of type deck and the second
		  of type int.  The method returns two deck types.
		- Split the deck that was passed in as a parameter into two decks using the
		  int passed in as the position on where the deck should be split.
		- Return the two decks

BEGIN CODE BELOW:
*/
	