package java7.wrapper;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class FloatTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(FloatTest.class);
    }
    
    @Test
    public void doubleToRawLongBits() {
        assertEquals(1076754509, Float.floatToRawIntBits(2.71828f));
    }
    
    @Test
    public void longBitsToDouble() {
        assertEquals(2.71828f, Float.intBitsToFloat(1076754509), 0);
    }
    
}
