package stdlib.basic.reflection;

import java.lang.reflect.Array;

import static helper.MyAssert.*;

// test java.lang.reflect.Array
public class ArrayTest implements Runnable {

    public static void main(String[] args) {
        new ArrayTest().run();
    }

    @Override
    public void run() {
        testNewArray();
        testGetNull();
        testSetNull();
        testGetNonArray();
        testSetNonArray();
        testGetBadIndex();
        testSetBadIndex();
        testGetObjectArray();
        testSetObjectArray();
        testGetPrimitiveArray();
        testSetPrimitiveArray();
    }

    private static void testNewArray() {
        Object[] arr = (Object[]) Array.newInstance(Object.class, 8);
        assertEquals(8, arr.length);
    }

    private static void testGetNull() {
        try {
            Object x = null;
            Array.get(x, 3);
            fail();
        } catch (NullPointerException e) {
            assertNull(e.getMessage());
        }
    }
    private static void testSetNull() {
        try {
            Object x = null;
            Array.set(x, 3, "a");
            fail();
        } catch (NullPointerException e) {
            assertNull(e.getMessage());
        }
    }

    private static void testGetNonArray() {
        try {
            String str = "abc";
            Array.get(str, 1);
            fail();
        } catch (IllegalArgumentException e) {
            assertEquals("Argument is not an array", e.getMessage());
        }
    }
    private static void testSetNonArray() {
        try {
            String str = "abc";
            Array.set(str, 1, "a");
            fail();
        } catch (IllegalArgumentException e) {
            assertEquals("Argument is not an array", e.getMessage());
        }
    }

    private static void testGetBadIndex() {
        try {
            int[] arr = {1, 2, 3};
            Array.get(arr, -1);
            fail();
        } catch (ArrayIndexOutOfBoundsException e) {
            assertNull(e.getMessage());
        }
    }
    private static void testSetBadIndex() {
        try {
            int[] arr = {1, 2, 3};
            Array.set(arr, -1, 4);
            fail();
        } catch (ArrayIndexOutOfBoundsException e) {
            assertNull(e.getMessage());
        }
    }

    private static void testGetObjectArray() {
        String[] arr = {"a", "b", "c"};
        assertEquals("c", Array.get(arr, 2));
    }
    private static void testSetObjectArray() {
        String[] arr = {"beyond"};
        Array.set(arr, 0, "5457");
        assertEquals("5457", Array.get(arr, 0));
    }

    private static void testGetPrimitiveArray() {
        assertEquals(true,     Array.get(new boolean[]{true}, 0));
        assertEquals((byte)2,  Array.get(new byte[]{2},       0));
        assertEquals('a',      Array.get(new char[]{'a'},     0));
        assertEquals((short)2, Array.get(new short[]{2},      0));
        assertEquals(2,        Array.get(new int[]{2},        0));
        assertEquals(2L,       Array.get(new long[]{2},       0));
        assertEquals(3.14f,    Array.get(new float[]{3.14f},  0));
        assertEquals(2.71,     Array.get(new double[]{2.71},  0));
    }
    private static void testSetPrimitiveArray() {
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
