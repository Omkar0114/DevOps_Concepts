- import <module path for module_1>/abc 

- A module is a collection of Go packages stored in a file tree with a go.mod file at its root. The go.mod file defines the module's module path, which is also the import path used for the root directory.

To know more about the Modules, refer this page : - https://go.dev/blog/using-go-modules

Q: Given a variable named points, declared outside of all the functions in a file in a package utils inside the module module_1

ans: It can be anywhere inside utils package, not the rest of module_1
explaination: For points to be accessible outside the package, within module_1, or even outside the module, it should start with a capital letter, such as Points. Using capital letters makes Points an exported variable.

Q: What information does go.mod file include?
ans: Go version
     Dependencies needed
     Module replacement information
     All of the above