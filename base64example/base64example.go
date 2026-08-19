package base64

import (
	"slices"
	"strings"
)

const base64 string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func EncodeBase64(s string) (string, error) {
	//split the string into bytes
	bytes := []byte(s)
	var output strings.Builder

	//chunk them
	for chunk := range slices.Chunk(bytes, 3) {
		chunkSize := len(chunk)

		var b [3]byte     // new array of 3 bytes
		copy(b[:], chunk) // slice operator produces a slice header that points the b array, copy only accepts slices

		position16to24 := int(b[0]) << 16
		position8to15 := int(b[1]) << 8
		position0to7 := int(b[2]) << 0

		// 3 letters are chunked into 3 sets of pairs
		// 1111 1111 2222 2222 3333 3333
		combined24int := position16to24 | position8to15 | position0to7

		// now split into sets of 4
		//111111 112222 222233 333333
		// shift the output by 18 then take the rest then chop of anything lower than 6 bits
		output.WriteByte(base64[(combined24int >> 18)]) //0x3F could be added here but redundant
		output.WriteByte(base64[(combined24int>>12)&0x3F])

		if chunkSize > 1 {
			//if chunk size is less than 1 there is no real data in the 3rd position
			//11111 1110000 000000 000000
			//              ^empty
			output.WriteByte(base64[(combined24int>>6)&0x3F])
		} else {
			output.WriteByte('=')
		}
		if chunkSize > 2 {
			//if chunk size is less than 2 there is no real data in the 4th position
			//111111 112222 222200 000000
			//                     ^empty
			output.WriteByte(base64[(combined24int)&0x3F])
		} else {
			output.WriteByte('=')
		}
	}

	return output.String(), nil
}
