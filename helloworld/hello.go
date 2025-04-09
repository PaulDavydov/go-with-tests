package helloworld

const englishHelloPrefix = "Hello"

func makeMap() map[string]string {
	m := make(map[string]string)
	m["Spanish"] = "Hola"
	m["French"] = "Bonjour"
	m["Russian"] = "Privet"

	return m
}

func greetingPrefix(language string) (prefix string) {
	m := makeMap()

	switch language {
	case "":
		prefix = englishHelloPrefix
	default:
		prefix = m[language]
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + " " + name
}

//func main() {
//	name := "Eddie"
//	fmt.Println(Hello(name, "Spanish"))
//}
