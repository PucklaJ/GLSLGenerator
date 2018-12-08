package glslgen

func (this *VertexGenerator) SetVersion(version string) *VertexGenerator {
	this.Version = version
	return this
}

func (this *VertexGenerator) AddOutput(output Variable) *VertexGenerator {
	this.Outputs = append(this.Outputs, output)
	return this
}

func (this *VertexGenerator) AddOutputs(outputs []Variable) *VertexGenerator {
	this.Outputs = append(this.Outputs, outputs...)
	return this
}

func (this *VertexGenerator) AddMakro(makro Makro) *VertexGenerator {
	this.Makros = append(this.Makros, makro)
	return this
}

func (this *VertexGenerator) AddMakros(makros []Makro) *VertexGenerator {
	this.Makros = append(this.Makros, makros...)
	return this
}

func (this *VertexGenerator) AddGlobal(global Variable) *VertexGenerator {
	this.Globals = append(this.Globals, global)
	return this
}

func (this *VertexGenerator) AddGlobals(globals []Variable) *VertexGenerator {
	this.Globals = append(this.Globals, globals...)
	return this
}

func (this *VertexGenerator) AddModule(module Module) *VertexGenerator {
	this.Modules = append(this.Modules, module)
	return this
}

func (this *VertexGenerator) AddAttribute(attribute Variable) *VertexGenerator {
	this.Attributes = append(this.Attributes, attribute)
	return this
}

func (this *VertexGenerator) AddAttributes(attributes []Variable) *VertexGenerator {
	this.Attributes = append(this.Attributes, attributes...)
	return this
}
