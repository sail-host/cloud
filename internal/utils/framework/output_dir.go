package framework

func OutputDir(framework string) string {
	// TODO: Add more frameworks

	switch framework {
	case "nextjs":
		return "out"
	case "nuxt":
		return ".output"
	case "react":
		return "build"
	case "vue":
		return "dist"
	case "vite":
		return "dist"
	default:
		return "dist"
	}
}
