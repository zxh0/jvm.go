package jvm.instructions;

import static helper.MyAssert.*;

// newarray, anewarray, arraylength, multianewarray
public class RefArrayTest implements Runnable {

    public static void main(String[] args) {
        new RefArrayTest().run();
    }

    @Override
    public void run() {
        testSimpleArray();
        testMultiDimensionalArray();
    }

    private static void testSimpleArray() {
        // newarray
        assertEquals("[Z", new boolean[8].getClass().getName());
        assertEquals("[B", new byte[8].getClass().getName());
        assertEquals("[C", new char[8].getClass().getName());
        assertEquals("[S", new short[8].getClass().getName());
        assertEquals("[I", new int[8].getClass().getName());
        assertEquals("[J", new long[8].getClass().getName());
        assertEquals("[F", new float[8].getClass().getName());
        assertEquals("[D", new double[8].getClass().getName());
        // anewarray
        assertEquals("[Ljava.lang.Object;", new Object[8].getClass().getName());
        // arraylength
        assertEquals(16, new Object[16].length);
    }

    private static void testMultiDimensionalArray() {
        int[][][] a = new int[1][2][3]; // multianewarray
        assertEquals("[[[I", a.getClass().getName());

        int[][][] b = {
            {
                {1},
                {1, 2},
                {1, 2, 3}
            }
        };
        assertEquals(1, b.length);
        assertEquals(3, b[0].length);
        assertEquals(1, b[0][0].length);
        assertEquals(2, b[0][1].length);
        assertEquals(3, b[0][2].length);
        assertEquals(2, b[0][2][1]);
    }

}
