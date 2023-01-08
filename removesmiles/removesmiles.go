package removesmiles

import "fmt"

func RemoveSmiles(s string) string {
	j := 0
	ls := len(s)
	bs := []byte(s)

	for i := 0; i < ls; i++ {
		for i+2 < ls && string(bs[i]) == ":" && string(bs[i+1]) == "-" && (string(bs[i+2]) == ")" || string(bs[i+2]) == "(") {
			i += 3
			for i < ls && bs[i] == bs[i-1] {
				i++
			}
		}

		if i < ls {
			j++
			// fmt.Println(string(bs[:i]))
			// bs[j] = bs[i]
		}
	}

	fmt.Println(string(bs[:j]))
	fmt.Println(j)
	return string(bs[:j])
}
