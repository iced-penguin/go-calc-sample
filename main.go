package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	var formula string
	var isResult bool
	var canEnterOperator bool

	init := func() {
		formula = ""
		isResult = false
		canEnterOperator = true
	}

	app := app.New()
	window := app.NewWindow("Calc")

	init()

	label := widget.NewLabel(formula)
	label.Alignment = fyne.TextAlignTrailing

	pushNum := func(n int) {
		// 数式の先頭が0とならないように
		if formula == "" && n == 0 {
			return
		}
		// 計算直後に数字を入力すると入力前にクリア
		if isResult {
			init()
		}
		formula += strconv.Itoa(n)
		label.SetText(formula)
	}

	pushOperator := func(s string) {
		if formula == "" || !canEnterOperator {
			return
		}
		// 計算結果に続けて演算子を入力
		if isResult {
			isResult = false
		}
		formula += " " + s + " "
		canEnterOperator = false
		label.SetText(formula)
	}

	pushEnter := func() {
		s := strings.Split(formula, " ")
		if len(s) != 3 {
			return
		}

		num1, err := strconv.Atoi(s[0])
		if err != nil {
			dialog.ShowError(err, window)
		}
		num2, err := strconv.Atoi(s[2])
		if err != nil {
			dialog.ShowError(err, window)
		}
		op := s[1]

		result, err := calc(num1, num2, op)
		if err != nil {
			dialog.ShowError(err, window)
		}

		formula = strconv.Itoa(result)
		isResult = true
		canEnterOperator = true
		label.SetText(formula)
	}

	clear := func() {
		init()
		label.SetText(formula)
	}

	window.SetContent(
		widget.NewVBox(
			label,
			widget.NewHBox(
				fyne.NewContainerWithLayout(
					layout.NewFixedGridLayout(fyne.NewSize(200, 200)),
					fyne.NewContainerWithLayout(
						layout.NewGridLayout(3),
						widget.NewButton("7", func() { pushNum(7) }),
						widget.NewButton("8", func() { pushNum(8) }),
						widget.NewButton("9", func() { pushNum(9) }),
						widget.NewButton("4", func() { pushNum(4) }),
						widget.NewButton("5", func() { pushNum(5) }),
						widget.NewButton("6", func() { pushNum(6) }),
						widget.NewButton("1", func() { pushNum(1) }),
						widget.NewButton("2", func() { pushNum(2) }),
						widget.NewButton("3", func() { pushNum(3) }),
						widget.NewButton("0", func() { pushNum(0) }),
						layout.NewSpacer(),
						widget.NewButton("=", pushEnter),
					),
				),
				widget.NewVBox(
					widget.NewButton("CL", clear),
					widget.NewButton("/", func() { pushOperator("/") }),
					widget.NewButton("*", func() { pushOperator("*") }),
					widget.NewButton("+", func() { pushOperator("+") }),
					widget.NewButton("-", func() { pushOperator("-") }),
				),
			),
		),
	)

	app.Settings().SetTheme(theme.DarkTheme())
	window.ShowAndRun()
}

func calc(num1 int, num2 int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		return result, fmt.Errorf("Invalid operand %s", op)
	}
	return result, nil
}
