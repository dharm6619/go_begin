package conditionals

func GetFullName(name string) string {
	fullName := ""
	switch name {
	case "go":
		fullName = "GoLang"
	case "k8s":
		fullName = "Kubernetes"
	case "cli":
		fullName = "Command Line Interface"
	default:
		fullName = "No Match Found"
	}
	return fullName
}
