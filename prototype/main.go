package main

import (
	"fmt"
	"strings"
	"time"
)

type Cloneable interface {
	Clone() Cloneable
}

type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

func (a *Address) Clone() *Address {
	return &Address{
		Street:  a.Street,
		City:    a.City,
		Country: a.Country,
		ZipCode: a.ZipCode,
	}
}

type Author struct {
	Name    string
	Email   string
	Address *Address
}

func (a *Author) Clone() *Author {
	return &Author{
		Name:    a.Name,
		Email:   a.Email,
		Address: a.Address.Clone(),
	}
}

type Document interface {
	Clone() Document
	GetInfo() string
	SetTitle(title string)
	SetContent(content string)
}

type BaseDocument struct {
	Title      string
	Content    string
	Author     *Author
	CreatedAt  time.Time
	ModifiedAt time.Time
	Version    int
}

type Resume struct {
	BaseDocument
	Experience []string
	Skills     []string
	Education  string
}

func (r *Resume) Clone() Document {
	experienceCopy := make([]string, len(r.Experience))
	copy(experienceCopy, r.Experience)
	
	skillsCopy := make([]string, len(r.Skills))
	copy(skillsCopy, r.Skills)
	
	return &Resume{
		BaseDocument: BaseDocument{
			Title:      r.Title,
			Content:    r.Content,
			Author:     r.Author.Clone(),
			CreatedAt:  r.CreatedAt,
			ModifiedAt: time.Now(),
			Version:    r.Version + 1,
		},
		Experience: experienceCopy,
		Skills:     skillsCopy,
		Education:  r.Education,
	}
}

func (r *Resume) GetInfo() string {
	var info []string
	info = append(info, fmt.Sprintf("=== RESUME ==="))
	info = append(info, fmt.Sprintf("Title: %s", r.Title))
	info = append(info, fmt.Sprintf("Author: %s (%s)", r.Author.Name, r.Author.Email))
	info = append(info, fmt.Sprintf("Location: %s, %s", r.Author.Address.City, r.Author.Address.Country))
	info = append(info, fmt.Sprintf("Education: %s", r.Education))
	info = append(info, fmt.Sprintf("Skills: %s", strings.Join(r.Skills, ", ")))
	info = append(info, fmt.Sprintf("Experience: %d positions", len(r.Experience)))
	info = append(info, fmt.Sprintf("Version: %d", r.Version))
	info = append(info, fmt.Sprintf("Created: %s", r.CreatedAt.Format("2006-01-02")))
	info = append(info, fmt.Sprintf("Modified: %s", r.ModifiedAt.Format("2006-01-02")))
	return strings.Join(info, "\n")
}

func (r *Resume) SetTitle(title string) {
	r.Title = title
}

func (r *Resume) SetContent(content string) {
	r.Content = content
}

type Report struct {
	BaseDocument
	ReportType string
	Department string
	Quarter    string
	Data       map[string]interface{}
}

func (r *Report) Clone() Document {
	dataCopy := make(map[string]interface{})
	for k, v := range r.Data {
		dataCopy[k] = v
	}
	
	return &Report{
		BaseDocument: BaseDocument{
			Title:      r.Title,
			Content:    r.Content,
			Author:     r.Author.Clone(),
			CreatedAt:  r.CreatedAt,
			ModifiedAt: time.Now(),
			Version:    r.Version + 1,
		},
		ReportType: r.ReportType,
		Department: r.Department,
		Quarter:    r.Quarter,
		Data:       dataCopy,
	}
}

