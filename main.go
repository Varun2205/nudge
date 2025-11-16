package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TodoItem struct {
	Text     string
	Complete bool
}

type TodoApp struct {
	todos  []TodoItem
	list   *widget.List
	data   binding.ExternalStringList
	entry  *widget.Entry
	window fyne.Window
}

func NewTodoApp() *TodoApp {
	app := &TodoApp{
		todos: make([]TodoItem, 0),
		data:  binding.BindStringList(&[]string{}),
	}

	return app
}

func (a *TodoApp) createUI() fyne.CanvasObject {
	// Title with Apple Notes-style typography
	title := widget.NewRichTextWithText("To-Do List")
	title.Segments[0].(*widget.TextSegment).Style = widget.RichTextStyle{
		TextStyle: fyne.TextStyle{Bold: true},
		SizeName:  theme.SizeNameHeadingText,
	}

	// Create input area
	a.entry = widget.NewEntry()
	a.entry.SetPlaceHolder("New task...")
	a.entry.OnSubmitted = func(s string) {
		if strings.TrimSpace(s) != "" {
			a.AddTodo(s)
			a.entry.SetText("")
		}
	}

	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		text := strings.TrimSpace(a.entry.Text)
		if text != "" {
			a.AddTodo(text)
			a.entry.SetText("")
		}
	})

	inputContainer := container.NewBorder(nil, nil, nil, addButton, a.entry)

	// Create todo list
	a.list = widget.NewList(
		func() int {
			return len(a.todos)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewCheck("", nil),
				widget.NewLabel("Template item"),
				layout.NewSpacer(),
				widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
			)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			box := item.(*fyne.Container)
			check := box.Objects[0].(*widget.Check)
			label := box.Objects[1].(*widget.Label)
			deleteBtn := box.Objects[3].(*widget.Button)

			// Update item content
			todo := &a.todos[id]
			label.SetText(todo.Text)

			// Style for completed items
			if todo.Complete {
				label.TextStyle = fyne.TextStyle{Italic: true}
			} else {
				label.TextStyle = fyne.TextStyle{}
			}

			// Update check state
			check.SetChecked(todo.Complete)
			check.OnChanged = func(checked bool) {
				a.todos[id].Complete = checked
				a.RefreshList()
			}

			// Update delete button action
			deleteBtn.OnTapped = func() {
				a.RemoveTodo(id)
			}
		},
	)

	// Apply Apple Notes-inspired styling
	a.list.Hide()

	// Main layout
	content := container.NewBorder(
		container.NewVBox(title, inputContainer), // Top: title and input
		nil,                                      // Bottom: nothing
		nil,                                      // Left: nothing
		nil,                                      // Right: nothing
		a.list,                                   // Center: todo list
	)

	return content
}

func (a *TodoApp) AddTodo(text string) {
	a.todos = append(a.todos, TodoItem{
		Text:     text,
		Complete: false,
	})
	a.RefreshList()
}

func (a *TodoApp) RemoveTodo(id int) {
	if id >= 0 && id < len(a.todos) {
		a.todos = append(a.todos[:id], a.todos[id+1:]...)
		a.RefreshList()
	}
}

func (a *TodoApp) RefreshList() {
	a.list.Refresh()
	a.list.Show()
}

func (a *TodoApp) LoadWindow(window fyne.Window) {
	a.window = window
	window.SetContent(a.createUI())

	// Set window size and style for widget-like appearance
	window.Resize(fyne.NewSize(400, 600))
	window.SetFixedSize(true)
}

func main() {
	// Create app and window
	myApp := app.NewWithID("com.todolist.widget")
	myApp.Settings().SetTheme(theme.LightTheme()) // Apple-like light theme

	window := myApp.NewWindow("To-Do List")

	todoApp := NewTodoApp()
	todoApp.LoadWindow(window)

	// Add some sample tasks
	todoApp.AddTodo("Buy groceries")
	todoApp.AddTodo("Call mom")
	todoApp.AddTodo("Finish project report")
	todoApp.AddTodo("Schedule dentist appointment")

	window.ShowAndRun()
}
