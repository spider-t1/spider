package douyin

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Enc 主要的编码函数，对应 JavaScript 中的 enc 函数
func Enc(url, data, userAgent string) string {
	// 生成乱码字符串
	params := url[strings.Index(url, "?")+1:] + "dhzx"
	data += "dhzx"
	garbledString := getGarbledString(params, data, userAgent)
	shortStr := "Dkdpgh2ZmsQB80/MfvV36XI1R45-WUAlEixNLwoqYTOPuzKFjJnry79HbGcaStCe"
	var aBogus strings.Builder

	// 依次生成七组字符串
	j := 0
	for i := 0; i <= len(garbledString); i += 3 {
		if (i + 3) <= len(garbledString) {
			charCodeAtNum0 := int(garbledString[i])
			charCodeAtNum1 := int(garbledString[i+1])
			charCodeAtNum2 := int(garbledString[i+2])
			baseNum := charCodeAtNum2 | charCodeAtNum1<<8 | charCodeAtNum0<<16

			str1 := string(shortStr[(baseNum&16515072)>>18])
			str2 := string(shortStr[(baseNum&258048)>>12])
			str3 := string(shortStr[(baseNum&4032)>>6])
			str4 := string(shortStr[baseNum&63])
			aBogus.WriteString(str1 + str2 + str3 + str4)
		}
		if i+3 > len(garbledString) {
			u := len(garbledString) - j
			if u == 2 {
				charCodeAtNum0 := int(garbledString[j])
				charCodeAtNum1 := int(garbledString[j+1])
				baseNum := charCodeAtNum1<<8 | charCodeAtNum0<<16
				str1 := string(shortStr[(baseNum&16515072)>>18])
				str2 := string(shortStr[(baseNum&258048)>>12])
				str3 := string(shortStr[(baseNum&4032)>>6])
				aBogus.WriteString(str1 + str2 + str3 + "=")
			}
			if u == 1 {
				charCodeAtNum0 := int(garbledString[j])
				baseNum := 0 | charCodeAtNum0<<16
				str1 := string(shortStr[(baseNum&16515072)>>18])
				str2 := string(shortStr[(baseNum&258048)>>12])
				aBogus.WriteString(str1 + str2 + "==")
			}
		}
		j += 3
	}
	return aBogus.String()
}

// getGarbledString 生成混淆字符串
func getGarbledString(params, data, userAgent string) string {
	timestamp1 := time.Now().UnixMilli()
	timestamp2 := timestamp1 - int64(math.Floor(rand.Float64()*10))
	arr29 := getArr29(timestamp1, timestamp2, params, data, userAgent)
	a := topHeaderRandomGarbledCharacters()
	b := abGarbledCharacters(string(arr29))
	return a + b
}

