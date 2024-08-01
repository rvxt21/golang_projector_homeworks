package enteties

type Question struct {
	Text    string
	Answers map[int]string
	Correct int
}
