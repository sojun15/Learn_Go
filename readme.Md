## 1. why function

- For maintaining single responsibility principle we use function.
- For making simple our main class we use function.

## 2. Init function

- there have a special type of function which is init function that call auto-metically before main function. Besides we need't call it into main function.
- we can't pass anything as para-meter into init function
- we can't return anything from init function

## 3. Scope of variable

- variable scope mainly tree types local, global
- local scope occur mainly three types of blocks such as if-else, switch, function within those block.
- global scope occur when we declare it out of all function or other blocks.

## Receiver function
- receiver function is that type of function which declare for custom type object and we can call it through object_name.function_name
- For another type of custom object(struct) it will not working

## Pointer
- In pass by value there will be copy only values into another array or variable or memory address 
- In pass by address there will go only one address of the array or variable there save memory space

## slice 
- in slice there have three elements such as pointer, lenght and capacity. here pointer means starting address, lenght means how much element we can takes from array, capacity means from where we take slice that starting to ending elements