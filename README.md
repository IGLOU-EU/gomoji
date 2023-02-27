# 🥰 Gomoji
[![Pull request are open](https://img.shields.io/badge/Pull_request-Open-green.svg?style=flat-square)](https://github.com/IGLOU-EU/gomoji/fork)
[![License: GPL 3.0](https://img.shields.io/badge/License-GPL_3.0_or_later-blue.svg?style=flat-square)](https://www.gnu.org/licenses/gpl-3.0.html)

The reason of this project, is that I like emoji and I wanted to make a small project to facilitate their use in my programs in Go.   
Fell free to request functionality and/or made PR 💞

## 🧠 Latest Release
**Date**: `0000-00-00`   
**Release ID**: `00000000000000000000000000000`

**Number of ...**
- **Emoji**: `0`
- **Keywords**: `0`
- **Category**: `0`

## 💻 Usage
To begin with the latest version you can import it in your project:
```go
import "github.com/IGLOU-EU/gomoji"
```
Don't forget to run `go mod tidy` to update the project and download it.

You can use it directly with global var, from global slice, or with exported function.   
**Exemple**:
```go
package main

import (
    "fmt"
    "github.com/IGLOU-EU/gomoji"
)

func main() {
    fmt.Println("I like", gomoji.Get("penguin").Picto);

    fmt.Println("There is multiple birds into emoji");
    for _, p := range gomoji.ByKeyword("bird") {
        fmt.Println("-", p.Picto);
    }

    fmt.Printf("Here, this emoji is new, I would like to know its information :\n%#v", gomoji.Info("🧌"));
)
```