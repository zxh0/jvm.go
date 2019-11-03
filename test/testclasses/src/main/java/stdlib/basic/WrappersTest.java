package stdlib.basic;

import static helper.MyAssert.*;

public class WrappersTest implements Runnable {
    
    public static void main(String[] args) {
        new WrappersTest().run();
    }

    @Override
    public void run() {
        assertEquals(1076754509, Float.floatToRawIntBits(2.71828f));
        assertEquals(2.71828f, Float.intBitsToFloat(1076754509));
        assertEquals(4614253070214989087L, Double.doubleToRawLongBits(3.14));
        assertEquals(3.14, Double.longBitsToDouble(4614253070214989087L));
    }

}
