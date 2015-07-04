# gogoa
Cocoa bindings for Go

Why?
---
Just for fun and curiosity. I was wondering how would it take to show a Cocoa Window through Go and here you Go :smile:

I don't know if this project will ever be useful and I'm rather sure this won't be useful anytime if it's just me coding it so I welcome contributions. They would make me feel I'm not alone :p

Getting started
---------------

The fastest way to try this project is running a Cocoa window. Actually this is the only thing you can do with it currently.

```go
package main

import (
  "github.com/alediaferia/gogoa"
)

func main() {
  app := gogoa.SharedApplication()
  window := gogoa.NewWindow(0, 0, 200, 200)
  window.SetTitle("Gogoga!")
  window.MakeKeyAndOrderFront()

  app.Run()
}
```

Screenshots
-----------
![Gogoa runs](/support/images/gogoa_run.png?raw=true "Gogoa runs!")

Contribute
----------
I would really love other people interested in this to help me build something useful.
Please, feel free to push whatever pull request comes to your mind. Also, I'd be happy to discuss the next steps with you.

License
-------
This codebase is currently licensed under MIT License you can find [here](LICENSE)

Copyright (c) Alessandro Diaferia <alediaferia@gmail.com>
