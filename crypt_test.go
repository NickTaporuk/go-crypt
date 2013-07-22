package crypt_test

import (
    "fmt"
    "."
)

func ExampleCrypt() {
    fmt.Println( crypt.Crypt("testtest", "es") )
    // Output:
    // esDRYJnY4VaGM
}
