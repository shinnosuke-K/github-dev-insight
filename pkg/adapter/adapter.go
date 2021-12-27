package adapter

type Adapter struct {
	GitHub GitHub
}

func (a *Adapter) Github() GitHub {
	return a.GitHub
}
