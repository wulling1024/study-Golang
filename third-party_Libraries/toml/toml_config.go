package toml

type tomlConfig struct {
	product  string
	tel      []int
	userDate user
	infoDate []info
}

type user struct{
	name string
	age  int
}

type info struct {
	ruanjianInfo ruanjian
	tumuInfo     tumu
	appleInfo    apple
	huaweiInfo   huawei
}

type ruanjian struct{
	name string
}
type tumu struct{
	name string
}
type apple struct{
	age int
}
type huawei struct{
	age int
}

