package sto

func hammingWeight(num uint32) int {
	var res uint32
	for num > 0 {
		res = res + (num & 1)
		num = num >> 1
	}
	return int(res)
}

func hammingWeight2(num uint32) int {
	var res int
	for num > 0 {
		num &= num-1
		res += 1
	}
	return res
}