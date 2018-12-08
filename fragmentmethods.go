package glslgen

func (this *FragmentGenerator) SetVersion(version string) *FragmentGenerator {
	this.Version = version
	return this
}

func (this *FragmentGenerator) AddOutput(output Variable) *FragmentGenerator {
	this.Outputs = append(this.Outputs, output)
	return this
}

func (this *FragmentGenerator) AddOutputs(outputs []Variable) *FragmentGenerator {
	this.Outputs = append(this.Outputs, outputs...)
	return this
}

func (this *FragmentGenerator) AddMakro(makro Makro) *FragmentGenerator {
	this.Makros = append(this.Makros, makro)
	return this
}

func (this *FragmentGenerator) AddMakros(makros []Makro) *FragmentGenerator {
	this.Makros = append(this.Makros, makros...)
	return this
}

func (this *FragmentGenerator) AddGlobal(global Variable) *FragmentGenerator {
	this.Globals = append(this.Globals, global)
	return this
}

func (this *FragmentGenerator) AddGlobals(globals []Variable) *FragmentGenerator {
	this.Globals = append(this.Globals, globals...)
	return this
}

func (this *FragmentGenerator) AddModule(module Module) *FragmentGenerator {
	this.Modules = append(this.Modules, module)
	return this
}

func (this *FragmentGenerator) AddInput(input Variable) *FragmentGenerator {
	this.Inputs = append(this.Inputs, input)
	return this
}

func (this *FragmentGenerator) AddInputs(inputs []Variable) *FragmentGenerator {
	this.Inputs = append(this.Inputs, inputs...)
	return this
}
