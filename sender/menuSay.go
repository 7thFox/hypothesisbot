package sender

type MenuSay struct {
	Tlt  string
	Desc string
	Txt  string
}

func (s MenuSay) Execute(sn Sender) error {
	return sn.Say(s.Txt)
}
func (s MenuSay) Title() string {
	return s.Tlt
}
