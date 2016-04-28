package animals

// Unexported type (lowercase first letter): Invisible from outside the package
type animal struct {
	name string		// Unexported field
	Age  int		// Exported field
}

// Exported type (uppercase first letter): Visible from outside the package
type Dog struct {
	animal		// Exports exported fields/methods of animal although animal itself is an unexported type
	BarkStrength int
}

// Exported method
func (d *Dog) Name() string {
	return d.name
}
