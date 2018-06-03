
import java.util.stream.IntStream;

public class RomanNumber {

	private static final java.util.Map<Character, Integer> rm;

	static {
		rm = new java.util.HashMap<>();
		rm.put('I', 1);
		rm.put('V', 5);
		rm.put('X', 10);
		rm.put('L', 50);
		rm.put('C', 100);
		rm.put('D', 500);
		rm.put('M', 1000);
	}

	public static void main(String[] args) {
		IntStream.range(0, 4999).parallel().forEach(i -> {
			String r = intToRoman(i);
			if (romanToInt(r) != i) throw new RuntimeException(r + " != " + i);
		});
	}

	/**
	 * 逆序遍历字符串
	 * 如果当前字符值大于等于前一个 则相加到sum
	 * 小于前一个 则减去
	 */
	public static int romanToInt(String roman) {
		int n = roman.length();
		int lastVal = 0;
		int sum = 0;
		for (int i = n - 1; i >= 0; i--) {
			int val = rm.get(roman.charAt(i));
			if (val >= lastVal) {
				sum += val;
			} else {
				sum -= val;
			}
			lastVal = val;
		}
		return sum;
	}

	/**
	 *
	 */
	public static String intToRoman(int num) {
		String M[] = {"", "M", "MM", "MMM", "MMMM"};
		String C[] = {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"};
		String X[] = {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"};
		String I[] = {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};
		return M[num / 1000] + C[(num % 1000) / 100] + X[(num % 100) / 10] + I[num % 10];
	}
}
