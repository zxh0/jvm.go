package jls8.ch12;

/**
 * Example 12.5-2.
 * Dynamic Dispatch During Instance Creation
 */
public class Eg12_5_2 {
    
    private static class Super {
        Super() { printThree(); }
        void printThree() { System.out.println("three"); }
    }
    private static class Test extends Super {
        int three = (int)Math.PI;  // That is, 3
        void printThree() { System.out.println(three); }
    }
    
    public static void main(String[] args) {
        Test t = new Test();
        t.printThree();
    }
    
}
