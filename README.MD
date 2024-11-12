# Push Swap

A highly efficient number sorting algorithm implementation in Go that sorts a stack of integers using a limited set of operations. This project implements a sorting algorithm using two stacks with a specific set of allowed operations.

## 🎯 Overview

### The project consists of two programs:

* #### Push Swap: A program that calculates and displays the smallest sequence of operations that sorts a random list of integers using two stacks (Stack A and Stack B).

* #### Checker: A program that takes Push Swap's output and verifies whether the operations actually sort the stack correctly.

### Allowed Operations

- `sa` (swap a): Swap the first 2 elements at the top of stack A
- `sb` (swap b): Swap the first 2 elements at the top of stack B
- `ss`: Execute `sa` and `sb` simultaneously
- `pa` (push a): Take the first element at the top of B and push it to A
- `pb` (push b): Take the first element at the top of A and push it to B
- `ra` (rotate a): Shift up all elements of stack A by 1
- `rb` (rotate b): Shift up all elements of stack B by 1
- `rr`: Execute `ra` and `rb` simultaneously
- `rra` (reverse rotate a): Shift down all elements of stack A by 1
- `rrb` (reverse rotate b): Shift down all elements of stack B by 1
- `rrr`: Execute `rra` and `rrb` simultaneously

## 🚀 Features

- ⚡ Efficient sorting algorithm with optimization for different stack sizes
- 🔄 Sorts integers using only two stacks (A & B)
- 🔍 Input validation and error handling
- 📊 Advanced sorting algorithm for larger sets using median and range-based approach
- ✅ Built-in operation validator
- 🧪 Zero external dependencies
- 💻 Pure Go implementation

## 📋 Requirements

- Go 1.18 or higher
- No external dependencies required

## 💻 Installation

```bash
# Clone the repository
git clone <repository-url>
cd push-swap

# Build the project
./build.sh
```

## 🎮 Usage

### Push Swap
```bash
# Run push-swap with a list of integers
./push-swap "4 67 3 87 23"
```
#### The program will output the sequence of operations needed to sort the stack.
### Checker
```bash
# Method 1: Manual checking
./checker "4 67 3 87 23"

# Then type operations one by one (press Enter after each)
# Press Ctrl+D when done

### Method 2: Pipe push-swap output to checker
./push-swap "4 67 3 87 23" | ./checker "4 67 3 87 23"
```

#### The checker will output:

- `OK` if the sequence of operations correctly sorts the stack
- `KO` if the stack is not properly sorted
- `Error` if an invalid operation is encountered

### Input Format

- Numbers must be integers
- No duplicates are allowed
- Numbers can be positive or negative
- Numbers should be space-separated within quotes

## 🏗 Project Structure
* ### Main Programs :
- `PushSwap/push-swap.go`: Main entry point of the push-swap program
- `Chechek-Programe/checker.go`: checker Program to verify the sorting operations
* ### Core Components :
- `Algorithm/operations.go`: Contains input validation and basic sorting operations
- `Algorithm/variables.go`: Defines types, variables, and utility functions
- `Algorithm/part_One.go`: Implements the first phase of the sorting algorithm
- `Algorithm/part_Two.go`: Implements the second phase of the sorting algorithm
- `Algorithm/moves.go`: The set of operations used for sorting

## 🔍 Algorithm Overview

The sorting algorithm is divided into two main parts:

### Small Sets (2-6 numbers)
- Specialized sorting algorithms for optimal performance
- Direct approach with minimal operations

### Large Sets
1. **Part One**:
   - Divides numbers into ranges
   - Uses median for efficient stack management
   - Pushes numbers to stack B in a strategic order

2. **Part Two**:
   - Sorts numbers back to stack A
   - Uses queue system for optimization
   - Implements rotation strategies based on number positions

## ⚡ Performance

The algorithm is optimized for different input sizes:
- 2 numbers: Maximum 1 operation
- 3 numbers: Maximum 3 operations
- 4 numbers: Maximum 5 operations
- 5 numbers: Maximum 9 operations
- 6 numbers: Maximum 13 operations
- 100 number: Maximum under 700 operation
- Larger sets: Optimized using range-based approach

## 🛠 Error Handling

The program includes comprehensive error handling for:
- Invalid input formats
- Duplicate numbers
- Non-integer inputs
- Already sorted sequences
