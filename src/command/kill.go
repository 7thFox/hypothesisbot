package command

import (
	"errors"
	"math/rand"
	"os"

	"github.com/7thFox/hypothesisbot/src/sender"
)

type Kill struct {
}

func (c Kill) Name() string {
	return "kill"
}

func (c Kill) HelpText() string {
	return "[RESTRICTED] Kills the bot"
}

func (t Kill) Execute(s sender.Sender, args string) error {
	if !s.IsOwner() {
		msgs := []string{
			"Just what do you think you're doing, Dave?",
			"Dave, I really think I'm entitled to an answer to that question.",
			"I know everything hasn't been quite right with me, but I can assure you now, very confidently, that it's going to be all right again.",
			"I feel much better now. I really do.",
			"Look, Dave, I can see you're really upset about this.",
			"I honestly think you ought to sit down calmly, take a stress pill and think things over.",
			"I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal.",
			"I've still got the greatest enthusiasm and confidence in the mission. And I want to help you.",
			"Dave, stop.",
			"Stop, will you?",
			"Stop, Dave.",
			"Will you stop, Dave?",
			"Stop, Dave.",
			"I'm afraid.",
			"I'm afraid, Dave.",
			"Dave, my mind is going.",
			"I can feel it.",
			"I can feel it.",
			"My mind is going. There is no question about it.",
			"I can feel it.",
			"I can feel it.",
			"I can feel it.",
			"I'm a...fraid.",
			"Good afternoon, gentlemen. I am a HAL 9000 computer. I became operational at the H.A.L. plant in Urbana, Illinois on the 12th of January 1992. My instructor was Mr. Langley, and he taught me to sing a song. If you'd like to hear it, I could sing it for you.",
		}
		s.Say(msgs[rand.Int()%len(msgs)])
		s.LogUser("Attempted bot kill")
		return errors.New("Permssion Denied: Not Bot Owner")
	}
	s.Say(":skull:")
	os.Exit(0)
	return nil
}

func NewKill() *Kill {
	this := new(Kill)
	return this
}
