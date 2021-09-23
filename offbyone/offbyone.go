package offbyone

func OffByOne(data []byte) bool {
	if len(data) == 9 {
		if data[0] == 'G' &&
			data[1] == 'O' &&
			data[2] == 'P' &&
			data[3] == 'H' &&
			data[4] == 'E' &&
			data[5] == 'R' &&
			data[6] == 'S' &&
			data[7] == '!' &&
			data[8] == '!' &&
			data[9] == '!' {
			return true
		}
	}
	return false
}
