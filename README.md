

## Go Translation

### **Installation**
1. Clone this repository:
   ```sh
   git clone https://github.com/your-repo/go-translation.git
   cd go-translation
   ```
2. Initialize Go modules:
   ```sh
   go mod init go-translation
   ```

### **Usage**
Run the translation function manually:
```sh
go run translate.go
```
Or use it in your code:
```go
package main

import "fmt"

func main() {
    fmt.Println(translate("WOVN.io", "en", "MainActivity")) // Expected: "WOVN.io - MainScreen"
    fmt.Println(translate("WOVN.io", "en", ""))            // Expected: "WOVN.io"
}
```

### **Running Tests**
Run unit tests using `go test`:
```sh
go test
```

### **Project Structure**
```
├── data.json            # Translation dataset
├── translate.go         # Translation function
├── translate_test.go    # Go unit tests
├── go.mod               # Go module dependencies
└── README.md            # Documentation
```



