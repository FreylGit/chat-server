package codes

type Code uint32

const (
	Ok               Code = 0
	Canceled         Code = 1
	Unknown          Code = 2
	InvalidArgument  Code = 3
	DeadLineExceeded Code = 4
	NotFound         Code = 5
	AlreadyExist     Code = 6
)
