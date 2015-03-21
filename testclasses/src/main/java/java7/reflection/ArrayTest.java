package java7.reflection;

import java.lang.reflect.Array;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ArrayTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ArrayTest.class);
    }
    
    @Test
    public void get() {
        int[] arr = {1, 2, 3};
        Object two = Array.get(arr, 1);
        assertEquals(2, two);
    }
    
}
