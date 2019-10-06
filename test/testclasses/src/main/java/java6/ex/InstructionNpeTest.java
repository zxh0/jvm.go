package java6.ex;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static helper.TestHelper.nullObj;

public class InstructionNpeTest {
    
    private int i;
    
    public static void main(String[] args) {
        UnitTestRunner.run(InstructionNpeTest.class);
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
        InstructionNpeTest x = (InstructionNpeTest) nullObj();
        int y = x.i;
    }
    
    @Test(expected = NullPointerException.class)
    public void monitorenter() {
        Object x = nullObj();
        synchronized(x) {
            System.out.println("BAD!");
        }
    }
    
    @Test(expected = NullPointerException.class)
    public void invokevirtual() {
        Object x = nullObj();
        x.toString();
    }
    
}
