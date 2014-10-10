go-crypt - Traditional Crypt function implemented in Go.
========================================================

Example
-------

```go
import (
    "github.com/nyarlabo/go-crypt"
    "fmt"
)

func main() {
    fmt.Println( crypt.Crypt("testtest", "es") ); // esDRYJnY4VaGM
}
```

# Why I forked it?

Originalcode is written by iasija at 2009-12-08,
And orignal code is hosting on [code.google.com/p/go-crypt](https://code.google.com/p/go-crypt/).

But original code is not supported golang 1.1 or later.

So I fixed this issue, added documentation, and tests code.

NOTE: I could't find to iasija's contactt address.

# License and Copyright

  * Copyright of original code (C) 2009 isija All rights reserved. ([BSD-3-Clause](http://opensource.org/licenses/BSD-3-Clause))
  * Copyright of modified code (C) 2013-2014 Naoki OKAMURA (Nyarla) *nyarla[ at ]thotep.net* Some rights reserved. ([BSD-3-Clause](http://opensource.org/licenses/BSD-3-Clause))

