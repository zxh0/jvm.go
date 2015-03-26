package jls8.ch12;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

/**
 * Example 12.4.1-1. 
 * Superclasses Are Initialized Before Subclasses
 */
public class Eg12_4_1_1 {
    
    private static String output = "";
    
    private static class Super {
        static { Eg12_4_1_1.output += "Super "; }
    }
    private static class One {
        static { Eg12_4_1_1.output += "One "; }
    }
    private static class Two extends Super {
        static { Eg12_4_1_1.output += "Two "; }
    }
    
    @Test
    public void test() {
        One o = null;
        Two t = new Two();
        output += ((Object)o == (Object)t);
        assertEquals("Super Two false", output);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(Eg12_4_1_1.class);
    }
    
}
