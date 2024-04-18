package integertoroman

func intToRoman(num int) string {
	satuan := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	puluhan := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ratusan := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ribuan := []string{"", "M", "MM", "MMM"}

	return ribuan[num/1000] + ratusan[(num%1000)/100] + puluhan[(num%100)/10] + satuan[num%10]
}
