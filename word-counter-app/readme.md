When building applications that eill automate tasks, analyze data, parse logs, talk to network services or address other 
requirements, writing a CLI tool in golang to achieve all the stated scenarios will be the fastest way to do it. In this
module or project, i will be learning how to develop CLI applications using the go programming language. Starting this 
journey with creating a simple basic word counter tool. This tool reads number of words or lines provided as input using 
the standard input [STDIN]. This versiob reads data from stdin and displays the number of words. This will be the  first 
version of the app.

1. First step will be creating a directory for our project and initializing a go module.

        mkdir word-counter-app
        cd word-counter-app
        go mod init github.com/marviigrey/cli-app-development/word-counter-app

2. Create our main.go file.

        touch main.go


In the count() function, We made use of the NewScanner from the bufio package which is used for reading data, the newScanner function is used to read data delimited by spaces or new line. 

We also created a testing case to test the newly created tool, we declared a variable using the NewBufferString function of the byte package to read a string content.
The test is passed.

Next step is to add command line flags.
To implement this we make use of the flag package. This is the package used for implementing command line flag parsing. Command line tools provide flexibility through options, we are going to use the parsing to add options when we want to count words in our application with the flag package.
The "lines" variable defined in line 14:10 sets a new -l option thay we will use to indicate whether to count lines, the option is set to false by default which means the normal behaviour is to count just words. 
