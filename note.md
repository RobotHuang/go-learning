# Golang study notes
## file
if you want to operate files on the computer, you need firstly import the package named "os"  
`import "os"`
### open&close file
```go
file, err := os.Open(path)
//"defer" could be omitted
defer file.Close()
```
### read file
#### file.Read
`func (f *File)Read(b []byte)(n int, err error)`
This function read bytes from file to the byte's slice, namely 'b' in above code.  
#### bufio
We also can use "bufio" to read files. "Bufio" encapsulate "File" and offer a series of API which is convenient for programmer to use it.
```go
file, _ := os.Open("/path")
reader := bufio.NewReader(file)
line, err := reader.ReadString('\n')
```
#### ioutil
"ioutil" is in package "io/ioutil" and it has a function called "ReadFile" which could read a complete content of file.
```go
content, err := ioutil.ReadFile("path")
```
### write file
#### os.OpenFile()
```go
func OpenFile(name string, flag int,perm FileMode)
```
flag has a number of values:
```
os.O_WRONLY
os.O_CREATE
os.O_RDONLY
os.O_RDWR	
os.O_TRUNC	
os.O_APPEND
```
#### file.Write([]byte) & file.WriteString(string)
#### bufio.NewWriter & writer.WriteString()
#### ioutil.WriteFile()
## package
1. All files in one directory belong only a package and the letters of package's name must all be lower case.   
2. An application must have a main package including a main function.  
3. The name of function, struct or interface must begin with the capital letter if you want to access it in other package.
## reflect
