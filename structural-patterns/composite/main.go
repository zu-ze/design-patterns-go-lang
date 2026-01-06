package main

import (
	"fmt"
	"strings"
)

type FileSystemComponent interface {
	GetName() string
	GetSize() int64
	Print(indent string)
}

type File struct {
	name string
	size int64
}

func NewFile(name string, size int64) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int64 {
	return f.size
}

func (f *File) Print(indent string) {
	fmt.Printf("%s📄 %s (%d bytes)\n", indent, f.name, f.size)
}

type Directory struct {
	name     string
	children []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: make([]FileSystemComponent, 0),
	}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int64 {
	var totalSize int64
	for _, child := range d.children {
		totalSize += child.GetSize()
	}
	return totalSize
}

func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

func (d *Directory) Remove(component FileSystemComponent) {
	for i, child := range d.children {
		if child.GetName() == component.GetName() {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Directory) Print(indent string) {
	fmt.Printf("%s📁 %s/ (%d bytes total)\n", indent, d.name, d.GetSize())
	for _, child := range d.children {
		child.Print(indent + "  ")
	}
}

type Employee interface {
	GetName() string
	GetPosition() string
	GetSalary() float64
	Print(indent string)
}

type Developer struct {
	name     string
	position string
	salary   float64
}

func NewDeveloper(name, position string, salary float64) *Developer {
	return &Developer{
		name:     name,
		position: position,
		salary:   salary,
	}
}

func (d *Developer) GetName() string {
	return d.name
}

func (d *Developer) GetPosition() string {
	return d.position
}

func (d *Developer) GetSalary() float64 {
	return d.salary
}

func (d *Developer) Print(indent string) {
	fmt.Printf("%s👨‍💻 %s - %s ($%.2f)\n", indent, d.name, d.position, d.salary)
}

type Designer struct {
	name     string
	position string
	salary   float64
}

func NewDesigner(name, position string, salary float64) *Designer {
	return &Designer{
		name:     name,
		position: position,
		salary:   salary,
	}
}

func (d *Designer) GetName() string {
	return d.name
}

func (d *Designer) GetPosition() string {
	return d.position
}

func (d *Designer) GetSalary() float64 {
	return d.salary
}

func (d *Designer) Print(indent string) {
	fmt.Printf("%s🎨 %s - %s ($%.2f)\n", indent, d.name, d.position, d.salary)
}

type Manager struct {
	name          string
	position      string
	salary        float64
	subordinates  []Employee
}

func NewManager(name, position string, salary float64) *Manager {
	return &Manager{
		name:         name,
		position:     position,
		salary:       salary,
		subordinates: make([]Employee, 0),
	}
}

func (m *Manager) GetName() string {
	return m.name
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) GetSalary() float64 {
	totalSalary := m.salary
	for _, subordinate := range m.subordinates {
		totalSalary += subordinate.GetSalary()
	}
	return totalSalary
}

func (m *Manager) AddSubordinate(employee Employee) {
	m.subordinates = append(m.subordinates, employee)
}

func (m *Manager) RemoveSubordinate(employee Employee) {
	for i, subordinate := range m.subordinates {
		if subordinate.GetName() == employee.GetName() {
			m.subordinates = append(m.subordinates[:i], m.subordinates[i+1:]...)
			break
		}
	}
}

func (m *Manager) Print(indent string) {
	fmt.Printf("%s👔 %s - %s ($%.2f, Team Budget: $%.2f)\n", 
		indent, m.name, m.position, m.salary, m.GetSalary())
	for _, subordinate := range m.subordinates {
		subordinate.Print(indent + "  ")
	}
}

type MenuComponent interface {
	GetName() string
	GetPrice() float64
	Print(indent string)
}

type MenuItem struct {
	name        string
	description string
	price       float64
}

func NewMenuItem(name, description string, price float64) *MenuItem {
	return &MenuItem{
		name:        name,
		description: description,
		price:       price,
	}
}

func (m *MenuItem) GetName() string {
	return m.name
}

func (m *MenuItem) GetPrice() float64 {
	return m.price
}

func (m *MenuItem) Print(indent string) {
	fmt.Printf("%s- %s: %s ($%.2f)\n", indent, m.name, m.description, m.price)
}

type Menu struct {
	name  string
	items []MenuComponent
}

func NewMenu(name string) *Menu {
	return &Menu{
		name:  name,
		items: make([]MenuComponent, 0),
	}
}

func (m *Menu) GetName() string {
	return m.name
}

func (m *Menu) GetPrice() float64 {
	var totalPrice float64
	for _, item := range m.items {
		totalPrice += item.GetPrice()
	}
	return totalPrice
}

func (m *Menu) Add(component MenuComponent) {
	m.items = append(m.items, component)
}

func (m *Menu) Print(indent string) {
	fmt.Printf("%s📋 %s Menu (Total: $%.2f)\n", indent, m.name, m.GetPrice())
	for _, item := range m.items {
		item.Print(indent + "  ")
	}
}

func printSeparator(title string) {
	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("%s\n", title)
	fmt.Printf("%s\n\n", strings.Repeat("=", 60))
}

func main() {
	fmt.Println("=== Composite Pattern Demo ===")

	printSeparator("Example 1: File System Structure")
	
	root := NewDirectory("root")
	
	home := NewDirectory("home")
	home.Add(NewFile("resume.pdf", 2048))
	home.Add(NewFile("photo.jpg", 4096))
	
	documents := NewDirectory("documents")
	documents.Add(NewFile("report.docx", 8192))
	documents.Add(NewFile("presentation.pptx", 16384))
	
	home.Add(documents)
	
	projects := NewDirectory("projects")
	projects.Add(NewFile("main.go", 1024))
	projects.Add(NewFile("README.md", 512))
	
	root.Add(home)
	root.Add(projects)
	root.Add(NewFile("config.txt", 256))
	
	root.Print("")
	fmt.Printf("\nTotal size: %d bytes\n", root.GetSize())

	printSeparator("Example 2: Company Organization Structure")
	
	ceo := NewManager("Alice Johnson", "CEO", 200000)
	
	cto := NewManager("Bob Smith", "CTO", 150000)
	cto.AddSubordinate(NewDeveloper("Charlie Brown", "Senior Developer", 100000))
	cto.AddSubordinate(NewDeveloper("Diana Prince", "Developer", 80000))
	cto.AddSubordinate(NewDeveloper("Eve Adams", "Junior Developer", 60000))
	
	cdo := NewManager("Frank Miller", "CDO", 140000)
	cdo.AddSubordinate(NewDesigner("Grace Lee", "Senior Designer", 90000))
	cdo.AddSubordinate(NewDesigner("Henry Wong", "UI/UX Designer", 75000))
	
	ceo.AddSubordinate(cto)
	ceo.AddSubordinate(cdo)
	
	ceo.Print("")
	fmt.Printf("\nTotal company payroll: $%.2f\n", ceo.GetSalary())

	printSeparator("Example 3: Restaurant Menu System")
	
	mainMenu := NewMenu("Main")
	
	breakfast := NewMenu("Breakfast")
	breakfast.Add(NewMenuItem("Pancakes", "Fluffy pancakes with syrup", 8.99))
	breakfast.Add(NewMenuItem("Omelette", "Three-egg omelette", 10.99))
	breakfast.Add(NewMenuItem("Coffee", "Fresh brewed coffee", 2.99))
	
	lunch := NewMenu("Lunch")
	lunch.Add(NewMenuItem("Burger", "Beef burger with fries", 12.99))
	lunch.Add(NewMenuItem("Salad", "Caesar salad", 9.99))
	lunch.Add(NewMenuItem("Pasta", "Spaghetti carbonara", 13.99))
	
	drinks := NewMenu("Drinks")
	drinks.Add(NewMenuItem("Soda", "Soft drink", 2.49))
	drinks.Add(NewMenuItem("Juice", "Fresh orange juice", 3.99))
	drinks.Add(NewMenuItem("Water", "Bottled water", 1.99))
	
	mainMenu.Add(breakfast)
	mainMenu.Add(lunch)
	mainMenu.Add(drinks)
	
	mainMenu.Print("")

	fmt.Println("\n--- Summary ---")
	fmt.Println("Composite pattern treats individual objects and compositions uniformly")
	fmt.Println("Tree structures can be built and traversed easily")
	fmt.Println("Operations apply to both leaves and composites")
}
