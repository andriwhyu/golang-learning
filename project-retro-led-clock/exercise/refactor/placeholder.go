package main

type placeholder [5]string

var separator = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var blankSeparator = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	"   ",
}

var digits_0 = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var digits_1 = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var digits_2 = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var digits_3 = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var digits_4 = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var digits_5 = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var digits_6 = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var digits_7 = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var digits_8 = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var digits_9 = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var digits = [...]placeholder{
	digits_0,
	digits_1,
	digits_2,
	digits_3,
	digits_4,
	digits_5,
	digits_6,
	digits_7,
	digits_8,
	digits_9,
}
