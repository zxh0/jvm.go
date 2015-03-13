package java8;

import java7.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class InterfaceMethodTest {
    
    public interface IF1 {
        static int x() {
            return 1;
        }
        default int y() {
            return 2;
        }
    }
    
    public class IMPL1 implements IF1 {
        
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(InterfaceMethodTest.class);
    }
    
    @Test
    public void staticMethod() {
        assertEquals(1, IF1.x());
    }
    
    @Test
    public void defaultMethod() {
        assertEquals(2, new IMPL1().y());
    }
    
}