// getArr29 生成29位数组
func getArr29(dateTime1, dateTime2 int64, params, data, userAgent string) []byte {
	parArr := getArr(string(getArr(params)))
	dataArr := getArr(string(getArr(data)))

	uaSalt := 0 // 固定为0，对应JavaScript中的逻辑

	browserArr := getArr(encryptionUa(uaGarbledCharacters(userAgent, uaSalt)))
	arr := make([]int, 55)
	arr2 := make([]byte, 50)
	num := getArr2(userAgent)
	dateTime3 := int((float64(time.Now().UnixMilli())-1721836800000)/1000/60/60/24/14) >> 0
	arr0 := randomGarbledCharactersArrayList()

	arr[0] = 41
	arr[1] = dateTime3
	arr[2] = 6
	arr[3] = (int(dateTime1-dateTime2) + 3) & 255
	arr[4] = int(dateTime1) >> 0 & 255
	arr[5] = int(dateTime1) >> 8 & 255
	arr[6] = int(dateTime1) >> 16 & 255
	arr[7] = int(dateTime1) >> 24 & 255
	arr[8] = int(dateTime1/256/256/256/256) & 255
	arr[9] = int(dateTime1/256/256/256/256/256) & 255
	arr[10] = 1 % 256 & 255
	arr[11] = 1 / 256 & 255
	arr[12] = 1 & 255
	arr[13] = 1 >> 8 & 255
	arr[14] = 0
	arr[15] = 0 % 101 % 256 & 255
	arr[16] = 0 % 201 % 256 & 255
	arr[17] = 0 % 101 % 256 & 255
	arr[18] = uaSalt & 255
	arr[19] = uaSalt >> 8 & 255
	arr[20] = uaSalt >> 16 & 255
	arr[21] = uaSalt >> 24 & 255
	arr[22] = int(parArr[9])
	arr[23] = int(parArr[18])
	arr[24] = 3
	arr[25] = int(parArr[3])
	arr[26] = int(dataArr[10])
	arr[27] = int(dataArr[19])
	arr[28] = 4
	arr[29] = int(dataArr[4])
	arr[30] = int(browserArr[11])
	arr[31] = int(browserArr[21])
	arr[32] = 5
	arr[33] = int(browserArr[5])
	arr[34] = int(dateTime2) >> 0 & 255
	arr[35] = int(dateTime2) >> 8 & 255
	arr[36] = int(dateTime2) >> 16 & 255
	arr[37] = int(dateTime2) >> 24 & 255
	arr[38] = int(dateTime2/256/256/256/256) & 255
	arr[39] = int(dateTime2/256/256/256/256/256) & 255
	arr[40] = 3
	arr32 := 97
	arr[41] = arr32 >> 0 & 255
	arr[42] = arr32 >> 8 & 255
	arr[43] = arr32 >> 16 & 255
	arr[44] = arr32 >> 24 & 255
	arr36 := 6399
	arr[45] = arr36 & 255
	arr[46] = arr36 >> 8 & 255
	arr[47] = arr36 >> 16 & 255
	arr[48] = arr36 >> 24 & 255
	lastNumOne := getLast3Num(dateTime1)
	arr[49] = len(num)
	arr[50] = len(num) & 255
	arr[51] = len(num) >> 8 & 255
	arr[52] = len(lastNumOne)
	arr[53] = len(lastNumOne) & 255
	arr[54] = len(lastNumOne) >> 8 & 255

	lastNum := getLastNum2(arr0, arr)

	// 重新排列数组
	arr2[0] = byte(arr[9])
	arr2[1] = byte(arr[18])
	arr2[2] = byte(arr[30])
	arr2[3] = byte(arr[35])
	arr2[4] = byte(arr[47])
	arr2[5] = byte(arr[4])
	arr2[6] = byte(arr[44])
	arr2[7] = byte(arr[19])
	arr2[8] = byte(arr[10])
	arr2[9] = byte(arr[23])
	arr2[10] = byte(arr[12])
	arr2[11] = byte(arr[40])
	arr2[12] = byte(arr[25])
	arr2[13] = byte(arr[42])
	arr2[14] = byte(arr[3])
	arr2[15] = byte(arr[22])
	arr2[16] = byte(arr[38])
	arr2[17] = byte(arr[21])
	arr2[18] = byte(arr[5])
	arr2[19] = byte(arr[45])
	arr2[20] = byte(arr[1])
	arr2[21] = byte(arr[29])
	arr2[22] = byte(arr[6])
	arr2[23] = byte(arr[43])
	arr2[24] = byte(arr[33])
	arr2[25] = byte(arr[14])
	arr2[26] = byte(arr[36])
	arr2[27] = byte(arr[37])
	arr2[28] = byte(arr[2])
	arr2[29] = byte(arr[46])
	arr2[30] = byte(arr[15])
	arr2[31] = byte(arr[48])
	arr2[32] = byte(arr[31])
	arr2[33] = byte(arr[26])
	arr2[34] = byte(arr[16])
	arr2[35] = byte(arr[13])
	arr2[36] = byte(arr[8])
	arr2[37] = byte(arr[41])
	arr2[38] = byte(arr[27])
	arr2[39] = byte(arr[17])
	arr2[40] = byte(arr[39])
	arr2[41] = byte(arr[20])
	arr2[42] = byte(arr[11])
	arr2[43] = byte(arr[0])
	arr2[44] = byte(arr[34])
	arr2[45] = byte(arr[7])
	arr2[46] = byte(arr[50])
	arr2[47] = byte(arr[51])
	arr2[48] = byte(arr[53])
	arr2[49] = byte(arr[54])

	// 合并数组
	newArr := append(arr2, num...)
	newArr2 := append(newArr, lastNumOne...)
	newArr2 = append(newArr2, byte(lastNum))

	return getNumList(arr0, newArr2)
}

