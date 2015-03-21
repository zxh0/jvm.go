package java7.reflection;

import java.lang.reflect.Array;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ArrayTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ArrayTest.class);
    }
    
    @Test(expected = NullPointerException.class)
    public void getNullArray() {
        Object x = null;
        Array.get(x, 3);
    }
    
    @Test
    public void getNonArray() {
        try {
            String str = "abc";
            Array.get(str, 1);
            fail();
        } catch (IllegalArgumentException e) {
            assertEquals("Argument is not an array", e.getMessage());
        }
    }
    
    //@Test
    public void get() {
        int[] arr = {1, 2, 3};
        Object two = Array.get(arr, 1);
        assertEquals(2, two);
    }
    
}
