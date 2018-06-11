package twofer

// ShareWith returns the appropriate sentence given the name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
