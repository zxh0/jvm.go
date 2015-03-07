package exception;

import junit.UnitTestRunner;
import org.junit.Test;

public class JvmExTest {
    
    int i;
    
    public static void main(String[] args) {
        UnitTestRunner.run(JvmExTest.class);
    }
    
    @Test(expected = NullPointerException.class)
    public void arraylength() {
        int[] x = (int[]) nullObj();
        int y = x.length;
    }
    
    @Test(expected = NullPointerException.class)
    public void athrow() throws Exception {
        Exception x = (Exception) nullObj();
        throw x;
    }
    
    @Test(expected = NullPointerException.class)
    public void getfield() {
        JvmExTest x = (JvmExTest) nullObj();
        int y = x.i;
    }
    
    @Test(expected = NullPointerException.class)
    public void monitorenter() {
        Object x = nullObj();
        synchronized(x) {
            System.out.println("BAD!");
        }
    }
    
    // ClassCastException
    static void checkcast() {
        Object x = "String";
        Integer y = (Integer) x;
    }
    
    // NegativeArraySizeException
    static void newarray() {
        int[] a = new int[-3];
    }
    
    // NegativeArraySizeException
    static void anewarray() {
        Object[] a = new Object[-1];
    }
    
    // ArrayIndexOutOfBoundsException
    static void aload() {
        int[] a = {1};
        int x = a[2];
    }
    
    // ArrayIndexOutOfBoundsException
    static void astore() {
        int[] a = {};
        a[1] = 2;
        int x = a[1];
    }
    
    // ArithmeticException
    static void idiv() {
        int x = 0;
        int y = 1 / x;
    }
    
    // ArithmeticException
    static void irem() {
        int x = 0;
        int y = 1 % x;
    }
    
    private Object nullObj() {
        return null;
    }
    
}