func (r *Report) GetInfo() string {
	var info []string
	info = append(info, fmt.Sprintf("=== REPORT ==="))
	info = append(info, fmt.Sprintf("Title: %s", r.Title))
	info = append(info, fmt.Sprintf("Type: %s", r.ReportType))
	info = append(info, fmt.Sprintf("Department: %s", r.Department))
	info = append(info, fmt.Sprintf("Quarter: %s", r.Quarter))
	info = append(info, fmt.Sprintf("Author: %s (%s)", r.Author.Name, r.Author.Email))
	info = append(info, fmt.Sprintf("Data Points: %d", len(r.Data)))
	info = append(info, fmt.Sprintf("Version: %d", r.Version))
	info = append(info, fmt.Sprintf("Created: %s", r.CreatedAt.Format("2006-01-02")))
	info = append(info, fmt.Sprintf("Modified: %s", r.ModifiedAt.Format("2006-01-02")))
	return strings.Join(info, "\n")
}

func (r *Report) SetTitle(title string) {
	r.Title = title
}

func (r *Report) SetContent(content string) {
	r.Content = content
}

type DocumentRegistry struct {
	documents map[string]Document
}

func NewDocumentRegistry() *DocumentRegistry {
	return &DocumentRegistry{
		documents: make(map[string]Document),
	}
}

func (r *DocumentRegistry) Register(name string, doc Document) {
	r.documents[name] = doc
}

func (r *DocumentRegistry) Create(name string) Document {
	if prototype, exists := r.documents[name]; exists {
		return prototype.Clone()
	}
	return nil
}

func main() {
	fmt.Println("=== Prototype Pattern Demo ===\n")

	address := &Address{
		Street:  "123 Tech Street",
		City:    "San Francisco",
		Country: "USA",
		ZipCode: "94102",
	}

	author := &Author{
		Name:    "John Doe",
		Email:   "john.doe@example.com",
		Address: address,
	}

	fmt.Println("--- Creating Original Resume ---")
	originalResume := &Resume{
		BaseDocument: BaseDocument{
			Title:      "Software Engineer Resume",
			Content:    "Experienced software engineer...",
			Author:     author,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
			Version:    1,
		},
		Experience: []string{"Company A - 3 years", "Company B - 2 years"},
		Skills:     []string{"Go", "Python", "Docker", "Kubernetes"},
		Education:  "BS Computer Science",
	}
	fmt.Println(originalResume.GetInfo())

	fmt.Println("\n--- Cloning Resume and Modifying ---")
	clonedResume := originalResume.Clone().(*Resume)
	clonedResume.SetTitle("Senior Software Engineer Resume")
	clonedResume.Author.Name = "Jane Smith"
	clonedResume.Author.Email = "jane.smith@example.com"
	clonedResume.Skills = append(clonedResume.Skills, "Rust", "GraphQL")
	fmt.Println(clonedResume.GetInfo())

	fmt.Println("\n--- Original Resume (unchanged) ---")
	fmt.Println(originalResume.GetInfo())

	fmt.Println("\n--- Creating Original Report ---")
	originalReport := &Report{
		BaseDocument: BaseDocument{
			Title:      "Q1 Sales Report",
			Content:    "Sales performance for Q1...",
			Author:     author,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
			Version:    1,
		},
		ReportType: "Sales",
		Department: "Sales & Marketing",
		Quarter:    "Q1 2024",
		Data: map[string]interface{}{
			"revenue":   1500000,
			"growth":    15.5,
			"customers": 1250,
		},
	}
	fmt.Println(originalReport.GetInfo())

	fmt.Println("\n--- Using Document Registry (Prototype Manager) ---")
	registry := NewDocumentRegistry()
	registry.Register("resume-template", originalResume)
	registry.Register("report-template", originalReport)

	fmt.Println("Creating new resume from template...")
	newResume := registry.Create("resume-template").(*Resume)
	newResume.SetTitle("DevOps Engineer Resume")
	newResume.Author.Name = "Bob Johnson"
	newResume.Skills = []string{"AWS", "Terraform", "Jenkins", "Ansible"}
	fmt.Println(newResume.GetInfo())

	fmt.Println("\n--- Creating new report from template ---")
	newReport := registry.Create("report-template").(*Report)
	newReport.SetTitle("Q2 Sales Report")
	newReport.Quarter = "Q2 2024"
	newReport.Data["revenue"] = 1750000
	fmt.Println(newReport.GetInfo())
}
