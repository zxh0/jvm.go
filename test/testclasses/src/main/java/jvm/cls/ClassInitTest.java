package jvm.cls;

import static helper.MyAssert.*;

public class ClassInitTest implements Runnable {

    private static int a;

    private static class A {
        static {
            a = 100;
        }
    }

    public static void main(String[] args) {
        new ClassInitTest().run();
    }

    @Override
    public void run() {
        A[] arr = new A[8];
        assertEquals(0, a);
    }

}
