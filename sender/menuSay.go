package sender

type MenuSay struct {
	Tlt  string
	Desc string
	Txt  string
}

func (s MenuSay) Execute(sn Sender, parentMsgID string) error {
	return sn.SayReplace(s.Txt, parentMsgID)
}
func (s MenuSay) Title() string {
	return s.Tlt
}
