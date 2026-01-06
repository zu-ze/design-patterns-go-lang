package main

import (
	"fmt"
	"strings"
)

type Computer struct {
	CPU          string
	RAM          int
	Storage      string
	GPU          string
	Motherboard  string
	PowerSupply  string
	CoolingType  string
	CaseType     string
	WiFi         bool
	Bluetooth    bool
	RGBLighting  bool
}

func (c *Computer) Specifications() string {
	var specs []string
	specs = append(specs, "=== Computer Specifications ===")
	specs = append(specs, fmt.Sprintf("CPU: %s", c.CPU))
	specs = append(specs, fmt.Sprintf("RAM: %d GB", c.RAM))
	specs = append(specs, fmt.Sprintf("Storage: %s", c.Storage))
	specs = append(specs, fmt.Sprintf("GPU: %s", c.GPU))
	specs = append(specs, fmt.Sprintf("Motherboard: %s", c.Motherboard))
	specs = append(specs, fmt.Sprintf("Power Supply: %s", c.PowerSupply))
	specs = append(specs, fmt.Sprintf("Cooling: %s", c.CoolingType))
	specs = append(specs, fmt.Sprintf("Case: %s", c.CaseType))
	
	features := []string{}
	if c.WiFi {
		features = append(features, "WiFi")
	}
	if c.Bluetooth {
		features = append(features, "Bluetooth")
	}
	if c.RGBLighting {
		features = append(features, "RGB Lighting")
	}
	
	if len(features) > 0 {
		specs = append(specs, fmt.Sprintf("Features: %s", strings.Join(features, ", ")))
	}
	
	return strings.Join(specs, "\n")
}

type ComputerBuilder struct {
	computer *Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{
		computer: &Computer{},
	}
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilder) SetStorage(storage string) *ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ComputerBuilder) SetMotherboard(motherboard string) *ComputerBuilder {
	b.computer.Motherboard = motherboard
	return b
}

func (b *ComputerBuilder) SetPowerSupply(powerSupply string) *ComputerBuilder {
	b.computer.PowerSupply = powerSupply
	return b
}

func (b *ComputerBuilder) SetCoolingType(coolingType string) *ComputerBuilder {
	b.computer.CoolingType = coolingType
	return b
}

func (b *ComputerBuilder) SetCaseType(caseType string) *ComputerBuilder {
	b.computer.CaseType = caseType
	return b
}

func (b *ComputerBuilder) AddWiFi() *ComputerBuilder {
	b.computer.WiFi = true
	return b
}

func (b *ComputerBuilder) AddBluetooth() *ComputerBuilder {
	b.computer.Bluetooth = true
	return b
}

func (b *ComputerBuilder) AddRGBLighting() *ComputerBuilder {
	b.computer.RGBLighting = true
	return b
}

func (b *ComputerBuilder) Build() *Computer {
	return b.computer
}

type Director struct{}

func (d *Director) BuildGamingPC(builder *ComputerBuilder) *Computer {
	return builder.
		SetCPU("Intel Core i9-13900K").
		SetRAM(32).
		SetStorage("2TB NVMe SSD").
		SetGPU("NVIDIA RTX 4090").
		SetMotherboard("ASUS ROG Maximus Z790").
		SetPowerSupply("1000W 80+ Gold").
		SetCoolingType("Liquid Cooling 360mm").
		SetCaseType("Full Tower RGB").
		AddWiFi().
		AddBluetooth().
		AddRGBLighting().
		Build()
}

func (d *Director) BuildOfficePC(builder *ComputerBuilder) *Computer {
	return builder.
		SetCPU("Intel Core i5-13400").
		SetRAM(16).
		SetStorage("512GB SSD").
		SetGPU("Integrated Graphics").
		SetMotherboard("ASUS Prime B660").
		SetPowerSupply("450W 80+ Bronze").
		SetCoolingType("Air Cooling").
		SetCaseType("Mid Tower").
		AddWiFi().
		Build()
}

func (d *Director) BuildWorkstationPC(builder *ComputerBuilder) *Computer {
	return builder.
		SetCPU("AMD Ryzen 9 7950X").
		SetRAM(64).
		SetStorage("4TB NVMe SSD + 8TB HDD").
		SetGPU("NVIDIA RTX 4080").
		SetMotherboard("ASUS Pro WS X670E").
		SetPowerSupply("850W 80+ Platinum").
		SetCoolingType("Liquid Cooling 280mm").
		SetCaseType("Mid Tower").
		AddWiFi().
		AddBluetooth().
		Build()
}

func main() {
	fmt.Println("=== Builder Pattern Demo ===\n")

	director := &Director{}

	fmt.Println("--- Building Gaming PC (using Director) ---")
	gamingPC := director.BuildGamingPC(NewComputerBuilder())
	fmt.Println(gamingPC.Specifications())

	fmt.Println("\n--- Building Office PC (using Director) ---")
	officePC := director.BuildOfficePC(NewComputerBuilder())
	fmt.Println(officePC.Specifications())

	fmt.Println("\n--- Building Workstation PC (using Director) ---")
	workstationPC := director.BuildWorkstationPC(NewComputerBuilder())
	fmt.Println(workstationPC.Specifications())

	fmt.Println("\n--- Building Custom PC (without Director) ---")
	customPC := NewComputerBuilder().
		SetCPU("AMD Ryzen 7 7800X3D").
		SetRAM(32).
		SetStorage("1TB NVMe SSD").
		SetGPU("AMD Radeon RX 7900 XTX").
		SetMotherboard("MSI MAG X670E").
		SetPowerSupply("750W 80+ Gold").
		SetCoolingType("Air Cooling - Noctua NH-D15").
		SetCaseType("Mid Tower").
		AddWiFi().
		AddRGBLighting().
		Build()
	fmt.Println(customPC.Specifications())
}
