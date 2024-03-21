package cmd

// cmd를 main에서 호출시킬텐데
type Cmd struct{}

// struct 안에 다른 요소들을 구성
func NewCmd() *Cmd {
	c := &Cmd{}

	return c
}

// newCmd를 main에서 받아서 서버 구동