// getArr 实现JavaScript中的自定义哈希算法
func getArr(input string) []byte {
	// 初始化寄存器
	reg := []uint32{1937774191, 1226093241, 388252375, 3666478592, 2842636476, 372324522, 3817729613, 2969243214}

	// 将输入字符串转换为字节数组
	data := []byte(input)
	size := len(data)

	// 添加填充
	data = append(data, 0x80)
	for (len(data) % 64) != 56 {
		data = append(data, 0x00)
	}

	// 添加长度信息（64位，大端序）
	bitLen := uint64(size * 8)
	for i := 0; i < 4; i++ {
		data = append(data, byte((bitLen>>(8*(7-i)))&0xff))
	}
	for i := 4; i < 8; i++ {
		data = append(data, byte((bitLen>>(8*(7-i)))&0xff))
	}

	// 处理每个64字节块
	for i := 0; i < len(data); i += 64 {
		reg = compressBlock(reg, data[i:i+64])
	}

	// 将结果转换为32字节数组
	result := make([]byte, 32)
	for i := 0; i < 8; i++ {
		val := reg[i]
		result[4*i+3] = byte(val & 0xff)
		val >>= 8
		result[4*i+2] = byte(val & 0xff)
		val >>= 8
		result[4*i+1] = byte(val & 0xff)
		val >>= 8
		result[4*i] = byte(val & 0xff)
	}

	return result
}

// compressBlock 压缩64字节数据块
func compressBlock(reg []uint32, block []byte) []uint32 {
	// 扩展消息
	w := make([]uint32, 132)

	// 前16个字（大端序）
	for i := 0; i < 16; i++ {
		w[i] = uint32(block[4*i])<<24 | uint32(block[4*i+1])<<16 | uint32(block[4*i+2])<<8 | uint32(block[4*i+3])
	}

	// 扩展到68个字
	for i := 16; i < 68; i++ {
		temp := w[i-16] ^ w[i-9] ^ rotateLeft(w[i-3], 15)
		temp = temp ^ rotateLeft(temp, 15) ^ rotateLeft(temp, 23)
		w[i] = temp ^ rotateLeft(w[i-13], 7) ^ w[i-6]
	}

	// 生成W'
	for i := 0; i < 64; i++ {
		w[i+68] = w[i] ^ w[i+4]
	}

	// 复制寄存器
	a, b, c, d, e, f, g, h := reg[0], reg[1], reg[2], reg[3], reg[4], reg[5], reg[6], reg[7]

	// 64轮压缩
	for j := 0; j < 64; j++ {
		ss1 := rotateLeft(rotateLeft(a, 12)+e+rotateLeft(getTj(j), uint32(j)), 7)
		ss2 := ss1 ^ rotateLeft(a, 12)
		tt1 := getFF(j, a, b, c) + d + ss2 + w[j+68]
		tt2 := getGG(j, e, f, g) + h + ss1 + w[j]

		d = c
		c = rotateLeft(b, 9)
		b = a
		a = tt1
		h = g
		g = rotateLeft(f, 19)
		f = e
		e = getP0(tt2)
	}

	// 更新寄存器
	newReg := make([]uint32, 8)
	newReg[0] = reg[0] ^ a
	newReg[1] = reg[1] ^ b
	newReg[2] = reg[2] ^ c
	newReg[3] = reg[3] ^ d
	newReg[4] = reg[4] ^ e
	newReg[5] = reg[5] ^ f
	newReg[6] = reg[6] ^ g
	newReg[7] = reg[7] ^ h

	return newReg
}

// rotateLeft 左旋转
func rotateLeft(x uint32, n uint32) uint32 {
	n = n % 32
	return (x << n) | (x >> (32 - n))
}

// getTj 获取常数Tj
func getTj(j int) uint32 {
	if j >= 0 && j < 16 {
		return 2043430169
	} else if j >= 16 && j < 64 {
		return 2055708042
	}
	return 0
}

// getFF 布尔函数FF
func getFF(j int, x, y, z uint32) uint32 {
	if j >= 0 && j < 16 {
		return x ^ y ^ z
	} else if j >= 16 && j < 64 {
		return (x & y) | (x & z) | (y & z)
	}
	return 0
}

// getGG 布尔函数GG
func getGG(j int, x, y, z uint32) uint32 {
	if j >= 0 && j < 16 {
		return x ^ y ^ z
	} else if j >= 16 && j < 64 {
		return (x & y) | (^x & z)
	}
	return 0
}

// getP0 置换函数P0
func getP0(x uint32) uint32 {
	return x ^ rotateLeft(x, 9) ^ rotateLeft(x, 17)
}

