/**
- A package consists of one or more .go source files in a single directory
that define what the package does.
- Each source file begins with a package declaration, here package main,
that states which package the file belongs to;
- A list of the imported packages follows
- Then the declarat ions of the program
- Package main is special. It defines a standalone executable program, not a library
- Within the package main, the function main is also special:
  it’s where execution of the program begins
 - To tell the compiler what packages are needed by this source file, we use the
import declaration that follows the package declaration.
- Import declarations must follow the package declaration.
- After that, a program consists of the declarations of:
	functions, variables, constants, and types
    (func, var, const, and type keywords)
**/

package main

import "fmt"

func main() {
	fmt.Println("生日快乐, Ricky")
	/** happy birthday, Ricky **/
}