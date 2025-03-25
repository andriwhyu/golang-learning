package model

type Saver interface {
	Save() error
}

type Outputable interface {
	Saver
	Show()
}
