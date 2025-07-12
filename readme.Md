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

## Context switching
- In our processor there are running multiple processes(tasks) if cpu complete some part of first process then go to the second process complete some part then go to third and more, finally back to the first process and start it task from the next part that is called context switching.

## Concurrency
- Cpu runs multiple task through context switching that at a time one task is running but our brain think that at a time running multiple task that is called concurrency.

## Process control block
- In context switching time,after completing first task some part cpu go to the second task but before moving it will store the first task last completed location into PCB so that when it will come again that task it will start the next part of that task.

## core discussion
- core i3 means in cpu there have three core and each core there have 2 virtual or logical cpu(processor)

## parallelism
- In our cpu there have multiple virtual cpu such as 10. When we run 2 program at a time then amoung 10 virtual cpu 2 virtual cpu run two task parallelly at a time that is called parallelism.

## different range of data type
- we can declare int8,int16,int32,int64 as our need in golong. Similary we can declare float32,float64 as our need in golong.

## defer
defer function execute at last when get return statement. defer follow LIFO operation