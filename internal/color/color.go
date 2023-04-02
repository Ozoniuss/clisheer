package color

type Color string

const (
	Reset  Color = "\x1b[0m"
	Red    Color = "\x1b[31m"
	Green  Color = "\x1b[32m"
	Yellow Color = "\x1b[33m"
	Blue   Color = "\x1b[34m"
	Purple Color = "\x1b[35m"
	Cyan   Color = "\x1b[36m"
	Gray   Color = "\x1b[37m"
	White  Color = "\x1b[97m"
)
