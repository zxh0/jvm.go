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
    
    @Test(expected = ClassCastException.class)
    public void checkcast() {
        Object x = "String";
        Integer y = (Integer) x;
    }
    
    //@Test(expected = NegativeArraySizeException.class)
    public void newarray() {
        int[] a = new int[-3];
    }
    
    //@Test(expected = NegativeArraySizeException.class)
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
    
    private Object nullObj() {
        return null;
    }
    
}
