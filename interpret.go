package dorklang

type InterpretCodeOptions struct {
	DebugMode bool
	SkipClean bool
}

var (
	InterpretCodeDefaultOptions = InterpretCodeOptions{
		DebugMode: false,
		SkipClean: false,
	}
)
