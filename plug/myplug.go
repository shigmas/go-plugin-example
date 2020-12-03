package plug

type (
	MyPlug interface {
		Init()
		DoPlug(arg string) string
	}
)
