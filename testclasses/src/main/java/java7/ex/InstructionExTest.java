package java7.ex;

import libs.junit.UnitTestRunner;
import org.junit.Test;

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
    
    @Test(expected = ArrayIndexOutOfBoundsException.class)
    public void aload() {
        int[] a = {1};
        int x = a[2];
    }
    
    @Test(expected = ArrayIndexOutOfBoundsException.class)
    public void astore() {
        int[] a = {};
        a[1] = 2;
        int x = a[1];
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
