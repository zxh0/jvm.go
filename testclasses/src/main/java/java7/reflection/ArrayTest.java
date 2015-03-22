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
    public void getNullArray() {
        try {
            Object x = null;
            Array.get(x, 3);
            fail();
        } catch (NullPointerException e) {
            assertNull(e.getMessage());
        }
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

    @Test
    public void getArrayBadIndex() {
        try {
            int[] arr = {1, 2, 3};
            Array.get(arr, -1);
            fail();
        } catch (ArrayIndexOutOfBoundsException e) {
            assertEquals(null, e.getMessage());
        }
    }

    @Test
    public void getObjectArray() {
        String[] arr = {"a", "b", "c"};
        assertEquals("c", Array.get(arr, 2));
    }
    
    @Test
    public void getPrimitiveArray() {
        int[] arr = {1, 2, 3};
        Object two = Array.get(arr, 1);
        assertEquals(2, two);
    }

    @Test
    public void setObjectArray() {
        String[] arr = {"beyond"};
        Array.set(arr, 0, "5457");
        assertEquals("5457", Array.get(arr, 0));
    }

    @Test
    public void setPrimitiveArray() {
        int[] arr = {5, 4, 5, 7};
        Array.set(arr, 0, 0);
        assertEquals(0, Array.get(arr, 0));

        //Array.set(arr, 0, "beyond");
    }

}
