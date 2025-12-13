# Golang

# Installation

Download the file from   
[https://go.dev/](https://go.dev/)

Follow these instructions to install on Linux  
[https://go.dev/doc/install](https://go.dev/doc/install)

NOTE:   
You will need to add the PATH string to your bashrc and call source on it

# Core Functions

## Main

When creating a new golang project we need a package called main with a function called main.  
This can exist in a file by a different name though.

So the function hello could contain:

######  *package main*

###### 

###### *func main() {*

###### 

###### *}*

NOTE:   
No two files in the same folder can have a “main” function.  
If you have two packages and each package has a main function then those packages need to be separated out into their own folders.

# Style Guide

Golang has a clear style guide that can be found here:  
[https://go.dev/doc/effective\_go](https://go.dev/doc/effective_go)

# CLI Commands

## Go

If you just type “go” on it’s own you see a list of all available commands.

## Run

Go Run allows you to run scripts

###### *go run [main.go](http://main.go)*

## Build

Allows you to build packages

###### 

###### *go build [main.go](http://main.go)*

Once a file has been compiled it will create a new file.  
In this instance the new file will be called “main”.

## fmt

This can be called on a file to fix it’s formatting

###### *go fmt [main.go](http://main.go)*

## mod {#mod}

This creates a new go module

###### *go mod init hello-go*

In our example we have one package called “main”, so that will be bundled into our module, but there could be multiple packages.

This will create a file called “[go.mod](#go.mod)”

This is where the name of the module, version, and required packages will be stored.

### tidy

We can use the mod tidy command to lint or tidy the comments in our module

###### *go mod tidy*

NOTE: You do not need to specify the file name for this to work.

# App Files

## go.mod {#go.mod}

This is the [module](#module) file created when you run the [mod](#mod) command. This file contains the module name, version and required packages.

## go.sum {#go.sum}

This file is created when you install a [new package](#3rd-party-packages) into a [module](#module). This is a generated system file that should not be modified.

# Go Packages

## fmt

The [Golang FMT](https://pkg.go.dev/fmt) package allows you to print to console and other functions

### Println

Prints strings to console

NOTE: Strings must be placed within double quotation marks

######  *package main*

###### 

###### *import "fmt"*

###### 

###### *func main() {*

###### 	*fmt.Println("I want a cup of coffee\!")*

###### *}*

Println automatically injects a space between the arguments that it receives.  
It also transforms floats slightly so this

###### *package main*

###### 

###### *import "fmt"*

###### 

###### *func main() {*

###### 	*var coffeeName \= "Espresso"*

###### 

###### 	*var size \= "Small"*

###### 

###### 	*var price \= 2.50*

###### 

###### 	*fmt.Println(size, coffeeName, "price is $", price)*

###### *}*

###### 

Will output this

###### *Small Espresso price is $ 2.5Small Espresso price is $ 2.5*

Notice the space between the dollar sign and the number, and the lack of a 0 on 2.5.

println also seems to add a new line at the end correctly as well, which is good for just printing things quickly to console I guess.

### Printf

This works just like sprintf in PHP, same syntax

###### *fmt.Printf(“%s is %d years old.\\n”, name, age)*

### Scanln {#scanln}

Allows you to accept CLI input as an argument

###### *fmt.Print("Enter your coffee name (or type 'exit' to quit): ")*

###### *fmt.Scanln(\&order)*

If we use Scanln within a [for loop](#infinite-for-loop-\(like-a-while\)) normally the loop will call it infinitely.  
It won’t stop to check to see what type of input was received.  
In order to make it stop and wait to check what input was received we need to pass in the variable as [a pointer to its position in memory](#pointers). If we do that, then when we reach Scanln the system will wait to receive input before continuing to the next iteration

### Stringer Interface

Go includes a special interface called Stringer, which is used by the fmt package to decide how a custom type should be printed. The Stringer interface is very simple:

## *type Stringer interface {*

##     *String() string*

## *}*

If your type has a String() method that returns a string, then it automatically satisfies this interface. This means that when you print your type using fmt.Println, fmt.Printf, or any other fmt function, Go will call your String() method to decide what text to display.

Here is an example:

## *type Coffee struct {*

##     *Name string*

##     *Size string*

## *}*

## 

## *func (c Coffee) String() string {*

##     *return fmt.Sprintf("A %s %s coffee", c.Size, c.Name)*

## *}*

## 

## *func main() {*

##     *order := Coffee{Name: "Latte", Size: "Large"}*

##     *fmt.Println(order)*

## *}*

Output:

## *A Large Latte coffee*

Because we added the String() method, Go knows exactly how to print the order variable. Without this method, printing the struct would show its raw field values, which is usually less readable.

The Stringer interface is useful for:

Making console output easier to read

Creating custom string representations for debugging

Formatting structs in a more user-friendly way

**NOTE:** You do not need to say “this type implements Stringer”. Go figures it out automatically as long as your type has a String() method.

## os

The Go os package provides tools for interacting with the operating system, such as working with files, directories, and environment variables.

### Open (os)

Opens a file for reading. Returns two values: a \*File and an error.

## *file, err := os.Open("coffee.txt")*

### Create (os)

Creates a new file or truncates an existing file for writing. Returns a \*File and an error.

file, err := os.Create("coffee\_orders.txt")

### Remove (os)

Deletes a file from the filesystem. Returns an error.

## *err := os.Remove("old\_orders.txt")*

Example

## *package main*

## *import (*

##     *"fmt"*

##     *"os"*

## *)*

## 

## *func main() {*

##     *file, err := os.Open("coffee\_orders.txt")*

##     *if err \!= nil {*

##         *fmt.Println("Error opening file:", err)*

##         *return*

##     *}*

##     *defer file.Close()*

##     *fmt.Println("Successfully opened file:", file.Name())*

## *}*

### File Interface (os)

\*os.File implements the io.Reader and io.Writer interfaces, which are used to read from and write to files and other streams.

## errors

The errors package provides tools for creating and handling error values in Go. Errors are treated as normal values, not exceptions.

Errors are not like exceptions, in that they do not stop the program from running. They are only there to detail that something went wrong. Only a [panic](#panic-\(defer\)) stops the program from running and subsequently requires the use of [recover](#recover-\(defer\)).

* Use errors for expected problems.  
* Use panic for truly unexpected situations that indicate a bug or unrecoverable state.

  ### New (errors)

  Creates a new error with a simple message.  
    
  err := errors.New("file not found")

  ### Is (errors)

  Checks if an error matches a specific target error, often used with standard library errors like os.ErrNotExist.  
    
  if errors.Is(err, os.ErrNotExist) { ... }

  ### As (errors)

  Checks if an error can be cast to a specific type and retrieves it. Useful for handling custom error types.


  ## *var pathErr \*os.PathError*

  ## *if errors.As(err, \&pathErr) { ... }*

    
  You can think of errors.As in Go as similar to instanceof in PHP or Java:  
    
  It checks if an error is a specific type (or wraps that type).  
    
  If it is, it gives you access to the actual typed error so you can inspect its fields.  
    
  Example analogy:  
    
  **PHP**


  ## *if ($error instanceof PathError) {*

  ##     *echo $error-\>path;*

  ## *}*

    
  **Go**


  ## *var pathErr \*os.PathError*

  ## *if errors.As(err, \&pathErr) {*

  ##     *fmt.Println("Path causing error:", pathErr.Path)*

  ## *}*

    
  Key point: errors.As is for type checking, while errors.Is is for value checking (specific error constants).


## runtime

The Go runtime package gives you access to functions that interact directly with Go’s runtime system — things like controlling how many CPU cores your program uses, inspecting the Go version, or getting memory statistics. These tools help you understand and tune how your program runs under the hood.

### Concurrent vs Parallel

Concurrent means tasks take turns running, but not necessarily at the same time.  
With concurrency, a single CPU rapidly switches between tasks, giving the illusion that they run simultaneously. It’s about managing multiple tasks at once, not executing them at the exact same time.

Parallel means tasks are actually running at the same time on multiple CPU cores.  
With parallelism, different cores execute different tasks simultaneously, giving true simultaneous execution.

**In short:**

* Concurrent: One CPU, tasks interleaved  
* Parallel: Multiple CPUs/cores, tasks truly simultaneous

  ### GOMAXPROCS

  Controls how many OS threads can execute Go code at the same time.  
    
  By default, Go uses all available CPU cores.  
  You can change this using runtime.GOMAXPROCS, which returns the previous value and lets you set a new one.


  ## *package main*

  ## 

  ## *import (*

  ##     *"fmt"*

  ##     *"runtime"*

  ## *)*

  ## 

  ## *func main() {*

  ##     *previous := runtime.GOMAXPROCS(1) // limits program to 1 CPU core*

  ##     *fmt.Println("Previous setting:", previous)*

  ## 

  ##     *fmt.Println("Running on:", runtime.GOMAXPROCS(0), "logical CPUs")*

  ## *}*

    
  Passing 0 returns the current value without changing it.  
  This is mainly useful when tuning performance or forcing predictable behavior in concurrent programs.


  ### NumCPU

  Returns the number of logical CPU cores available on the machine.

  ## 

  ## *cores := runtime.NumCPU()*

  ## *fmt.Println("This machine has", cores, "CPUs")*

  Good for printing system info or making decisions about concurrency.


  ### NumGoroutine

  Returns how many goroutines are currently running.


  ## *fmt.Println("Active goroutines:", runtime.NumGoroutine())*

  Useful for debugging concurrency issues or checking whether goroutines are being cleaned up properly.


  ### GoVersion

  Returns the version of Go used to compile the program.


  ## *fmt.Println("Go version:", runtime.GoVersion())*


  Helpful for logging, debugging, or ensuring certain runtime behaviors match your expectations.


  These runtime functions let you query system details and influence how your program interacts with the machine it’s running on, which can be especially important when working with goroutines, performance tuning, or debugging.

# 3rd Party Packages {#3rd-party-packages}

Third party packages can be found here:  
[https://pkg.go.dev/](https://pkg.go.dev/)

Each package site should contain a command that lets you install it.

Packages can only be installed in [modules](#module)

After we install the package a new section will be added to the module called “required” which will contain all of the package “dependencies” needed for the module.

It will also add a file called “[go.sum](#go.sum)”

Now we import the 3rd party package using its full name at the top of the package.  
Then we call the package using the last name in its url path

######  *package main*

###### 

###### *import (*

###### 	*"fmt"*

###### 

###### 	*"github.com/fatih/color"*

###### *)*

###### 

###### *func main() {*

###### 	*fmt.Println("I want a cup of coffee\!")*

###### 	*color.Blue("I want a cup of coffee")*

###### *}*

# Project Structure

## Module {#module}

A module contains packages which in turn contain files. Modules are recommended for most projects.

Module \> Packages \> Files

We can create a module with the [mod](#mod) command.

# Variables

## Typehinting (variables)

Once a value has been assigned to a var, that vars type is inferred by the type of the value it originally received.

###### *var number \= 1 // good*

###### *number \= false // rejected, must be a number.*

Alternatively we can declare a string type explicitly

###### *var coffeeName string \= “Espresso”*

Vars can also be assigned via [a shorthand using colon equals](#var-shorthand-\(:=\)).

###### *price := 2.5*

We can declare a variable outside of a function using standard notation, but shorthand can only be used when working inside of a function.

You can declare a variable without assigning it a value and it will assume the empty equivalent.  
For example if you declare a string but don’t assign it a value then the output will be a blank string, like this

###### *var wordOfTheDay string*

We’ve declared our variable will be a string, it has no value yet, as a result it will default to an empty string.

### Combined types {#combined-types}

If you happen to have two types that are different and you attempt to multiply or subtract one from the other for instance, the type that the result receives will not be inferred automatically.  
Instead you need to overtly declare or cast the types that you use as part of the equation, like so

###### *func main() {*

###### 	*price := 2.5*

###### 	*quantity := 5*

###### 

###### 	*total := price \* float64(quantity)*

###### 

###### 	*fmt.Println(total)*

###### *}*

In the above example you can see we have to overtly cast our quantity as a float so that the total that we receive at the end is a float as well.

NOTE: This does not apply to [const](#constants) values. A const that has an inferred type value can be added to a float and the net result will be a float without conversion.

###### *func main() {*

###### 	*// Untyped constant with integer value*

###### 	*const rewardPoints \= 10*

###### 

###### 	*fmt.Printf("Default type of rewardPoints is %T\\n", rewardPoints) // int*

###### 

###### 	*var totalRewardPoints float64 \= 150.3*

###### 

###### 	*// Adding untyped constant to a float64 \- valid, constant adapts*

###### 	*totalRewardPoints \= totalRewardPoints \+ rewardPoints*

###### 

###### 	*fmt.Printf("Updated loyalty points %.2f\\n", totalRewardPoints)*

###### *}*

NOTE: If we overtly declare the type of our const then this will not work.

## Casting Types

You can cast a type like so

###### *int(2.5)*

NOTE: The output of our example would be 2, so just the integer value, not the float, it does not round the value up or down, it just takes the integer from it.

## Types available

Available types include:

* string  
* float32  
* float64  
* int (int16, 32,64, 8\) etc

## var

This allows you to declare and assign variables.

###### *var example \= 1*

Any variables that you declared like this MUST be used or you will receive a compilation error.  
This is not true for [const](#constants) values though.  
You can declare const variables without using them and receive no errors as a result.const

### var shorthand (:=) {#var-shorthand-(:=)}

Vars can also be assigned via a shorthand using colon equals.

###### *price := 2.5*

## Multiple variable declaration

Go allows you to declare multiple variables at the same time.

###### *func main() {*

###### 	*var coffee, milk, sugar, isReady \= 1, 2, 1, true*

###### 

###### 	*// var coffee \= 1*

###### 	*// var milk \= 2*

###### 	*// var sugar \= 1*

###### 

###### 	*fmt.Println("Coffee:", coffee, "Milk:", milk, "Sugar:", sugar, "Ready:", isReady)*

###### *}*

###### 

Here we have declared four variables all on the same line.  
Notice that the fourth variable, isReady, is a boolean.  
Go is intelligent enough to infer the type of each variable it declares.

We can explicitly typehint these variables as well

###### *func main() {*

###### 	*var coffee, milk, sugar int \= 1, 2, 1, true*

###### 

###### 	*// var coffee \= 1*

###### 	*// var milk \= 2*

###### 	*// var sugar \= 1*

###### 

###### 	*fmt.Println("Coffee:", coffee, "Milk:", milk, "Sugar:", sugar, "Ready:", isReady)*

###### *}*

We can also declare multiple variables by using bracket notation

###### *var (*

###### 	*customerName string \= "Bogdan"*

###### 	*tableNum     int    \= 8*

###### 	*isReadyToPay bool   \= false*

###### *)*

###### 

###### *fmt.Printf("Customer %s at table %d is ready to pay: %t\\n", customerName, tableNum, isReadyToPay)*

###### 

###### *// No unused variables compilation error for const*

###### *const (*

###### 	*sizeSmall  \= "S"*

###### 	*sizeMedium \= "M"*

###### 	*sizeLarge  \= "L"*

###### *)*

## Constants {#constants}

We can declare a variable as a constant using “const”, which in turn means that it cannot be overridden.

Just like with typescript or JavaScript you should use const by default and only use var or alternatives if you intend to change the value later.

###### *func main() {*

###### 	*const shopName \= "Brew & Beans"*

###### 

###### 	*// shopName \= "Latte Palace" // ❌ can not assign*

###### 

###### 	*fmt.Println("Welcome to", shopName)*

###### *}*

###### 

Any values that are declared as a const must be assigned a value, you cannot leave them with null values.

######  *const example string \= "hello"*

NOTE: Const values are not declared with uppercase notation as they would be with PHP. Instead you can use camel case or pascal case notation instead. If you declare a variable using pascal case it will essentially become a global variable that can be used in other applications.

ALSO… [const values are not subject the to the same restrictions when combined with other types](#combined-types)

## Pascal Case values are Exported Automatically

If you declare a variable or a function using the Pascal case it will be automatically exported. This means that it will essentially be global and as a result it will be available to any other packages in that folder.

# Functions {#functions}

## Typehinting (functions)

Typehinting for functions in Go is very similar to that of typescript, so here we declare the type of the argument immediately after it.

###### *func updateTotalPoints(currentPoints int, newPoints int) int {*

###### 	*return currentPoints \+ newPoints*

###### *}*

Notice that the return type is different.  
Firstly there is no such thing as a “void” or “undefined” return type.  
If a function does not have a return value then we do not declare a return type  
Secondly, unlike typescript, we do not use a colon to delineate the return type, we just add it at the end of our function declaration, like the int above.

## Return Multiple Values

Functions can return multiple arguments

###### *func processPayment(orderTotal float64, tip float64, amountPaid float64) (float64, float64) {*

###### 	*totalAmountDue := orderTotal \+ tip*

###### 	*change := amountPaid \- totalAmountDue*

###### 	*return totalAmountDue, change*

###### *}*

###### 

###### *func main() {*

###### 	*totalCost, change := processPayment(6.50, 2.00, 10.00)* 

###### 	*fmt.Printf("Total cost (with tip): $%.2f\\n", totalCost)*

###### 	*fmt.Printf("Change returned to the customer: $%.2f\\n", change)*

###### *}*

* The return type hinting for multiple values is contained within brackets  
* We return the values directly, not within an array or similar.  
* We assign the values to variables directly

## Named Return Values

In Go you can use the return typehint to name the variables that you want to return from that function.  
That means that within the function you can declare variable x.  
In the return type state that you want to return the value of x.   
Then at the end you just simply return without passing an argument.  
Go is clever enough to recognise the name of the variable x from the return type and return the value of that specific variable.

###### *func estimateBrewTime(cupsQty int, secondsPerCup int) (totalTimeSeconds int, info string) {*

###### 	*totalTimeSeconds \= cupsQty \* secondsPerCup*

###### 	*info \= "Estimated total brew time:"*

###### 

###### 	*// naked return*

###### 	*return*

###### *}*

###### 

###### *func main() {*

###### 	*// 12 cups, 20 seconds per cup*

###### 	*// 12 \* 20 \= 240 seconds*

###### 	

###### 	*brewTime, info := estimateBrewTime(12, 20\)*

###### 	*fmt.Println(info, brewTime)*

###### *}*

In the above example we have a naked return type but, because we have specified the names of those variables that we want to return in the return type for the function, we are able to extract the values of specific variables from the function.

## Named Function Literals / Callback Functions

A Named Function Literal is essentially identical to a callback function in PHP

We declare an anonymous function and we assign it a variable. Then we call that variable and pass in any arguments that it needs to get the result.

###### *func main() {*

###### 	*taxRate := 0.10 // 10% tax*

###### 

###### 	*calculateTax := func(amount float64) float64 {*

###### 		*return amount \* taxRate*

###### 	*}*

###### 

###### 	*subtotal := 25.00*

###### 	*tax := calculateTax(subtotal)*

###### 	*total := subtotal \+ tax*

###### 	*fmt.Printf("Total amount to pay: $%.2f\\n", total)*

###### *}*

## Closures in Functions

In this example we have a function that returns a closure (callback) so we need to typehint that

######  *func createTemperatureAdjuster() (func(change float64) float64, float64) {*

###### 	*baseTemperature := 90.0*

###### 

###### 	*adjustTemperature := func(change float64) float64 {*

###### 		*baseTemperature \= baseTemperature \+ change*

###### 		*return baseTemperature*

###### 	*}*

###### 

###### 	*return adjustTemperature, baseTemperature*

###### *}*

Highlighted in yellow we have:

###### *func(change float64) float64*

This is because this function returns two values. The first is a closure (callback) and the second is a variable.

The closure, called “adjustTemperature” accepts a variable called “change” that is of type “float64” and the “adjustTemperature” function returns a float64 as well.  
So the type hint for our closure is

######  *func(change float64) float64*

That is the first typehint we return, the second is for “baseTemperature” which is just a float64 value.

### Closures retain external values in memory

Look at this code

###### *func createTemperatureAdjuster() (func(change float64) float64, float64) {*

###### 	*baseTemperature := 90.0*

###### 

###### 	*adjustTemperature := func(change float64) float64 {*

###### 		*baseTemperature \= baseTemperature \+ change*

###### 		*return baseTemperature*

###### 	*}*

###### 

###### 	*return adjustTemperature, baseTemperature*

###### *}*

###### 

###### *func main() {*

###### 	*adjustTemp, originalTemp := createTemperatureAdjuster()*

###### 	*fmt.Printf("Original temperature is %.1f\\n", originalTemp)*

###### 

###### 	*fmt.Printf("Adjusted Temp \+1.5: %.1f grad C\\n", adjustTemp(1.5))  // baseTemperature is changed*

###### 	*fmt.Printf("Adjusted Temp \-3.0: %.1f grad C\\n", adjustTemp(-3.0)) // baseTemperature is changed*

###### 	*fmt.Printf("Adjusted Temp \+5.0: %.1f grad C\\n", adjustTemp(5.0))  // baseTemperature is changed*

###### 	*fmt.Printf("Adjusted Temp \+5.0: %.1f grad C\\n", adjustTemp(5.0))  // baseTemperature is changed*

###### 

###### 	*fmt.Printf("Original temperature is %.1f\\n", originalTemp) // originalTemp is not changed*

###### 	*// 90.0*

###### 	*// \+1.5 \-\> 91.5*

###### 	*// \-3.0 \-\> 88.5*

###### 	*// \+5.0 \-\> 93.5*

###### *}*

This is important because here we have a function that returns a closure (callback).

adjustTemperature references the value baseTemperature which exists within its parent method createTemperatureAdjuster.

It holds the value of baseTemperature in memory as it was at the time that adjustTemperature was created or returned by createTemperatureAdjuster.

We have stored that returned callback against the variable “adjustTemp”. So adjustTemp now has its own internal value for baseTemperature and each time we call adjustTemp that baseTemperature value changes. So if we call it with 5 twice and the value of baseTemperature in adjustTemp was originally 90 that value will now be 100 (90 \+ 5 \+ 5).

It's as though baseTemperature is being passed into the callback by reference.

# Pointers {#pointers}

PHP does not copy an array in full unless that array is modified, so it automatically makes some effort to reduce memory overhead by limiting the amount of times it copies an array or similar.

Go does not. Go will instead copy an array in full unless expressly told otherwise. So to avoid copying an array in its entirety each time we need to use it we can use a pointer to directly change the variables value in memory.

###### *func applyDiscount(price \*float64, discountRate float64) {*

###### 	*\*price \= \*price \- (\*price \* discountRate)*

###### *}*

###### 

###### *func main() {*

###### 	*var coffeePrice float64 \= 5.00*

###### 	*var discount float64 \= 0.10*

###### 	*fmt.Printf("Basic coffee price: $%.2f\\n", coffeePrice) // 5.00*

###### 

###### 	*applyDiscount(\&coffeePrice, discount)*

###### 	*fmt.Printf("Price with discount: $%.2f\\n", coffeePrice) // 4.50*

###### 	*// 5.00*

###### 	*// 10%*

###### 	*// 5.00 \- 5.00 \* 0.10 \= 5.00 \- 0.50 \= 4.50*

###### *}*

Notice in the above example that we have [typehinted our function as a pointer](#typehinting-\(pointer\)) and as a result we have to pass in a variable using the [address-of operator](#address-of-operator) or “ampersand” symbol (&).

## Typehinting (pointer) {#typehinting-(pointer)}

Pointer typehints are precluded by a \* symbol like this

###### *var pointerToCoffeePrice \*float64 \= \&coffeePrice*

Here we have a pointer for a float64 value, notice the \* at the beginning.

## Address-of operator {#address-of-operator}

The & operator returns the memory location where a variable is stored.

NOTE: It only **shows** the location, it does not allow you to directly change it.

###### *n := 5*

###### *p := \&n*

###### *fmt.Println(p)*

Expected output:

###### *0xc0000140a8   // actual address will vary*

## Dereference operator

The \* operator accesses or modifies the value stored at the memory location pointed to by a pointer.

###### *n := 5*

###### *p := \&n*

###### *fmt.Println(\*p)  // read value at pointer*

###### *\*p \= 10          // modify value at pointer*

###### *fmt.Println(n)   // print modified value*

Expected output:

###### *5*

###### *10*

This shows & gives the address and \* lets you read or write the value at that address.

# Arrays

## Typehinting (array)

Arrays in Go are very similar to PHP except that their typehints are more explicit.  
You have to typehint the content and size of the array

###### *var coffeeSizes \[3\]string*

In this example we have overtly declared an array that will contain three elements and all of them will be strings.

## Array Literals

We can typehint and declare arrays all on one line

######  *coffeeTypes := \[3\]string{"Espresso", "Latte", "Cappuccino"}*

This differs from PHP, in that in PHP we would use square bracket notation, here we use braces.

NOTE:  
An array in PHP is flexible. You can add or remove values from the array and it will shrink or grow accordingly. In Go an Array is a fixed size. You can change the value of the entries within that array, but you cannot increase or decrease its size. For that you need to use a [slice](#slice-\(array-type\)).

## len (array)

This is similar to count in PHP, it returns the length of an array.

###### *coffeeTypes\[len(coffeeTypes)-1\] \= "Milk" // access last element in the array*

## cap (array)

Cap refers to the capacity of a slice or array each time that a slice is modified the length of that slice will change because it will either increase or decrease and then the capacity will change as well so if the length of a ray increases by one the capacity will increase by the value like 2 to allow for not only the addition of one more value but in anticipation of you adding potentially two more values if you then reach whatever the new capacity happens to be and you can continue to append to that slice then the capacity will update again with a new value  and will grow geometrically typically you can expect the capacity of the slice to double.

######  *menu := \[\]string{"Cake", "Pie"}*

###### 

###### *fmt.Println("Initial menu:", menu)*

###### *fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) // 2, 2*

###### 

###### *menu \= append(menu, "Donut")*

###### *fmt.Println("Menu after adding donut:", menu)*

###### *fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) // 3, 4*

In the above example you can see that we start with a slice that contains just two values and as a result has length of two and capacity of two but then we append one more value and the length increases to three to accommodate the three values it now contains while the capacity increases to four in anticipation of further values being added. 

## slice (function) {#slice-(function)}

This works in much the same way as slice in PHP where the first number provided is the first key we return and the last number represents the upper limit, so we return the key below that of the second value

###### *slice := dessertMenu\[1:3\] // elements with indexes 1, 2*

###### *fmt.Println("Slice of the Dessert Menu \[1:3\]:", slice)*

So to call slice we just have to pass the name of the array we want to access followed by square bracket notation of those values we want to return.

Here we are going to enter \[1:3\] which will return values 1 and 2 from the array dessertMenu.

### slice \[:2\] (function) {#slice-[:2]-(function)}

We can also get a slice by just declaring the right slice value instead

######  *slice := menu\[:2\]*

Once you have returned a slice from an array like this you can treat that slice, like a [slice](#slice-\(array-type\)), in that it can be appended or changed in size as much as you need.

### slice \[1\] (function)

Just like we can use the right value like [\[:2\]](#slice-[:2]-\(function\)) to cap the upper limit of a slice, we can instead just pass the minimum key that we want to use followed by a colon and we will receive a slice from that key until the end of the array or slice.

###### *slice := menu\[1:\]*

So in this example we are passing a value of 1 this means  we can expect to receive all values from position one onwards. 

###### *originalArray := \[3\]string{"Bill", "Bob", "Ben"}*

###### 

###### *arraySlice := originalArray\[1:\]*

###### *// arraySlice \= \[Bob Ben\]*

## Slice & Array Memory DependencyNOTE

When you create a slice, it has a length and a capacity. The capacity reserves space in memory. If you modify the slice without exceeding its capacity, your changes affect the original array because both use the same underlying memory. But if you append past the slice’s capacity, Go creates a new underlying array in a new memory location. After that, changes to the slice no longer affect the original array.

###### *originalArray := \[3\]int{1, 2, 3}*

###### *sharedSlice := originalArray\[:2\]     // len=2 cap=3, still backed by originalArray*

###### 

###### *sharedSlice\[0\] \= 10                  // originalArray becomes \[10 2 3\]*

###### 

###### *sharedSlice \= append(sharedSlice, 4\) // exceeds cap, new array created*

###### *sharedSlice\[1\] \= 20                  // originalArray stays \[10 2 3\]*

## Delete array values

In Go there is no function to delete or remove values from an array.

Instead you have to use [slice](#slice-\(function\)) to copy only those values that you want to keep from an existing array.

######  *func main() {*

######     *coffees := \[\]string{"Espresso", "Latte", "Mocha", "Cappuccino"}*

######     *// coffees \= \[Espresso Latte Mocha Cappuccino\]*

######     *// Output: Length: 4 Capacity: 4*

###### 

######     *// Remove "Latte" (index 1\)*

######     *indexToRemove := 1*

######     *coffees \= append(coffees\[:indexToRemove\], coffees\[indexToRemove+1:\]...)*

######     *// append combines slices.* 

######     *// \`...\` spreads the second slice elements individually.*

######     *// coffees \= \[Espresso Mocha Cappuccino\]*

######     *// Output: Length: 3 Capacity: 4 (capacity may remain the same)*

###### *}*

 Here is another more straight forward example

###### *arrayWithoutValue := append(originalArray\[:2\], originalArray\[4:\]...)*

###### *// arrayWithoutValue \= \[1 2 5 6 7 8 9 10\]* *// Capacity \= 10* *// arrayWithoutValue\[3\] \= 6*

Notice how the capacity of the array remains the same, because the slice has not “increased” the capacity of the underlying array, so we are still accessing the same array in memory.  
Also notice that once we remove an item from the array, the keys shuffle to close the gap, so here 6 moves from position 5 to position 3\.

## Spread operator

The spread operator behaves in much the same way as PHP, where it exposes the values contained within an array.

## For Loops

### For in Range

In order to create a for loop we need to use the range command as below.

###### *// Iterate over a slice*

###### *numbers := \[\]int{10, 20, 30}*

###### *for index, value := range numbers {*

######     *fmt.Println(index, value)*

###### *}*

###### *// Output:*

###### *// 0 10*

###### *// 1 20*

###### *// 2 30*

###### 

###### *// If you only need the value, use \_*

###### *for \_, value := range numbers {*

######     *fmt.Println(value)*

###### *}*

###### *// Output:*

###### *// 10*

###### *// 20*

###### *// 30*

Notice that in the second example, where we don’t want to use index, We just use an underscore value instead and that tells the system that we don't want to compile the index and that we just want to use value instead. 

### For I Loop

For I loops behave just as they normally do.

###### *for i := coffeeCups; i \>= 1; i-- { // 10, 9, ... 1*

###### 	*fmt.Printf("Preparing coffee cup \#%d\\n", i)*

###### *}*

### Infinite For Loop (Like a while) {#infinite-for-loop-(like-a-while)}

It is possible to create an infinite for loop that behaves just like a while loop in PHP

###### *tokens := 3*

###### 

###### *for tokens \> 0 {*

###### 	*fmt.Println("Making another cup of coffee...")*

###### 	*tokens--*

###### *}*

#### break and continue (for loops)

Just like normal for loops we can use break and continue

###### *for { // infinite loop*

###### 	*var order string*

###### 	*fmt.Print("Enter your coffee name (or type 'exit' to quit): ")*

###### 	*fmt.Scanln(\&order)*

###### 

###### 	*if order \== "exit" {*

###### 		*fmt.Println("Thank you for visiting Brew\&Beans. Good bye\!")*

###### 		*break*

###### 	*}*

###### 

###### 	*if order \== "" {*

###### 		*fmt.Println("Please enter a valid order")*

###### 		*continue*

###### 	*}*

###### 

###### 	*fmt.Println("Preparing your order...", order)*

###### *}*

**NOTE:** We use a pointer in the [Scanln](#scanln) method to ensure that we stop and wait to receive input before continuing

# Map (associative arrays, objects) {#map-(associative-arrays,-objects)}

A “map” is much like an “associative array” in php or an “object” in javascript.

###### *menu := map\[string\]float64{*

###### 	*"Espresso":   2.50,*

###### 	*"Latte":      3.75,*

###### 	*"Cappuccino": 3.50,*

###### 	*"Americano":  2.75,*

###### *}*

We call “map”, then the type of key we expect, which in this case is “string”.  
Then we declare the return type, which here is “float64”

## Exists (map)

“exists” in golang is most similar to key\_exists in php.  
It does not check the value.

###### *price, exists := menu\[drink\]*

###### 

###### *fmt.Println("Exists:", exists)*

###### 

###### *if exists {*

###### 	*fmt.Printf("%s costs $%.2f\\n", drink, price)*

###### *} else {*

###### 	*fmt.Printf("%s is not on the menu\\n", drink)*

###### *}*

**NOTE:** The name of the second variable here is “exists”, but the name does not matter. It could be “ready” or “ok”, it makes no difference.  
You can use this “exists” pattern for “maps”, “interfaces” and “channels”

## delete (map)

We can delete value from a “map” by using the delete command

###### *menu := map\[string\]float64{*

###### 	*"Latte":      3.75,*

###### 	*"Espresso":   2.50,*

###### *}*

###### 

###### *delete(menu, "Latte") // remove key with value*

## nil (map) {#nil-(map)}

## *nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type. (It can also be used with errors that have no value)*

The empty value for a map or slice is nil.

###### *var stock map\[string\]int // zero value is nil*

Here we have declared a map but we have not initialized it or added any values.  
In Go, a map is a reference type that refers to an internal data structure where its keys and values are stored.  
Since we have not called [make](#make-\(map\)) to initialize the map, the internal data structure has not been allocated, so the map variable does not reference any valid storage.  
As a result, the value of the map is nil.  
The same is true for [slices](#slice-\(array-type\)), [pointers](#pointers) and [functions](#functions)

If we try to assign a value to this map directly like this

###### *stock\["Espresso"\] \= 10 // Stock map is nil\!panic: assignment to entry in nil map*

Then it falls over. We cannot assign a new value to a map directly like this, because the maps value is nil.

## make (map) {#make-(map)}

The important thing to remember is that:

“Make allocates memory”

Without it we might declare a variable but that variable is just a pointer to a place in memory that does not exist. Make ensures that the area in memory exists so that we can check and interact with it. 

We can call “make” on a “map” in order to initialise it.  
[make can be called on a slice as well.](#make-can-be-called-on-a-slice-as-well.make-\(slice-function\))  
Once we have initialised a map we can then begin to add values to it directly.

######  *stock := make(map\[string\]int)*

###### 

###### *stock\["Espresso"\] \= 10*

###### *stock\["Latte"\] \= 25*

###### 

###### *fmt.Println("Products in stock:", stock)*

Here we have intialised our map, and as a result we can add values directly to it.

We can also provide an optional capacity hint when calling make on a map.

###### *stock := make(map\[string\]int, 100\) // capacity hint \= 100*

The hint tells Go roughly how many elements we expect to store. This allows the runtime to pre-allocate memory efficiently, reducing the need to resize the map as we add items, which can improve performance for larger maps.

**NOTE:** It’s a good idea to provide a capacity hint if you have a rough idea of how many elements the map will hold.

## Accessing Map Values

The key distinction when accessing maps rather than [structs](#struct) is that we can access values on a struct via dot notation, whereas a map requires a string value instead

######  *orderStruct := CoffeeOrder{*

###### 	*CustomerName: "Bogdan",*

###### 	*CoffeeType:   "Latte",*

###### 	*CoffeeSize:   "Medium",*

###### *}*

###### 

###### *// Map version*

###### *orderMap := map\[string\]string{*

###### 	*"CustomerName": "Bogdan",*

###### 	*"CoffeeType":   "Latte",*

###### 	*"CoffeeSize":   "Medium",*

###### *}*

###### 

###### *fmt.Println("---- Using Struct \----")*

###### *fmt.Println("Customer:", orderStruct.CustomerName)*

###### *fmt.Println("Coffee Type:", orderStruct.CoffeeType)*

###### *fmt.Println("Coffee Size:", orderStruct.CoffeeSize)*

###### 

###### *fmt.Println()*

###### 

###### *fmt.Println("---- Using Map \----")*

###### *fmt.Println("Customer:", orderMap\["CustomerName"\])*

###### *fmt.Println("Coffee Type:", orderMap\["CoffeeType"\])*

###### *fmt.Println("Coffee Size:", orderMap\["CoffeeSize"\])*

**NOTE:** Maps can be adjusted at runtime, whereas a struct cannot.

# Slice (array type) {#slice-(array-type)}

If you do not overtly specify the length of an array then it is a “slice” because it is “dynamically-sized” instead

A slice looks like this

###### *ratings := \[\]int{5, 4, 5, 5, 3}*

Note that there is no number within the square brackets, which is what determines this value to be a slice.

## make can be called on a slice as well.make (slice function) {#make-can-be-called-on-a-slice-as-well.make-(slice-function)}

Make is a [slice](#slice-\(array-type\)) specific function, it cannot be used with an array, but it can be used on a [map](#map-\(associative-arrays,-objects\)).

Make accepts three arguments. The first is the type, which in the example below is a slice of an array containing integer values (\[\]int).

###### *s := make(\[\]int, 3, 10\) // length 3, capacity 10*

The second value sets the original size of the slice. So here we will receive a slice that has keys 0, 1 and 2 (3 in total). These values will be blank by default unless we give them a value.

The third argument to make is a hard cap on the total number of elements that the array can contain. In this example the maximum value is ten. So at most our slice could look like this: \[0, 1, 2, 3, 4, 5, 6, 7, 8, 9\]

**NOTE:**  
The third argument to the slice is optional, you do not have to specify an upper limit to the size of the slice.

In Go, when you create a slice using make, such as s := make(\[\]int, 3, 10), it does not create a slice of ten elements. Instead, it creates a slice of length 3 with capacity 10\. The length is the number of elements the slice currently contains and that you can safely access using indices 0 to len-1. For example:

###### *s := make(\[\]int, 3, 10\) // length 3, capacity 10*

###### *s\[0\] \= 1*

###### *s\[1\] \= 2*

###### *s\[2\] \= 3*

###### *fmt.Println(s) // Output: \[1 2 3\]*

You cannot access s\[5\] yet, because the length is only 3 — doing so will cause a runtime panic. The capacity is the total number of elements the slice can grow to without allocating a new array. To safely use the extra capacity, you can use append, which grows the slice:

###### *s \= append(s, 4, 5, 6\)* 

###### *fmt.Println(s) // Output: \[1 2 3 4 5 6\]*

So, the capacity is like reserved space for growth, while the length is how many elements actually exist at that moment.

# If statements

If statements are handled much like python, without braces.

######  *if orderTotal \> 10 {*

###### 	*fmt.Println("You get a free cookie with your order\!")*

###### *}*

**NOTE:**  
It is not possible to compare two variables of two different types.  
So you can’t compare a float against an integer, it will throw an error.

## Declare variables within conditionals (Initialiser)

We can declare variables within conditionals using the semi-colon to separate values

###### *if points := 15; points \> 10 {*

###### 	*fmt.Println("You are eligible for coffee discount")*

###### *}*

###### 

###### *if fullAmount := getOrderWithTax(14.50, 0.1); fullAmount \> 15 {*

###### 	*fmt.Println("You can join coffee club")*

###### *}*

###### 

###### *totalLoyaltyPoints := 150*

###### 

###### *if totalLoyaltyPoints++; totalLoyaltyPoints \> 120 {*

###### 	*fmt.Println("Total loyalty points:", totalLoyaltyPoints)*

###### 	*fmt.Println("You can become Gold member")*

###### *}*

###### 

###### *if totalLoyaltyPoints \+= 10; totalLoyaltyPoints \> 120 {*

###### 	*fmt.Println("Total loyalty points:", totalLoyaltyPoints)*

###### 	*fmt.Println("You can become Gold member")*

###### *}*

Here we are declaring the “points” variable on the left of our conditional. Then on the right we are checking that value as part of our if.

The value on the left is called an “initialiser”.

# Switch

A “switch” in go behaves more like a “match” in PHP, in that it does not allow values to fall through.

###### *day := "Saturday"*

###### 

###### *switch day {*

###### *case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":*

###### 	*fmt.Println("Weekday special: Buy one get one 50% off")*

###### *case "Saturday":*

###### *case "Sunday":*

###### 	*fmt.Println("Weekend special: Free croissant with any coffee\!")*

###### *default:*

###### 	*fmt.Println("Unknown day")*

###### *}*

In the above example the value “Saturday” will not print any output, because none has been given. Whereas in PHP it would fall through to the Sunday clause instead.

# Struct {#struct}

A struct is similar to a DTO in that it groups related data fields under a single type. The struct definition itself does not allocate memory — it only describes the shape of the data. Memory is allocated only when you create a variable and declare its type as that struct.

Here we have a struct that declares the type for three fields:

###### *type ExampleDefinition struct {*

###### 	*one   int*

###### 	*two   string*

###### 	*three float64*

###### *}*

Notice there are no commas to separate each line.  
Each field has its own name and type.  
We declare our struct with the type keyword first, then the name, which will be public or private depending on whether it starts with a capital letter (PascalCase) or lowercase letter (camelCase).

To actually use this struct, you must create a variable of that struct type, which allocates memory for its fields.  
When printed, the struct’s fields will contain their zero values—in this case 0, "", and 0.0.

We can then interact with a struct instance much like we would with an object in JavaScript:

###### *var order CoffeeOrder*

###### 

###### *fmt.Println(order) // {   0}*

###### 

###### *order.CoffeeType \= "Cappuccino"*

###### *order.CoffeeSize \= "Large"*

###### *order.CustomerName \= "Bogdan"*

or

###### *order := CoffeeOrder{*

###### 	*CoffeeType: "Latte",*

###### 	*Size:       "Large",*

###### *}*

## Functions within Structs

### unsafe struct function {#unsafe-struct-function}

We can pass a function into a Struct.  
We do this with the word “func” followed by the typehint for that functions expected arguments.

######  *type CoffeeShop struct {*

###### 	*Name  string*

###### 	*Greet func(shop CoffeeShop)*

###### *}*

###### 

###### *func greetShop(shop CoffeeShop) {*

###### 	*fmt.Println("Welcome to the", shop.Name)*

###### *}*

###### 

###### *func main() {*

###### 	*myShop := CoffeeShop{*

###### 		*Name:  "Brew & Beans",*

###### 		*Greet: greetShop,*

###### 	*}*

###### 

###### 	*myShop.Greet(myShop)*

###### *}*

The above example is not great, because if we called shop.Greet within greetShop we would get infinite recursion.

The above approach is non-standard.

### safe struct function

Here is an example of what good looks like

######  *type CoffeeShop struct {*

###### 	*Name string*

###### *}*

###### 

###### *func (c CoffeeShop) Greet() {*

###### 	*fmt.Println("Welcome to the", c.Name)*

###### *}*

###### 

###### *func main() {*

###### 	*myShop := CoffeeShop{Name: "Brew & Beans"}*

###### 	*myShop.Greet()*

###### *}*

So here we create our [function](#functions) as we normally do, but instead of a function name, we add our struct name in brackets. Notice that this is where we inject our argument “c”, this is the value that our struct method will be instantiated with.

######  *func (c CoffeeShop)*

This tells golang that we want to assign a function to this struct. Here we assign the “Greet” function.

###### *func (c CoffeeShop) Greet()*

This is better than the [unsafe approach](#unsafe-struct-function) because the unsafe approach treats the function like a typehint, that can be overwritten with a function of the developers choosing. Whereas this safer approach overtly declares and assigns the function that want to have available as part of this struct.

### Struct Method Arguments

We can also set Greet to receive an argument as well. To do this, we define the parameter inside the parentheses after the method name, just like a regular function. For example, if we want to pass a string representing the item the customer ordered:

###### *func (c CoffeeShop) Greet(item string)*

Here’s the full picture:

###### *type CoffeeShop struct {*

######     *Name string*

###### *}*

###### 

###### *// Method with a parameter*

###### *func (c CoffeeShop) Greet(item string) {*

######     *fmt.Printf("Welcome to %s\! Enjoy your %s.\\n", c.Name, item)*

###### *}*

###### 

###### *func main() {*

######     *myShop := CoffeeShop{Name: "Brew & Beans"}*

######     *myShop.Greet("pickle")*

###### *}*

c CoffeeShop is the receiver, giving the method access to the struct’s fields.

item string is the argument passed when calling the method.

When we call myShop.Greet("pickle"), Go automatically passes myShop as c and "pickle" as item.

This is safer than storing a function in a struct field because:

The method is statically bound to the struct type.

It cannot be accidentally overwritten by another function.

The compiler enforces the method signature, so the parameter type and number of arguments are checked.

**NOTE:** There are no classes or class inheritance in Golang. Instead, we declare methods on the struct as above. This can be difficult to read, because it means that methods that belong to structs are not as immediately obvious as methods on classes.

The benefit of this is that Golang allows types to be quickly modified. 

# Adding Methods to existing Types

Here we have a variable with the string value “Espresso”. If we then decide that we want a function to be associated with this string we can create a new type, which in this case is “CoffeeType” and assign a method to it. That allows us to quickly specify that myCoffee has a “Describe” method.

##  *type CoffeeType string*

## 

## *func (coffee CoffeeType) Describe() {*

## 	*fmt.Println("This is delicious", coffee)*

## *}*

## 

## *func main() {*

## 	*var myCoffee CoffeeType \= "Espresso"*

## 	*myCoffee.Describe()*

## *}*

**NOTE:** It is not possible to add additional “fields” (which are similar properties in PHP. Fields are just keys on a struct) to a type like this. Here the methods relate explicitly to the CoffeeType string. This is useful for things such as validation.

# Interfaces

Interfaces allow us to define a set of behaviors that multiple types can implement, without needing to know the exact type. A type automatically satisfies an interface in Go if it has all the methods the interface requires. No explicit “implements” keyword is needed.

For example, we can create a CoffeeMachine interface that requires a Brew() method:

## *type CoffeeMachine interface {*

##     *Brew() string*

## *}*

## 

## *type CapsuleMachine struct {*

##     *Brand string*

## *}*

## 

## *func (capsule CapsuleMachine) Brew() string {*

##     *return fmt.Sprintf("%s has brewed one cup of coffee", capsule.Brand)*

## *}*

## 

## *func main() {*

##     ***// Here we create a variable of type CoffeeMachine***

##     ***// This is important because it tells Go that this variable***

##     ***// can hold any type that satisfies the CoffeeMachine interface***

##     *var coffeeMachine CoffeeMachine*

## 

##     ***// We assign a CapsuleMachine to the CoffeeMachine variable***

##     *coffeeMachine \= CapsuleMachine{Brand: "Nespresso"}*

## 

##     ***// We can now call Brew() on the interface variable***

##     *fmt.Println(coffeeMachine.Brew())*

## *}*

CapsuleMachine has a Brew() method, so it satisfies the CoffeeMachine interface.

The variable coffeeMachine is declared as type CoffeeMachine, which is important: it allows us to assign any type that implements the interface to it.

When we call coffeeMachine.Brew(), Go dynamically calls the correct Brew() method for the underlying type (CapsuleMachine).

This allows us to write code that works with any type that satisfies the interface, without knowing the exact type in advance. For example, later we could create an EspressoMachine type with its own Brew() method and assign it to the same variable.

**NOTE:** Interfaces in Go only define behavior, not data. The actual data comes from the type that implements the interface. This is different from class inheritance in other languages; instead of extending a class, Go encourages composition and explicit behavior contracts via interfaces.

## Empty interface

Go has a special type called the empty interface, written as interface{}.

## *var value interface{}*

An empty interface can hold any type of value, because every type satisfies zero methods. In Go 1.18 and later, any is a shorthand alias for interface{}:

## *var value any*

This is useful for:

* Functions that accept multiple types  
* Flexible containers like slices or maps of mixed types  
* Handling unknown input, such as from JSON  
    
  Summary: Empty interfaces (interface{}) or any let Go work with values of any type without knowing their exact type in advance.

# Passing an Interface as a Function Argument

We can also pass an interface into a function as an argument. This allows the function to accept any type that satisfies the interface, which makes our code flexible and easy to extend.

For example, we can write a function that expects a CoffeeMachine as its argument. This means the function does not care what the actual machine type is—as long as it has a Brew() method, it can be used.

## *func MakeCoffee(machine CoffeeMachine) {*

##     *fmt.Println(machine.Brew())*

## *}*

We can now call this function with any type that satisfies CoffeeMachine:

## *func main() {*

##     *nespresso := CapsuleMachine{Brand: "Nespresso"}*

##     *MakeCoffee(nespresso)*

## *}*

The argument machine CoffeeMachine tells Go that this function will accept any type that has a Brew() method.

Inside the function, we call machine.Brew() without knowing whether it is a CapsuleMachine, EspressoMachine, or some future type we might create later.

This is one of the main benefits of interfaces: we can write code that works with many different types, without needing to change the function when new types are added.

# Defer {#defer}

The defer keyword delays the execution of a function until the surrounding function finishes, even if the function returns early. Multiple deferred calls run in last-in, first-out (LIFO) order.  
(In PHP we might use “finally” instead)

## *package main*

## 

## *import "fmt"*

## 

## *func main() {*

##     *fmt.Println("Start")*

## 

##     *defer fmt.Println("First defer")*

##     *defer fmt.Println("Second defer")*

## 

##     *fmt.Println("Before return")*

##     *return // deferred calls still run here*

## *}*

Output:

## *Start*

## *Before return*

## *Second defer*

## *First defer*

Deferred calls run after the function returns.

The last deferred call runs first, showing LIFO order.

Commonly used for cleanup tasks, like closing files or unlocking resources.

### Defer use case

A good example of a defer use case is as follows

##  *file, \_ := os.Open("coffee.txt")*

## *defer file.Close()*

Here we are ensuring that our open file command closes correctly after everything else in the file has been handled appropriately.

### Defer Anonymous Functions

We can defer anonymous functions like this

## *defer func() {*

## 	*fmt.Println("Cleaning a coffee machine...")*

## 	*fmt.Println("Suspending coffee machine...")*

## *}()*

### Recover (defer) {#recover-(defer)}

recover is a built-in Go function that catches a panic so the program doesn’t crash. It can only be used inside a deferred function.

## *defer func() {*

##     *if r := recover(); r \!= nil {*

##         *fmt.Println("Machine error:", r)*

##     *}*

## *}()*

If a panic occurs (like dividing by zero), recover() returns the panic value, letting you handle the error.

If no panic occurs, recover() returns nil and nothing happens.

Multiple deferred functions can each call recover(), but only the one directly in the panicking function can catch it.

Use case: catching runtime errors safely and allowing the program to continue, similar to try/catch in other languages.

### Panic (defer) {#panic-(defer)}

A panic is a way to signal an unexpected runtime error that stops the current function immediately.  
(A panic in golang is similar to an exception in PHP)  
[Recover](#recover-\(defer\)) is used inside a defer to catch a panic and allow the program to continue running instead of crashing.

## *package main*

## 

## *import "fmt"*

## 

## *func DispenseCoffee(coffeeAmount, cups int) {*

##     *defer func() {*

##         *if r := recover(); r \!= nil {*

##             *fmt.Println("Machine error:", r)*

##         *}*

##     *}()*

## 

##     *if cups \== 0 {*

##         *panic("cannot divide by zero cups")*

##     *}*

## 

##     *fmt.Printf("Dispensing %d grams into %d cups\\n", coffeeAmount, cups)*

## *}*

## 

## *func main() {*

##     *DispenseCoffee(750, 3\) // normal*

##     *DispenseCoffee(340, 0\) // triggers panic, caught by recover*

##     *DispenseCoffee(500, 2\) // still runs*

## *}*

Output:

## *Dispensing 750 grams into 3 cups*

## *Machine error: cannot divide by zero cups*

## *Dispensing 500 grams into 2 cups*

Summary:

panic() signals an unexpected error and stops execution immediately.

recover() catches the panic inside a defer so the program can continue.

In this example, the division by zero triggers a panic, but recover() handles it, allowing the coffee machine to keep running.

# Handling Errors

In Go, errors are treated as regular values, not exceptions. This means functions that can fail usually return a second value of type error. You can then check this value to see if something went wrong and handle it appropriately.

The os package provides tools for working with files and the filesystem. Functions like os.Open return two values:

1. The file object (if opening succeeded)  
2. An error value (nil if no error occurred, otherwise contains details about what went wrong)  
     
   The errors package helps us check the type of error, for example using errors.Is to see if a file doesn’t exist.  
     
   Example

   ## *package main*

   ## 

   ## *import (*

   ##     *"errors"*

   ##     *"fmt"*

   ##     *"os"*

   ## *)*

   ## 

   ## *func main() {*

   ##     *file, err := os.Open("coffee\_orders.txt") // returns file and error*

   ## 

   ##     *if err \!= nil {*

   ##         *if errors.Is(err, os.ErrNotExist) {*

   ##             *fmt.Println("File doesn't exist")*

   ##         *} else {*

   ##             *fmt.Println("General file opening error:", err)*

   ##         *}*

   ##         *return*

   ##     *}*

   ## 

   ##     *fmt.Println("Successfully accessed file:", file.Name())*

   ## *}*

     
   Here os.Open returns two values: the file and an error.  
     
   We check err \!= nil to see if the operation failed.  
     
   errors.Is(err, os.ErrNotExist) lets us check the type of error specifically, rather than just printing it.  
     
   This pattern keeps Go code explicit, safe, and predictable, as errors are handled immediately rather than propagating silently.  
   

## Using error Directly

In Go, error is a built-in interface type. Any value that implements the Error() string method satisfies this type. You can create an error directly using fmt.Errorf or other constructors, and then check it like any other value.

Example:

## *var err error*

## *err \= fmt.Errorf("Some interesting coffee machine error")*

## 

## *if err \!= nil {*

##     *fmt.Println("Error occurred\!", err)*

## *} else {*

##     *fmt.Println("There is no error")*

## *}*

Here, err has the type error. Because fmt.Errorf returns a value of type error, we can store it in err and check it against nil. This allows us to handle the error immediately and explicitly.

## Using the inbuilt error interface

Here we have a variable called “CoffeeError” which is of type “string”.  
It has one method assigned to it called “Error”.  
By virtue of having one method called “Error” it matches [the built in error interface](https://pkg.go.dev/builtin@go1.25.5#error).  
As a result it now uses the “error” interface and can be treated like an error variable.  
We assign it to the “err” variable, check to make sure that variable is not [nil](#nil-\(map\)).  
It’s not nil, so it has a value, we print that value and because this is a known error interface when we print the variable on its own we get the error string within it.

## *type CoffeeError string*

## 

## *func (e CoffeeError) Error() string {*

## 	*return string(e)*

## *}*

## 

## *func main() {*

## 	*var err error*

## 	*err \= CoffeeError("No coffee beans loaded\!")*

## 

## 	*if err \!= nil {*

## 		*fmt.Println("Error:", err)*

## 	*}*

## *}*

## Using errors.New

The errors package provides a simple way to create basic error values using errors.New.  
This is useful when you want to return or check for a fixed, predefined error message.

errors.New returns a value of type error, which can be compared or returned just like any other error.

## *package main*

## 

## *import (*

##     *"errors"*

##     *"fmt"*

## *)*

## 

## *var ErrNoCoffee \= errors.New("no coffee available")*

## 

## *func main() {*

##     *err := ErrNoCoffee*

## 

##     *if err \!= nil {*

##         *fmt.Println("Error:", err)*

##     *}*

## *}*

Here, ErrNoCoffee is a predefined error created with errors.New.  
Because it’s a static value, it can be compared directly (e.g., if err \== ErrNoCoffee) and reused anywhere in your program.  
This keeps error declarations simple, clear, and consistent.

# WaitGroups

The sync.WaitGroup type in Go allows you to wait for a collection of goroutines to finish before continuing. It’s useful when you want to launch multiple concurrent tasks but need to make sure they all complete before moving on.

## Add

Tells the WaitGroup how many goroutines you are going to wait for.

## *var wg sync.WaitGroup*

## *wg.Add(3) // we plan to wait for 3 goroutines*

## Done

Signals that a goroutine has finished its work.  
It decrements the counter that was set by Add.

## *wg.Done() // call this at the end of a goroutine*

Note: It is a good practice to call wg.Done() using [defer](#defer) at the start of the goroutine. This ensures that Done is called even if the goroutine exits early due to an error.

## *go func() {*

##     *defer wg.Done() // guarantees Done is called*

##     *// work here*

## *}()*

## Wait

Blocks the program until the WaitGroup counter reaches zero.  
This is what makes the program wait for all goroutines to finish.

## *wg.Wait() // program pauses here until all Done calls are made*

## *package main*

## 

## *import (*

##     *"fmt"*

##     *"sync"*

##     *"time"*

## *)*

## 

## *func brewCoffee(id int, wg \*sync.WaitGroup) {*

##     *defer wg.Done() // ensures Done is called even if something goes wrong*

##     *fmt.Println("Starting coffee", id)*

##     *time.Sleep(time.Second) // simulate work*

##     *fmt.Println("Finished coffee", id)*

## *}*

## 

## *func main() {*

##     *var wg sync.WaitGroup*

## 

##     *wg.Add(3) // we will start 3 goroutines*

## 

##     *for i := 1; i \<= 3; i++ {*

##         *go brewCoffee(i, \&wg)*

##     *}*

## 

##     *wg.Wait() // wait for all goroutines to finish*

##     *fmt.Println("All coffee brewed\!")*

## *}*

**Output:**

* Starting coffee 1  
* Starting coffee 2  
* Starting coffee 3  
* Finished coffee 1  
* Finished coffee 2  
* Finished coffee 3  
* All coffee brewed\!  
    
    
  This pattern ensures that:  
    
  The program waits for all goroutines to finish (Wait)  
    
  Each goroutine signals completion (Done)  
    
  The counter is correctly set at the start (Add)  
    
  Using defer Done protects against early returns or errors inside the goroutine