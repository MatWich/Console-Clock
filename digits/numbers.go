package digits

type Placeholder [5]string

func GetDigits() [10]Placeholder {
	zero := Placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := Placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := Placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := Placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := Placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := Placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := Placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := Placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := Placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := Placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	Digits := [...]Placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	return Digits
}

func GetColon() Placeholder {
	return Placeholder{
		"   ",
		" ▒ ",
		"   ",
		" ▒ ",
		"   ",
	}
}
