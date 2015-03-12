package jvmgo.java8;

import jvmgo.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class InterfaceMethodTest {
    
    public interface IF1 {
        public static int x() {
            return 1;
        }
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(InterfaceMethodTest.class);
    }
    
    @Test
    public void interfaceStaticMethod() {
        assertEquals(1, IF1.x());
    }
    
}
