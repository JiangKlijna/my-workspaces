public class Hex{
	private String hex;
	private int ten;
	
	public Hex(){
		this.hex="0";
		this.ten=0;
	}
	public Hex(String hex){
		this.hex=hex;
		this.ten=hex_ten(hex);
	}
	public Hex(int ten){
		this.ten=ten;
		this.hex=ten_hex(ten);
	}
	
	public String toString(){
		return hex;
	}
	
	public int hex_ten(String aa) {
		char[] aaa = aa.toCharArray();
		int re = 0;
		for (int i = aaa.length - 1; i >= 0; i--) {
			int b = aaa.length - 1 - i;
			int t;
			if ((int) aaa[i] >= 65) {
				t = (int) (aaa[i]) - 55;
			} else {
				t = Integer.parseInt(String.valueOf(aaa[i]));
			}
			re += t * Math.pow(16, b);
		}
		return re;
	}
	
	public String ten_hex(int a) {
		String re = "";
		while (a > 0) {
			int t = a % 16;
			if (t >= 10) {
				re = (char) (t + 55) + re;
			} else {
				re = String.valueOf(t) + re;
			}
			a = a / 16;
		}
		return re;
	}
	
	public Hex add(Hex h){
		return new Hex(this.ten+h.ten);
	}
	
	public static void main(String[] args){
		Hex h1=new Hex(123);
		Hex h2=new Hex("123");
		System.out.println(h1.add(h2));
		
	}
	
}