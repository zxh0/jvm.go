package java6.string;

public class Mutf8Test {
    
    public static void main(String[] args) {
        System.out.println("1\u0001~\u007F"); // 1-byte code points
        System.out.println("2\u0000&\u0080~\u07FF"); // 2-bytes code points
        System.out.println("3\u0800~\uFFFF"); // 3-bytes code points
        //System.out.println(new String(new int[] {0x10000}, 0, 1));
    }
    
}
