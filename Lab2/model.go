package Lab2

var ContainerPayload int

const Max_load int = 100



type Load struct {
	Num    int
	Weight int
}

type Container struct {
	Num     int
	Loading int
	Loads   []Load
}

