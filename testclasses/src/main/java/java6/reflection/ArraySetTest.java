package java6.reflection;

import java.lang.reflect.Array;

import libs.junit.UnitTestRunner;
import org.junit.Test;

import static org.junit.Assert.*;

public class ArraySetTest {

    public static void main(String[] args) {
        UnitTestRunner.run(ArraySetTest.class);
    }

    @Test
    public void setNullArray() {
        try {
            Object x = null;
            Array.set(x, 3, "a");
            fail();
        } catch (NullPointerException e) {
            assertNull(e.getMessage());
        }
    }

    @Test
    public void setNonArray() {
        try {
            String str = "abc";
            Array.set(str, 1, "a");
            fail();
        } catch (IllegalArgumentException e) {
            assertEquals("Argument is not an array", e.getMessage());
        }
    }

    @Test
    public void setArrayTypeMismatch() {
        try {
            int[] arr = {1, 2, 3};
            Array.set(arr, 0, "beyond");
            fail();
        } catch (IllegalArgumentException e) {
            assertEquals("argument type mismatch", e.getMessage());
        }
    }

    @Test
    public void setArrayBadIndex() {
        try {
            int[] arr = {1, 2, 3};
            Array.set(arr, -1, 4);
            fail();
        } catch (ArrayIndexOutOfBoundsException e) {
            assertEquals(null, e.getMessage());
        }
    }

    @Test
    public void setObjectArray() {
        String[] arr = {"beyond"};
        Array.set(arr, 0, "5457");
        assertEquals("5457", Array.get(arr, 0));
    }

    @Test
    public void setPrimitiveArray() {
        Array.set(new boolean[]{true}, 0, false);
        Array.set(new byte[]{2}, 0, (byte) 3);
        Array.set(new char[]{'a'}, 0, 'b');
        Array.set(new short[]{2}, 0, (short) 3);
        Array.set(new int[]{2}, 0, 3);
        Array.set(new long[]{2}, 0, 3L);
        Array.set(new float[]{3.14f}, 0, 2.71f);
        Array.set(new double[]{2.71}, 0, 3.14);

        int[] arr = {5, 4, 5, 7};
        Array.set(arr, 0, 0);
        assertEquals(0, Array.get(arr, 0));

    }

}
