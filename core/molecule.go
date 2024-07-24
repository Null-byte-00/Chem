package core

func ElementToNumber(element string) int {
	switch element {
	case "H":
		return 1
	case "C":
		return 6
	case "N":
		return 7
	case "O":
		return 8
	case "F":
		return 9
	case "P":
		return 15
	case "S":
		return 16
	case "Cl":
		return 17
	case "Br":
		return 35
	case "I":
		return 53
	default:
		return 0
	}
}


func NumberToElement(num int) string {
	switch num {
	case 1:
		return "H"
	case 6:
		return "C"
	case 7:
		return "N"
	case 8:
		return "O"
	case 9:
		return "F"
	case 15:
		return "P"
	case 16:
		return "S"
	case 17:
		return "Cl"
	case 35:
		return "Br"
	case 53:
		return "I"
	default:
		return "X"
	}
}

type Bond struct {
	Idx1 int
	Idx2 int
	Type int
}

type Molecule struct {
	Atoms []int
	Bonds []Bond
}
