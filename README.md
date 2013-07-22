go-crypt
========

Traditional Crypt function implemented in Go.

# Example

```go
import (
    "github.com/nyarlabo/go-crypt"
    "fmt"
)

func main() {
    fmt.Println( crypt.Crypt("testtest", "es") ); // esDRYJnY4VaGM
}
```

# Why do I have to re-package?

Original code is written by iasija at 08 Dec, 2009,
and hosting on [code.google.com/p/go-crypt](https://code.google.com/p/go-crypt/).

But this code does not work current golang 1.1,
So, I fix for supports current golang 1.1,
and added to some documentation.

And, I could not found to iasijya's contacts address.

# License and Copyright

* Copyright 2009 iasijya All rights reserved.
* Original code is hosting on [code.google.com/p/go-crypt](https://code.google.com/p/go-crypt/).
* This library under the [New BSD License](http://opensource.org/licenses/BSD-3-Clause).
  * (License is specified in the google code page.)
* Some Fix by Naoki OKAMURA (Nyarla) *nyarla[ at ]thotep.net*

