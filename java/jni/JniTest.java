
public class JniTest {
	static {
		System.loadLibrary("jni");
	}
	
	public static void main(String[] args){
		int r = add(1, 2);
		println(r);
	}
	
	public static native int add(int i, int j);
	
	public static native void println(int i);
}