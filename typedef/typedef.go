package typedef

type Permission int

const (
	AdminPerm       Permission = 0
	UnprotectedPerm            = -1
	PublicPerm                 = -2
)
