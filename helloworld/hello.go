package helloworld

const (
	defaultName = "World"

	korean = "Korean"
	french = "French"

	englishHelloPrefix = "Hello, "
	koreanHelloPrefix  = "안녕, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case korean:
		prefix = koreanHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
