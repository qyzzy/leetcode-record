package addStrings

func addStrings(num1 string, num2 string) string {
	carry, i, j, sum := 0, len(num1)-1, len(num2)-1, 0
	res := ""
	for i >= 0 && j >= 0 {
		sum = int(num1[i]-'0') + int(num2[j]-'0') + carry
		if sum >= 10 {
			carry = 1
			res = string(sum-10+'0') + res
		} else {
			carry = 0
			res = string(sum+'0') + res
		}
		i--
		j--
	}
	for i >= 0 {
		sum = int(num1[i]-'0') + carry
		if sum >= 10 {
			carry = 1
			res = string(sum-10+'0') + res
		} else {
			carry = 0
			res = string(sum+'0') + res
		}
		i--
	}
	for j >= 0 {
		sum = int(num2[j]-'0') + carry
		if sum >= 10 {
			carry = 1
			res = string(sum-10+'0') + res
		} else {
			carry = 0
			res = string(sum+'0') + res
		}
		j--
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}
