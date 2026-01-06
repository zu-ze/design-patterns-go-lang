package main

import "fmt"

type Button interface {
	Render() string
	OnClick() string
}

type Checkbox interface {
	Render() string
	Toggle() string
}

type Input interface {
	Render() string
	SetValue(value string) string
}

type LightButton struct{}

func (b *LightButton) Render() string {
	return "[Light Button] ☐ with white background and dark text"
}

func (b *LightButton) OnClick() string {
	return "[Light Button] Clicked with subtle shadow effect"
}

type LightCheckbox struct{}

func (c *LightCheckbox) Render() string {
	return "[Light Checkbox] ☐ with light gray border"
}

func (c *LightCheckbox) Toggle() string {
	return "[Light Checkbox] Toggled with smooth transition"
}

type LightInput struct{}

func (i *LightInput) Render() string {
	return "[Light Input] ___ with white background and thin border"
}

func (i *LightInput) SetValue(value string) string {
	return fmt.Sprintf("[Light Input] Value set to: %s (dark text on white)", value)
}

type DarkButton struct{}

func (b *DarkButton) Render() string {
	return "[Dark Button] ☐ with dark background and light text"
}

func (b *DarkButton) OnClick() string {
	return "[Dark Button] Clicked with neon glow effect"
}

type DarkCheckbox struct{}

func (c *DarkCheckbox) Render() string {
	return "[Dark Checkbox] ☐ with bright border on dark background"
}

func (c *DarkCheckbox) Toggle() string {
	return "[Dark Checkbox] Toggled with glowing animation"
}

type DarkInput struct{}

func (i *DarkInput) Render() string {
	return "[Dark Input] ___ with dark background and bright border"
}

func (i *DarkInput) SetValue(value string) string {
	return fmt.Sprintf("[Dark Input] Value set to: %s (light text on dark)", value)
}

type UIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
	CreateInput() Input
}

type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton() Button {
	return &LightButton{}
}

func (f *LightThemeFactory) CreateCheckbox() Checkbox {
	return &LightCheckbox{}
}

func (f *LightThemeFactory) CreateInput() Input {
	return &LightInput{}
}

type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton() Button {
	return &DarkButton{}
}

func (f *DarkThemeFactory) CreateCheckbox() Checkbox {
	return &DarkCheckbox{}
}

func (f *DarkThemeFactory) CreateInput() Input {
	return &DarkInput{}
}

type Application struct {
	factory UIFactory
}

func NewApplication(factory UIFactory) *Application {
	return &Application{factory: factory}
}

func (app *Application) RenderUI() {
	button := app.factory.CreateButton()
	checkbox := app.factory.CreateCheckbox()
	input := app.factory.CreateInput()

	fmt.Println(button.Render())
	fmt.Println(button.OnClick())
	fmt.Println()

	fmt.Println(checkbox.Render())
	fmt.Println(checkbox.Toggle())
	fmt.Println()

	fmt.Println(input.Render())
	fmt.Println(input.SetValue("Hello World"))
}

func main() {
	fmt.Println("=== Abstract Factory Pattern Demo ===\n")

	fmt.Println("--- Light Theme Application ---")
	lightFactory := &LightThemeFactory{}
	lightApp := NewApplication(lightFactory)
	lightApp.RenderUI()

	fmt.Println("\n--- Dark Theme Application ---")
	darkFactory := &DarkThemeFactory{}
	darkApp := NewApplication(darkFactory)
	darkApp.RenderUI()
}
