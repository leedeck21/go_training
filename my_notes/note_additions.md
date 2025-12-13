# Note Additions

This file contains only the new sections that are present in the updated `Golang Notes.md` but not in the original `original_golang_notes.md`.

---

## _p := \&n_

## _fmt.Println(\*p) // read value at pointer_

## _\*p = 10 // modify value at pointer_

## _fmt.Println(n) // print modified value_

Expected output:

## _5_

## _10_

This shows & gives the address and \* lets you read or write the value at that address.

## Slice & Array Memory DependencyNOTE

When you create a slice, it has a length and a capacity. The capacity reserves space in memory. If you modify the slice without exceeding its capacity, your changes affect the original array because both use the same underlying memory. But if you append past the slice’s capacity, Go creates a new underlying array in a new memory location. After that, changes to the slice no longer affect the original array.

## _originalArray := [3]int{1, 2, 3}_

## _sharedSlice := originalArray[:2] // len=2 cap=3, still backed by originalArray_

## _sharedSlice[0] = 10 // originalArray becomes [10 2 3]_

## _sharedSlice = append(sharedSlice, 4) // exceeds cap, new array created_

## _sharedSlice[1] = 20 // originalArray stays [10 2 3]_

## Spread operator

The spread operator behaves in much the same way as PHP, where it exposes the values contained within an array.

### For I Loop

For I loops behave just as they normally do.

## _for i := coffeeCups; i >= 1; i-- { // 10, 9, ... 1_

## _fmt.Printf("Preparing coffee cup #%d\n", i)_

## _}_

### Infinite For Loop (Like a while) {#infinite-for-loop-(like-a-while)}

It is possible to create an infinite for loop that behaves just like a while loop in PHP

## _tokens := 3_

## _for tokens > 0 {_

## _fmt.Println("Making another cup of coffee...")_

## _tokens--_

## _}_

#### break and continue (for loops)

Just like normal for loops we can use break and continue

## _for { // infinite loop_

## _var order string_

## _fmt.Print("Enter your coffee name (or type 'exit' to quit): ")_

## _fmt.Scanln(&order)_

## _if order == "exit" {_

## _fmt.Println("Thank you for visiting Brew&Beans. Good bye!")_

## _break_

## _}_

## _if order == "" {_

## _fmt.Println("Please enter a valid order")_

## _continue_

## _}_

## _fmt.Println("Preparing your order...", order)_

## _}_

**NOTE:** We use a pointer in the [Scanln](#scanln) method to ensure that we stop and wait to receive input before continuing

## _make can be called on a slice as well.make (slice function) {#make-can-be-called-on-a-slice-as-well.make-(slice-function)}_

Make is a [slice](<#slice-(array-type)>) specific function, it cannot be used with an array, but it can be used on a [map](<#map-(associative-arrays,-objects)>).

Make accepts three arguments. The first is the type, which in the example below is a slice of an array containing integer values (\[\]int).

## _s := make([]int, 3, 10) // length 3, capacity 10_

The second value sets the original size of the slice. So here we will receive a slice that has keys 0, 1 and 2 (3 in total). These values will be blank by default unless we give them a value.

The third argument to make is a hard cap on the total number of elements that the array can contain. In this example the maximum value is ten. So at most our slice could look like this: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

**NOTE:**  
The third argument to the slice is optional, you do not have to specify an upper limit to the size of the slice.

In Go, when you create a slice using make, such as s := make([]int, 3, 10), it does not create a slice of ten elements. Instead, it creates a slice of length 3 with capacity 10. The length is the number of elements the slice currently contains and that you can safely access using indices 0 to len-1. For example:

## _s := make([]int, 3, 10) // length 3, capacity 10_

## _s[0] = 1_

## _s[1] = 2_

## _s[2] = 3_

## _fmt.Println(s) // Output: [1 2 3]_

You cannot access s[5] yet, because the length is only 3 — doing so will cause a runtime panic. The capacity is the total number of elements the slice can grow to without allocating a new array. To safely use the extra capacity, you can use append, which grows the slice:

## _s = append(s, 4, 5, 6)_

## _fmt.Println(s) // Output: [1 2 3 4 5 6]_

So, the capacity is like reserved space for growth, while the length is how many elements actually exist at that moment.
