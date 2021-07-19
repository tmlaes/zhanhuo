package main

type IBingYing interface {
	Create(id int)

	ShengJi(id int)

	ZhaoMu(id int)
}

type BingYing struct {
	id int
	p  Point
}

func (b *BingYing) Create(id int) {

}
