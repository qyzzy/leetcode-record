package textJustification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	tmp := ""
	for i := 0; i < len(words); i++ {
		if tmp == "" {
			tmp += words[i]
			continue
		}
		if len(tmp+" "+words[i]) > maxWidth {
			res = append(res, build(tmp, maxWidth))
			tmp = words[i]
		} else {
			tmp += " " + words[i]
		}
	}
	wds := strings.Split(tmp, " ")
	tmp = wds[0]
	for i := 1; i < len(wds); i++ {
		tmp += " "
		tmp += wds[i]
	}
	length := len(tmp)
	for j := 0; j < maxWidth-length; j++ {
		tmp += " "
	}
	res = append(res, tmp)
	return res
}

func build(tmp string, maxW int) string {
	words := strings.Split(tmp, " ")
	if len(words) == 1 {
		res := words[0]
		spaces := maxW - len(words[0])
		for i := 0; i < spaces; i++ {
			res += " "
		}
		return res
	}
	spacesAvg := (maxW - (len(tmp) - len(words) + 1)) / (len(words) - 1)
	spacesMore := maxW - (len(tmp) - len(words) + 1) - spacesAvg*(len(words)-1)
	res := strings.Builder{}
	res.WriteString(words[0])
	for i := 1; i < len(words); i++ {
		for j := 0; j < spacesAvg; j++ {
			res.WriteByte(' ')
		}
		if spacesMore > 0 {
			res.WriteByte(' ')
			spacesMore--
		}
		res.WriteString(words[i])
	}
	return res.String()
}
