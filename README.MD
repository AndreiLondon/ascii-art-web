Description:

Project converts input text into ascii-art representation based on selected banner. Coverted text is displayed in new web page.

Authors:

Aliaksei Vidaseu
Andrei Martynenko

Usage:
Start server by running go run main.go command
Open browser and in addres bar put localhost:8080 and press enter
In the text area enter text that has to be converted.
In drop-down list select banner
Press submit button

Implementation details:
Home page is displayed using http template package
Once user pressed submit button, POST request is being send to /ascii-art along with a text and a banner as parameters
Using the banner, server creates map with a character as a key and a 2D array as a value
Using this map, server converts each character into ascii-art representation, and sends output into http response writer 