package typedef

type Permission int

const (
	SpravcePerm     Permission = 4
	TechnikPerm                = 3
	DispecerPerm               = 2
	RidicPerm                  = 1
	AdminPerm                  = 0
	UnprotectedPerm            = -1
	PublicPerm                 = -2
)
