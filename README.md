
# ASCII-ART

A command-line tool that converts text input into ASCII art representations. Utilizing a collection of standard ASCII characters, the program transforms text input into visually appealing designs. Users can generate ASCII art banners enhancing text-based communication with creative visuals.


## Installation

Prerequisites:

Go 1.18 or higher installed
Git installed
Operating System: Windows, macOS, Linux

Download/Clone:
```
$ git clone https://learn.zone01kisumu.ke/git/cowalla/ascii-art
```

Preconditions:
```
If you encounter any issues, make sure your input text contains only ASCII characters (32 to 126). Non-ASCII characters are not supported. The number of arguments must be strictly 2 for the program to work. The standard.txt file must not be empty for the program to work.

Error: Input contains non-ASCII characters

```

## Features

- Command-Line Interface: The project is a command-line tool, allowing users to interact with it through their terminal or command prompt.

- Input Handling: It accepts text input from the command line and processes it to generate ASCII art representations.

- Special Character Handling: It handles special characters such as newline (\n), tab (\t), and backspace (\b) in the input text, providing functionality to interpret and display them correctly in the ASCII art output.

- Error Handling: It includes error handling mechanisms to handle incorrect input or file reading errors gracefully. For example, it checks for incorrect number of command-line arguments and file reading errors.

- File Reading: It reads content from a file named "standard.txt" to generate ASCII art. This indicates the ability to read external resources or files for generating art. By default, this program reads content from the standard.txt file. You can also replace standard.txt with shadow.txt and thinkertoy.txt files in the program.

- ASCII Art Generation: It generates ASCII art based on the input text and the contents of the "standard.txt" file. The specifics of how the ASCII art is generated would depend on the implementation details within the ascii.AsciiArt function, which is likely contained in the ascii package imported in your code.

- Output Display: It displays the generated ASCII art to the user, typically in the terminal or command prompt where the tool is executed.


## Usage/Examples

```golang
import ascii/ascii

func main() {
  ascii.AsciiArt(words, contents2)
}
```

## Executing program

To run the program, navigate to the directory where the program is installed and use the following command:

``` go run main.go "Your Text Here" ```

Replace "Your Text Here" with the text you want to convert into ASCII art.

## Running Tests

To run tests, run the following command

```bash
  go test -v
```

## License

[MIT](https://choosealicense.com/licenses/mit/)


## Demo

https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExam4wOTZzdmJmOHIzNGFpbTZsOWhsdm95YndzcWYxaGRoeTV4YmN2NiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/JHlHhoAkrA7pDqAItE/giphy.gif


## Authors

- [@cowalla](https://learn.zone01kisumu.ke/git/cowalla)

- [@mombewa](https://learn.zone01kisumu.ke/git/mombewa)

- [@skisenge](https://learn.zone01kisumu.ke/git/skisenge)