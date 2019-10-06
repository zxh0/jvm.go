package jls8.ch12;

import jls8.StringOut;
import libs.junit.UnitTestRunner;
import static org.junit.Assert.*;

/**
 * Example 12.5-2.
 * Dynamic Dispatch During Instance Creation
 */
public class Eg12_5_2 {
    
    private static final StringOut out = new StringOut();
    
    private static class Super {
        Super() { printThree(); }
        void printThree() { out.println("three"); }
    }
    private static class Test extends Super {
        int three = (int)Math.PI;  // That is, 3
        void printThree() { out.println(three); }
    }
    
    @org.junit.Test
    public void test() {
        Test t = new Test();
        t.printThree();
        assertEquals("0\n3\n", out.toString());
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(Eg12_5_2.class);
    }
    
}