// GetArrTest 导出的测试函数，用于验证 getArr 函数
func GetArrTest(input string) []byte {
	return getArr(input)
}

// abGarbledCharacters 字符混淆算法
func abGarbledCharacters(userAgent string) string {
	arr256 := abArr256()
	n4 := 0
	var ans strings.Builder

	for i := 0; i < len(userAgent); i++ {
		n2 := (i + 1) % 256
		n3 := n4 + arr256[n2]
		n4 = n3 % 256
		oldArrN2 := arr256[n2]
		arr256[n2] = arr256[n4]
		arr256[n4] = oldArrN2
		n5 := int(userAgent[i])
		n6 := arr256[n2] + oldArrN2
		n7 := n6 % 256
		n8 := n5 ^ arr256[n7]
		ans.WriteByte(byte(n8))
	}
	return ans.String()
}

// abArr256 生成256位数组
func abArr256() []int {
	nums := make([]int, 256)
	for i := 0; i < 256; i++ {
		nums[i] = 255 - i
	}

	previousValue := 0
	lm := 211 // String.fromCharCode(211)

	for i := 0; i < len(nums); i++ {
		num1 := previousValue * nums[i]
		previousValue = (num1 + previousValue + lm) % 256
		tmp := nums[i]
		nums[i] = nums[previousValue]
		nums[previousValue] = tmp
	}
	return nums
}

// topHeaderRandomGarbledCharacters 生成随机头部字符
func topHeaderRandomGarbledCharacters() string {
	arr := make([]byte, 4)
	random := rand.Float64() * 65535
	num1 := int(random) & 255
	num2 := rand.Intn(40)

	arr[0] = byte((num1 & 170) | (3 & 85))
	arr[1] = byte((num1 & 85) | (3 & 170))
	arr[2] = byte((num2 & 170) | (82 & 85))
	arr[3] = byte((num2 & 85) | (82 & 170))

	return string(arr)
}

// randomGarbledCharactersArrayList 生成随机混淆字符数组
func randomGarbledCharactersArrayList() []byte {
	arr1 := randomGarbledCharactersArray1()
	arr2 := randomGarbledCharactersArray2()
	return append(arr1, arr2...)
}

func randomGarbledCharactersArray1() []byte {
	arr := make([]byte, 4)
	random := rand.Float64() * 65535
	num1 := int(random) & 255
	num2 := (int(random) >> 8) & 255

	arr[0] = byte((num1 & 170) | (1 & 85))
	arr[1] = byte((num1 & 85) | (1 & 170))
	arr[2] = byte((num2 & 170) | (0 & 85))
	arr[3] = byte((num2 & 85) | (0 & 170))

	return arr
}

func randomGarbledCharactersArray2() []byte {
	arr := make([]byte, 4)
	num1 := rand.Intn(240)
	num2 := (rand.Intn(255) & 77) | 2 | 16 | 32 | 128

	arr[0] = byte((num1 & 170) | (1 & 85))
	arr[1] = byte((num1 & 85) | (1 & 170))
	arr[2] = byte((num2 & 170) | (0 & 85))
	arr[3] = byte((num2 & 85) | (0 & 170))

	return arr
}

// encryptionUa Base64类似的编码算法
func encryptionUa(ss string) string {
	str := "ckdp1h4ZKsUB80/Mfvw36XIgR25+WQAlEi7NLboqYTOPuzmFjJnryx9HVGDaStCe"
	var uaEncryption strings.Builder
	j := 0

	for i := 0; i < len(ss); i += 3 {
		if i+3 <= len(ss) {
			number := (((int(ss[i]) & 255) << 16) | ((int(ss[i+1]) & 255) << 8)) | ((int(ss[i+2]) & 255) << 0)
			uaEncryption.WriteByte(str[(number&16515072)>>18])
			uaEncryption.WriteByte(str[(number&258048)>>12])
			uaEncryption.WriteByte(str[(number&4032)>>6])
			uaEncryption.WriteByte(str[(number&63)>>0])
		}
		if i+3 > len(ss) {
			u := len(ss) - j
			if u == 2 {
				charCodeAtNum0 := int(ss[j])
				charCodeAtNum1 := int(ss[j+1])
				baseNum := charCodeAtNum1<<8 | charCodeAtNum0<<16
				str1 := str[(baseNum&16515072)>>18]
				str2 := str[(baseNum&258048)>>12]
				str3 := str[(baseNum&4032)>>6]
				uaEncryption.WriteByte(str1)
				uaEncryption.WriteByte(str2)
				uaEncryption.WriteByte(str3)
				uaEncryption.WriteByte('=')
			}
			if u == 1 {
				charCodeAtNum0 := int(ss[j])
				baseNum := 0 | charCodeAtNum0<<16
				str1 := str[(baseNum&16515072)>>18]
				str2 := str[(baseNum&258048)>>12]
				uaEncryption.WriteByte(str1)
				uaEncryption.WriteByte(str2)
				uaEncryption.WriteString("==")
			}
		}
		j += 3
	}
	return uaEncryption.String()
}

