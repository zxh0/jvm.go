package jls8.ch12;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

/**
 * Example 12.4.1-3.
 * Interface Initialization Does Not Initialize Superinterfaces
 */
public class Eg12_4_1_3 {
    
    private static String output = "";
    
    private static interface I {
        int i = 1, ii = out("ii", 2);
    }
    private static interface J extends I {
        int j = out("j", 3), jj = out("jj", 4);
    }
    private static interface K extends J {
        int k = out("k", 5);
    }
    
    static int out(String s, int i) {
        println(s + "=" + i);
        return i;
    }
    
    private static void println(String s) {
        output += s;
        output += "\n";
    }
    
    @Test
    public void test() {
        println("" + J.i);
        println("" + K.j);
        assertEquals("1\nj=3\njj=4\n3\n", output);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(Eg12_4_1_3.class);
    }
    
}
