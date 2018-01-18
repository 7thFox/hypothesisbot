package sender

type Submenu struct {
	Tlt     string
	Desc    string
	Options []MenuOption
}

func (s Submenu) Execute(sn Sender, parentMsgID string) error {
	return sn.MenuReplace(s.Tlt, s.Desc, s.Options, parentMsgID)
}
func (s Submenu) Title() string {
	return s.Tlt
}
