package java6.wrapper;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class DoubleTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(DoubleTest.class);
    }
    
    @Test
    public void doubleToRawLongBits() {
        assertEquals(4614253070214989087L, Double.doubleToRawLongBits(3.14));
    }
    
    @Test
    public void longBitsToDouble() {
        assertEquals(3.14, Double.longBitsToDouble(4614253070214989087L), 0);
    }
    
}
