package sender

type Submenu struct {
	Tlt     string
	Desc    string
	Options []MenuOption
}

func (s Submenu) Execute(sn Sender) error {
	return sn.Menu(s.Tlt, s.Desc, s.Options)
}
func (s Submenu) Title() string {
	return s.Tlt
}
