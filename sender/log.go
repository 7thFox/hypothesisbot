package sender

import (
	"fmt"
)

func (s Sender) Log(m string) {
	s.logger.Log(m)
}

func (s Sender) LogUser(m string) {
	s.logger.Log(fmt.Sprintf("%s: UID %s (%s)\n", m, s.user.ID, s.user.Username))
}
