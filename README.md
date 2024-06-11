# Project: Text Formatter in Go

This project is a text formatter written in Go. It takes a text file as input and performs various modifications to the text, saving the modified version to a specified output file.

## Supported Modifications:

    Hexadecimal to Decimal: Convert words prefixed with "(hex)" to their decimal value (e.g., "1E (hex)" becomes "30").
    Binary to Decimal: Convert words prefixed with "(bin)" to their decimal value (e.g., "10 (bin)" becomes "2").
    Uppercase/Lowercase: Convert words with "(up)" or "(low)" to uppercase/lowercase, respectively. Optional numeric argument specifies the number of words to affect.
    Capitalize: Convert a word with "(cap)" to title case.
    Punctuation Spacing: Add spaces around punctuation marks (".", ",", "!", "?", ":", ";") except for ellipses ("...") and question/exclamation combinations ("!?").
    Single Quotes: Surround a word or phrase with single quotes, handling multiple words appropriately.
    "An" Conversion: Convert "a" to "an" before a word starting with a vowel or silent "h".

## Usage
This program takes two arguments:
- `Input file`:The text file we are modifying in this case; `sample.txt`
- `Output file`: the text file where the modified file will be saved; result.txt
	
## How to use:
Open your terminal window
Create a `sample.txt` file and insert the content you want modify.
```go
	$ echo "Your text to modify" > sample.txt
```
To run the program you run the following command:
```go
	$ go run . sample.txt result.txt
	$ cat result.txt (to see the modified text)
	$ 	(modified text appear here)
```	
	
## Testing
Test file has also been included. To run the file to test the functions, on the terminal type:
```go
	$ go test -v
```
## Author
[skisenge](https://learn.zone01kisumu.ke/git/skisenge)