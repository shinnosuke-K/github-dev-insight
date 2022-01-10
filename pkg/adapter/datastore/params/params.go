package params

type GetByStatusWithPaging struct {
	Status bool
	Limit  int
	Offset int
}
