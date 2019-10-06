package jls8.ch12;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

/**
 * Example 12.4.1-2.
 * Only The Class That Declares static Field Is Initialized
 */
public class Eg12_4_1_2 {
    
    private static class Super {
        static int taxi = 1729;
    }
    private static class Sub extends Super {
        static { 
            System.out.print("Sub ");
            if (true) {
                throw new RuntimeException("BAD");
            }
        }
    }
    
    @Test
    public void test() {
        assertEquals(1729, Sub.taxi);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(Eg12_4_1_2.class);
    }
    
}
