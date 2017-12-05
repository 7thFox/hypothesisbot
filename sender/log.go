package sender

import (
	"fmt"
)

func (s Sender) Log(m string) {
	fmt.Println(m)
}

func (s Sender) LogUser(m string) {
	fmt.Printf("%s: UID %s (%s)\n", m, s.user.ID, s.user.Username)
}
