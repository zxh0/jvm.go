package java6.string;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

// http://www.oracle.com/technetwork/articles/javase/supplementary-142654.html
public class Mutf8Test {
    
    public static void main(String[] args) {
        UnitTestRunner.run(Mutf8Test.class);
    }
    
    @Test
    public void test() {
        assertEquals("A", 1, "A".length()); // U+0041
        assertEquals("NULL", 1, "\u0000".length());
        assertEquals("ß", 1, "ß".length()); // U+00DF
        assertEquals("東", 1, "東".length()); // U+6771
        assertEquals("U+10400", 2, "\uD801\uDC00".length()); // U+10400
    }
    
}
