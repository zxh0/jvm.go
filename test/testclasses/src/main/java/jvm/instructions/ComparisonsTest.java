package jvm.instructions;

import static helper.MyAssert.*;

public class ComparisonsTest implements Runnable {

    public static void main(String[] args) {
        new ComparisonsTest().run();
    }

    @Override
    public void run() {
        testCmp();
        testIf();
        testIfICmp();
        testExtended();
    }

    private static void testCmp() {
        long j = 101;
        float f = 101.0f;
        double g = 101.0;
        assertTrue(j < 102 && j > 100 && j == 101); // lcmp
        assertTrue(f > 100.0f); // fcmpl
        assertTrue(f < 102.0f); // fcmpg
        assertTrue(g > 100.0);  // gcmpl
        assertTrue(g < 102.0);  // gcmpg
    }

    private static void testIf() {
        int i = 100;
        assertFalse(i == 0); // ifne
        assertTrue(i != 0);  // ifeq
        assertFalse(i < 0);  // ifge
        assertFalse(i <= 0); // ifgt
        assertTrue(i > 0);   // ifle
        assertTrue(i >= 0);  // iflt
    }

    private static void testIfICmp() {
        int i = 100;
        assertTrue(i == 100); // if_icmpne
        assertTrue(i != 99);  // if_icmpeq
        assertTrue(i < 101);  // if_icmpge
        assertTrue(i <= 101); // if_icmpgt
        assertTrue(i > 99);   // if_icmple
        assertTrue(i >= 99);  // if_icmplt
    }

    private static void testExtended() {
        Object o = null;
        assertTrue(o == null);  // ifnonnull
        assertFalse(o != null); // ifnull
    }

}
