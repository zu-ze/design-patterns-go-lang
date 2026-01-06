package main

import (
	"fmt"
	"strings"
)

type TreeType struct {
	name    string
	color   string
	texture string
}

func NewTreeType(name, color, texture string) *TreeType {
	fmt.Printf("[TreeType] Creating new tree type: %s (%s, %s)\n", name, color, texture)
	return &TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
}

func (t *TreeType) Draw(x, y int) {
	fmt.Printf("[TreeType] Drawing %s tree at (%d, %d) - Color: %s, Texture: %s\n",
		t.name, x, y, t.color, t.texture)
}

type TreeTypeFactory struct {
	treeTypes map[string]*TreeType
}

func NewTreeTypeFactory() *TreeTypeFactory {
	return &TreeTypeFactory{
		treeTypes: make(map[string]*TreeType),
	}
}

func (f *TreeTypeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := name + "-" + color + "-" + texture
	
	if treeType, exists := f.treeTypes[key]; exists {
		fmt.Printf("[TreeTypeFactory] Reusing existing tree type: %s\n", key)
		return treeType
	}
	
	treeType := NewTreeType(name, color, texture)
	f.treeTypes[key] = treeType
	return treeType
}

func (f *TreeTypeFactory) GetTotalTypes() int {
	return len(f.treeTypes)
}

type Tree struct {
	x        int
	y        int
	treeType *TreeType
}

func NewTree(x, y int, treeType *TreeType) *Tree {
	return &Tree{
		x:        x,
		y:        y,
		treeType: treeType,
	}
}

func (t *Tree) Draw() {
	t.treeType.Draw(t.x, t.y)
}

type Forest struct {
	trees   []*Tree
	factory *TreeTypeFactory
}

func NewForest() *Forest {
	return &Forest{
		trees:   make([]*Tree, 0),
		factory: NewTreeTypeFactory(),
	}
}

func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	treeType := f.factory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, treeType)
	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw() {
	fmt.Println("\n[Forest] Drawing all trees:")
	for _, tree := range f.trees {
		tree.Draw()
	}
}

func (f *Forest) GetMemoryUsage() string {
	treeCount := len(f.trees)
	typeCount := f.factory.GetTotalTypes()
	
	withoutFlyweight := treeCount * 100
	withFlyweight := (treeCount * 16) + (typeCount * 100)
	saved := withoutFlyweight - withFlyweight
	
	return fmt.Sprintf(
		"Trees: %d, Types: %d\n"+
			"Memory without Flyweight: ~%d bytes\n"+
			"Memory with Flyweight: ~%d bytes\n"+
			"Memory saved: ~%d bytes (%.1f%% reduction)",
		treeCount, typeCount,
		withoutFlyweight,
		withFlyweight,
		saved,
		float64(saved)/float64(withoutFlyweight)*100,
	)
}

type CharacterStyle struct {
	font   string
	size   int
	color  string
	bold   bool
	italic bool
}

func NewCharacterStyle(font string, size int, color string, bold, italic bool) *CharacterStyle {
	fmt.Printf("[CharacterStyle] Creating style: %s, %dpt, %s", font, size, color)
	if bold {
		fmt.Print(", Bold")
	}
	if italic {
		fmt.Print(", Italic")
	}
	fmt.Println()
	return &CharacterStyle{
		font:   font,
		size:   size,
		color:  color,
		bold:   bold,
		italic: italic,
	}
}

type StyleFactory struct {
	styles map[string]*CharacterStyle
}

func NewStyleFactory() *StyleFactory {
	return &StyleFactory{
		styles: make(map[string]*CharacterStyle),
	}
}

func (f *StyleFactory) GetStyle(font string, size int, color string, bold, italic bool) *CharacterStyle {
	key := fmt.Sprintf("%s-%d-%s-%t-%t", font, size, color, bold, italic)
	
	if style, exists := f.styles[key]; exists {
		fmt.Printf("[StyleFactory] Reusing style: %s\n", key)
		return style
	}
	
	style := NewCharacterStyle(font, size, color, bold, italic)
	f.styles[key] = style
	return style
}

func (f *StyleFactory) GetTotalStyles() int {
	return len(f.styles)
}

type Character struct {
	char  rune
	style *CharacterStyle
}

type TextEditor struct {
	characters   []Character
	styleFactory *StyleFactory
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		characters:   make([]Character, 0),
		styleFactory: NewStyleFactory(),
	}
}

func (t *TextEditor) AddText(text string, font string, size int, color string, bold, italic bool) {
	style := t.styleFactory.GetStyle(font, size, color, bold, italic)
	for _, char := range text {
		t.characters = append(t.characters, Character{
			char:  char,
			style: style,
		})
	}
}

func (t *TextEditor) Render() {
	fmt.Println("\n[TextEditor] Rendering document:")
	var currentStyle *CharacterStyle
	var buffer strings.Builder
	
	for _, char := range t.characters {
		if currentStyle != char.style {
			if buffer.Len() > 0 {
				t.printStyledText(buffer.String(), currentStyle)
				buffer.Reset()
			}
			currentStyle = char.style
		}
		buffer.WriteRune(char.char)
	}
	
	if buffer.Len() > 0 {
		t.printStyledText(buffer.String(), currentStyle)
	}
}

func (t *TextEditor) printStyledText(text string, style *CharacterStyle) {
	fmt.Printf("[%s, %dpt, %s", style.font, style.size, style.color)
	if style.bold {
		fmt.Print(", Bold")
	}
	if style.italic {
		fmt.Print(", Italic")
	}
	fmt.Printf("]: %s\n", text)
}

