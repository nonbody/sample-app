package print

type PrintFunc func(string) string

func NewPrintFunc() PrintFunc {
	return func(s string) string {
		return "Print " + s
	}
}
