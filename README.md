# HandsOn Exercise - 01
    1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
        IDENTIFIERS “x” and “y” and “z”
        a. 42
        b. “James Bond”
        c. true
    2. Now print the values stored in those variables using
        a. a single print statement
        b. multiple print statements
    code: here’s the solution: https://play.golang.org/p/yYXnWXIQNa
  
# HandsOn Exercise - 02
    1. Use var to DECLARE three VARIABLES. The variables should have package level
    scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
        a. identifier “x” type int
        b. identifier “y” type string
        c. identifier “z” type bool
    2. in func main
        a. print out the values for each identifier
        b. The compiler assigned values to the variables. What are these values called?
    code: here’s the solution: https://play.golang.org/p/jzHwSlles9

# HandsOn Exercise - 03
    Using the code from the previous exercise,
    1. At the package level scope, assign the following values to the three variables
       a. for x assign 42
       b. for y assign “James Bond”
       c. for z assign true
    2. in func main
      a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
      b. print out the value stored by variable “s”

# HandsOn Exercise - 04
    ● FYI - nice documentation and new terminology “underlying type”
    ○ https://golang.org/ref/spec#Types For this exercise
    1. Create your own type. Have the underlying type be an int.
    2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR”
      keyword
    3. in func main
        a. print out the value of the variable “x”
        b. print out the type of the variable “x”
        c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
        d. print out the value of the variable “x”