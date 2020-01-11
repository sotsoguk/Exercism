package variablelengthquantity

func prependByte(x []byte, y byte) []byte{
	x = append(x,0)
	copy(x[1:],x)
	x[0] = y
	return x

	}	
func EncodeVarint(input []uint32) []byte {
	result := make([]byte,0)
	for _,i :=range input {
		if i== 0 {
			result = append(result,0x0)
		}
		tmp := make([]byte,0)
		lastByte := true
		for i > 0 {
			b := i & 0x7f
			if !lastByte {
				b |= 0x80
			}
			// tmp = prependByte(tmp,byte(b))
			tmp = append([]byte{byte(b)},tmp...)
			// result = append([]byte{byte(b)},result...)
			lastByte = false
			i >>=7
		}
		result = append(result,tmp...)
	}
	return result
}