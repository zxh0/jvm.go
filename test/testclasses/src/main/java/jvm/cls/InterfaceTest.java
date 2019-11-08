package jvm.cls;

import static helper.MyAssert.*;

public class InterfaceTest implements Runnable {

    private interface I {
        static int m1() { return 1; }
        default int m2() { return 2; }
        private int m3() { return 3; }
        int m4();
    }

    private static class C implements I {
        public int m4() {
            return 4;
        }
    }

    public static void main(String[] args) {
        new InterfaceTest().run();
    }

    @Override
    public void run() {
        assertEquals(1, I.m1()); // invokestatic I.m1
        assertEquals(2, new C().m2()); // invokevirtual C.m2
    }

}
