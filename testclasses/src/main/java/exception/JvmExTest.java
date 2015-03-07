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
        int[] x = null;
        int y = x.length;
    }
    
    static void athrow(Exception ex) throws Exception {
        throw ex;
    }
    
    static void getfield(JvmExTest x) {
        int y = x.i;
    }
    
    static void checkcast() {
        Object x = "String";
        Integer y = (Integer) x;
    }
    
    static void newarray() {
        int[] a = new int[-3];
    }
    
    static void anewarray() {
        Object[] a = new Object[-1];
    }
    
    static void aload() {
        int[] a = {1};
        int x = a[2];
    }
    
    static void astore() {
        int[] a = {};
        a[1] = 2;
        int x = a[1];
    }
    
    static void idiv() {
        int x = 0;
        int y = 1 / x;
    }
    
    static void irem() {
        int x = 0;
        int y = 1 % x;
    }
    
    static void monitorenter() {
        Object x = null;
        synchronized(x) {
            System.out.println("BAD!");
        }
    }
    
}
