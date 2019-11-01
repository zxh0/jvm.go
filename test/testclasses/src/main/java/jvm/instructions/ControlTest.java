package jvm.instructions;

import static helper.MyAssert.*;

public class ControlTest implements Runnable {

    public static void main(String[] args) {
        new ControlTest().run();
    }

    @Override
    public void run() {
        testTableSwitch();
        testLookupSwitch();
    }

    private static void testTableSwitch() {
        assertEquals(0, chooseNear(0));
        assertEquals(1, chooseNear(1));
        assertEquals(2, chooseNear(2));
        assertEquals(-1, chooseNear(3));
    }

    private static void testLookupSwitch() {
        assertEquals(-1, chooseFar(-100));
        assertEquals(0, chooseFar(0));
        assertEquals(1, chooseFar(100));
        assertEquals(-2, chooseFar(1));
    }

    // tableswitch
    private static int chooseNear(int i) {
        switch (i) {
            case 0:  return  0;
            case 1:  return  1;
            case 2:  return  2;
            default: return -1;
        }
    }

    // lookupswitch
    private static int chooseFar(int i) {
        switch (i) {
            case -100: return -1;
            case 0:    return  0;
            case 100:  return  1;
            default:   return -2;
        }
    }

}
