package main

import "fmt"

type rect struct {
	Length, Width, Height int
}

func (r *rect) volume() (int, error) {

	var errs []string

	if r.Height < 0 {
		errs = append(errs, "height")
	}
	if r.Length < 0 {
		errs = append(errs, "length")
	}
	if r.Width < 0 {
		errs = append(errs, "width")
	}

	if len(errs) > 0 {
		var str string
		for i, e := range errs {
			if len(errs)-1 == i {
				// if it's the last item
				str = str + e
			} else {
				str = str + e + ", "
			}

		}
		return 0, fmt.Errorf("%s are negative", str)

	}

	return r.Height * r.Length * r.Width, nil
}

func main() {

	tests := []rect{
		{Length: 10, Width: 10, Height: 10},
		{Length: 10, Width: 10, Height: -10},
		{Length: 0, Width: 0, Height: 0},
		{Length: -200, Width: 10000, Height: -20000},
	}

	for i, rct := range tests {
		vol, err := rct.volume()
		if err != nil {
			fmt.Printf("test %d, error: %s\n", i, err.Error())
		} else {
			fmt.Printf("test %d, volume: %d\n", i, vol)
		}

	}

}
