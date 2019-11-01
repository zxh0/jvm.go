package jvm.instructions;

import static helper.MyAssert.*;

public class RefFieldTest implements Runnable {

    private static int f1;
    private double f2;

    public static void main(String[] args) {
        new RefFieldTest().run();
    }

    @Override
    public void run() {
        RefFieldTest.f1 = 100;              // putstatic
        assertEquals(100, RefFieldTest.f1); // getstatic
        this.f2 = 1.5;                      // putfield
        assertEquals(1.5, this.f2);         // getfield
    }

}
