package java6.ex;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class InstructionExTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(InstructionExTest.class);
    }
    
    @Test(expected = ClassCastException.class)
    public void checkcast() {
        Object x = "String";
        Integer y = (Integer) x;
    }
    
    @Test(expected = NegativeArraySizeException.class)
    public void newarray() {
        int[] a = new int[-3];
    }
    
    @Test(expected = NegativeArraySizeException.class)
    public void anewarray() {
        Object[] a = new Object[-1];
    }
    
    @Test
    public void aload() {
        try {
            int[] a = {1};
            int x = a[2];
            fail();
        } catch (ArrayIndexOutOfBoundsException e) {
            assertEquals("2", e.getMessage());
        }
    }
    
    @Test
    public void astore() {
        try {
            int[] a = {};
            a[1] = 2;
            int x = a[1];
            fail();
        } catch (ArrayIndexOutOfBoundsException e) {
            assertEquals("1", e.getMessage());
        }
    }
    
    @Test(expected = ArithmeticException.class)
    public void idiv() {
        int x = 0;
        int y = 1 / x;
    }
    
    @Test(expected = ArithmeticException.class)
    public void irem() {
        int x = 0;
        int y = 1 % x;
    }
    
}