// uaGarbledCharacters UA字符混淆
func uaGarbledCharacters(userAgent string, uaSalt int) string {
	arr256 := uaArr256(userAgent, uaSalt)
	n4 := 0
	var ans strings.Builder

	for i := 0; i < len(userAgent); i++ {
		n2 := (i + 1) % 256
		n3 := n4 + arr256[n2]
		n4 = n3 % 256
		oldArrN2 := arr256[n2]
		arr256[n2] = arr256[n4]
		arr256[n4] = oldArrN2
		n5 := int(userAgent[i])
		n6 := arr256[n2] + oldArrN2
		n7 := n6 % 256
		n8 := n5 ^ arr256[n7]
		ans.WriteByte(byte(n8))
	}
	return ans.String()
}

// uaArr256 生成UA专用的256位数组
func uaArr256(userAgent string, uaSalt int) []int {
	nums := make([]int, 256)
	for i := 0; i < 256; i++ {
		nums[i] = 255 - i
	}

	previousValue := 0
	lm := []byte{0, 1, byte(uaSalt)}

	for i := 0; i < len(nums); i++ {
		num1 := previousValue * nums[i]
		previousValue = (num1 + previousValue + int(lm[i%3])) % 256
		tmp := nums[i]
		nums[i] = nums[previousValue]
		nums[previousValue] = tmp
	}
	return nums
}

// getLast3Num 获取最后3位数字
func getLast3Num(dateTime1 int64) []byte {
	intStr := int(dateTime1) + 3
	timestamp1 := fmt.Sprintf("%d,", intStr&255)
	num := make([]byte, len(timestamp1))
	for i := 0; i < len(timestamp1); i++ {
		num[i] = timestamp1[i]
	}
	return num
}

// getArr2 获取数组2
func getArr2(userAgent string) []byte {
	var data string
	if strings.Contains(userAgent, "Win") {
		data = "523|695|1536|816|1536|816|1536|864|Win32"
	} else {
		data = "523|695|1536|816|1536|816|1536|864|Win32"
	}

	sum := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		sum[i] = data[i]
	}
	return sum
}

// getLastNum2 获取最后数字2
func getLastNum2(arr1 []byte, arr []int) int {
	xorResult := int(arr1[0]) ^ int(arr1[1]) ^ int(arr1[2]) ^ int(arr1[3]) ^ int(arr1[4]) ^ int(arr1[5]) ^ int(arr1[6]) ^ int(arr1[7])
	for i := 0; i < len(arr); i++ {
		// 对齐 JS 版本：排除索引 24、28、32、49、52
		if i == 24 || i == 28 || i == 32 || i == 49 || i == 52 {
			continue
		}
		xorResult ^= arr[i]
	}
	return xorResult
}

// getNumList 获取数字列表
func getNumList(arr0 []byte, arrAr []byte) []byte {
	var numList []byte

	for i := 0; i < len(arrAr); i += 3 {
		if i+2 >= len(arrAr) {
			if i+1 >= len(arrAr) {
				num1 := arrAr[i]
				numList = append(numList, num1)
			} else {
				num1 := arrAr[i]
				num2 := arrAr[i+1]
				numList = append(numList, num1, num2)
			}
		} else {
			random := rand.Intn(1000) & 255
			num1 := (random & 145) | (int(arrAr[i]) & 110)
			num2 := (random & 66) | (int(arrAr[i+1]) & 189)
			num3 := (random & 44) | (int(arrAr[i+2]) & 211)
			num4 := ((int(arrAr[i]) & 145) | (int(arrAr[i+1]) & 66)) | (int(arrAr[i+2]) & 44)
			numList = append(numList, byte(num1), byte(num2), byte(num3), byte(num4))
		}
	}

	return append(arr0, numList...)
}
