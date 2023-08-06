package depth

type Depth struct {
	coms chan int
}

func New() *Depth {
	return &Depth{
		coms: make(chan int),
	}
}
