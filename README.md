# go-calc-sample

Go-calc-sample is a simple calculator.
I made it to learn how to develop Go GUI application.

## Demo

![](https://i.gyazo.com/971d59a92d81d82fbe13a11d21e84c66.gif)

## Requirement

- Go 1.16+
- Fyne 1.4.3+

```
go get fyne.io/fyne
```

## Installation

Clone this repository:

```
git clone https://github.com/icedpenguin0504/go-calc-sample.git
```

Run the program using `go run`:

```
go run main.go
```

Or you can package the application. First, install `fyne` command:

```
go get fyne.io/fyne/cmd/fyne
```

Then execute the `fyne` package command:

```
fyne package -icon icon.png
```

You can use any image you like instead of `icon.png`.

The above commands are the same for Windows, Mac OS, and Linux.
They will create the appropriate application formats for each platform.