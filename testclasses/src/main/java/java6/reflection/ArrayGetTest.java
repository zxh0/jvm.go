package java6.reflection;

import java.lang.reflect.Array;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ArrayGetTest {

    public static void main(String[] args) {
        UnitTestRunner.run(ArrayGetTest.class);
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
        assertEquals(true,     Array.get(new boolean[]{true}, 0));
        assertEquals((byte)2,  Array.get(new byte[]{2},       0));
        assertEquals('a',      Array.get(new char[]{'a'},     0));
        assertEquals((short)2, Array.get(new short[]{2},      0));
        assertEquals(2,        Array.get(new int[]{2},        0));
        assertEquals(2L,       Array.get(new long[]{2},       0));
        assertEquals(3.14f,    Array.get(new float[]{3.14f},  0));
        assertEquals(2.71,     Array.get(new double[]{2.71},  0));
    }

}
