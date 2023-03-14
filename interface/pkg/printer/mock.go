package printer

type TestPrinter struct {
	value string
}

func NewTestPrinter() *TestPrinter {
	return &TestPrinter{}
}

func (p *TestPrinter) Println(a string) error {
	p.value = a
	return nil
}

func (p *TestPrinter) Value() string {
	return p.value
}
