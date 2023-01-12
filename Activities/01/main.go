/*
1. Declare the package name to be able to execute this file
2. Import fmt
3. MAIN ACTIVITY:
	global variables:
        Add the following global variables to the file (values only goes up to four to keep it simple)

		type Card struct{
			suit string,
			value string,
		}

		suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
		values := []string{"Ace", "Two", "Three", "Four"}

	main method:
		- Add your main method that calls a function called newCard() with two string parameters
		  passed in (to keep it simple, you may pass in two 'hard-coded' string values instead of
		  asking for user input).
		  	- newCard() function will return a string
			- Store this returned value in a string variable
		- Print out the string variable to the console.

	newCard method:
		- Create the newCard method so that it takes in two string parameter values.
			- The first parameter is the value of the card
			- The second parameter is the suit of the card
		- This method will check the user's input against the expected values in the global variables
		- Once the variables are verified, it will return a card Struct that is made up of the passed
			in suit and value.
		- If the user input is incorrect,return an error

			HINT: Think to use two seperate for loops to check the two parameters.

	print method:
		- Create a receiver function of type Card that takes no arguments
		- Prints a card to the console/terminal
		- The print statement should be formatted to read the value of the card first followed
		  by " of " and then the suit.
			- EXAMPLE: "Four of Spades"

BEGIN CODE BELOW:
*/