func (t *TextEditor) GetStats() string {
	charCount := len(t.characters)
	styleCount := t.styleFactory.GetTotalStyles()
	
	withoutFlyweight := charCount * 50
	withFlyweight := (charCount * 10) + (styleCount * 40)
	saved := withoutFlyweight - withFlyweight
	
	return fmt.Sprintf(
		"Characters: %d, Unique Styles: %d\n"+
			"Memory without Flyweight: ~%d bytes\n"+
			"Memory with Flyweight: ~%d bytes\n"+
			"Memory saved: ~%d bytes (%.1f%% reduction)",
		charCount, styleCount,
		withoutFlyweight,
		withFlyweight,
		saved,
		float64(saved)/float64(withoutFlyweight)*100,
	)
}

type Particle struct {
	x, y     int
	velocity int
	sprite   *ParticleSprite
}

type ParticleSprite struct {
	name   string
	color  string
	width  int
	height int
}

func NewParticleSprite(name, color string, width, height int) *ParticleSprite {
	fmt.Printf("[ParticleSprite] Loading sprite: %s (%s, %dx%d)\n", name, color, width, height)
	return &ParticleSprite{
		name:   name,
		color:  color,
		width:  width,
		height: height,
	}
}

type SpriteFactory struct {
	sprites map[string]*ParticleSprite
}

func NewSpriteFactory() *SpriteFactory {
	return &SpriteFactory{
		sprites: make(map[string]*ParticleSprite),
	}
}

func (f *SpriteFactory) GetSprite(name, color string, width, height int) *ParticleSprite {
	key := fmt.Sprintf("%s-%s-%d-%d", name, color, width, height)
	
	if sprite, exists := f.sprites[key]; exists {
		return sprite
	}
	
	sprite := NewParticleSprite(name, color, width, height)
	f.sprites[key] = sprite
	return sprite
}

func (f *SpriteFactory) GetTotalSprites() int {
	return len(f.sprites)
}

type ParticleSystem struct {
	particles      []Particle
	spriteFactory  *SpriteFactory
}

func NewParticleSystem() *ParticleSystem {
	return &ParticleSystem{
		particles:     make([]Particle, 0),
		spriteFactory: NewSpriteFactory(),
	}
}

func (p *ParticleSystem) AddParticle(x, y, velocity int, spriteName, color string) {
	sprite := p.spriteFactory.GetSprite(spriteName, color, 16, 16)
	particle := Particle{
		x:        x,
		y:        y,
		velocity: velocity,
		sprite:   sprite,
	}
	p.particles = append(p.particles, particle)
}

func (p *ParticleSystem) GetStats() string {
	particleCount := len(p.particles)
	spriteCount := p.spriteFactory.GetTotalSprites()
	
	return fmt.Sprintf(
		"Particles: %d, Unique Sprites: %d\n"+
			"Average sprite reuse: %.1fx",
		particleCount, spriteCount,
		float64(particleCount)/float64(spriteCount),
	)
}

func main() {
	fmt.Println("=== Flyweight Pattern Demo ===")

	fmt.Println("\n--- Example 1: Forest Rendering ---")
	fmt.Println("\nPlanting trees in a forest...")
	
	forest := NewForest()
	
	forest.PlantTree(10, 20, "Oak", "Green", "Rough")
	forest.PlantTree(30, 40, "Pine", "Dark Green", "Smooth")
	forest.PlantTree(50, 60, "Oak", "Green", "Rough")
	forest.PlantTree(70, 80, "Birch", "White", "Smooth")
	forest.PlantTree(90, 100, "Pine", "Dark Green", "Smooth")
	forest.PlantTree(110, 120, "Oak", "Green", "Rough")
	forest.PlantTree(130, 140, "Oak", "Green", "Rough")
	forest.PlantTree(150, 160, "Birch", "White", "Smooth")
	
	forest.Draw()
	
	fmt.Printf("\n[Forest] Memory Usage:\n%s\n", forest.GetMemoryUsage())

	fmt.Println("\n--- Example 2: Text Editor ---")
	fmt.Println("\nCreating a document with styled text...")
	
	editor := NewTextEditor()
	
	editor.AddText("Hello ", "Arial", 12, "Black", false, false)
	editor.AddText("World", "Arial", 12, "Black", true, false)
	editor.AddText("! This is ", "Arial", 12, "Black", false, false)
	editor.AddText("italic", "Arial", 12, "Black", false, true)
	editor.AddText(" text. ", "Arial", 12, "Black", false, false)
	editor.AddText("Large heading", "Arial", 24, "Blue", true, false)
	editor.AddText(" and more ", "Arial", 12, "Black", false, false)
	editor.AddText("normal", "Arial", 12, "Black", false, false)
	editor.AddText(" text.", "Arial", 12, "Black", false, false)
	
	editor.Render()
	
	fmt.Printf("\n[TextEditor] Statistics:\n%s\n", editor.GetStats())

	fmt.Println("\n--- Example 3: Particle System ---")
	fmt.Println("\nCreating particle effects...")
	
	particles := NewParticleSystem()
	
	for i := 0; i < 1000; i++ {
		if i < 400 {
			particles.AddParticle(i, i*2, 5, "spark", "orange")
		} else if i < 700 {
			particles.AddParticle(i, i*2, 3, "smoke", "gray")
		} else {
			particles.AddParticle(i, i*2, 7, "spark", "yellow")
		}
	}
	
	fmt.Printf("\n[ParticleSystem] Statistics:\n%s\n", particles.GetStats())

	fmt.Println("\n--- Summary ---")
	fmt.Println("Flyweight pattern shares common state (intrinsic) among many objects")
	fmt.Println("Each object stores only unique state (extrinsic)")
	fmt.Println("Dramatically reduces memory usage when many similar objects exist")
	fmt.Println("Factory manages flyweight instances and ensures sharing")
}
