package jls8;

/**
 * Example 12.4.1-2.
 * Only The Class That Declares static Field Is Initialized
 */
public class Eg12_4_1_2 {
    
    private static class Super {
        static int taxi = 1729;
    }
    private static class Sub extends Super {
        static { System.out.print("Sub "); }
    }
    
    public static void main(String[] args) {
        System.out.println(Sub.taxi);
    }
    
}
