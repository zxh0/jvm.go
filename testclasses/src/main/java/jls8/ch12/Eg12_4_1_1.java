package jls8.ch12;

import jls8.StringOut;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

/**
 * Example 12.4.1-1. 
 * Superclasses Are Initialized Before Subclasses
 */
public class Eg12_4_1_1 {
    
    private static final StringOut out = new StringOut();
    
    private static class Super {
        static { out.print("Super "); }
    }
    private static class One {
        static { out.print("One "); }
    }
    private static class Two extends Super {
        static { out.print("Two "); }
    }
    
    @Test
    public void test() {
        One o = null;
        Two t = new Two();
        out.println((Object)o == (Object)t);
        assertEquals("Super Two false\n", out.toString());
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(Eg12_4_1_1.class);
    }
    
}
