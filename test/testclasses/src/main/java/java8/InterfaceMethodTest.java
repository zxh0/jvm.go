package java8;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class InterfaceMethodTest {
    
    public interface If1 {
        static int x() {
            return 1;
        }
        default int y() {
            return 2;
        }
    }
    
    public static class Impl1 implements If1 {
        
    }
    
    public static class Impl2 implements If1 {        
        @Override
        public int y() {
            return 12;
        }
    }
    
    public static class Impl3 implements If1 {
        @Override
        public int y() {
            return 100 + If1.super.y();
        }
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(InterfaceMethodTest.class);
    }
    
    @Test
    public void staticMethod() {
        assertEquals(1, If1.x());
    }
    
    @Test
    public void defaultMethod() {
        assertEquals(2, new Impl1().y());
        assertEquals(12, new Impl2().y());
        assertEquals(102, new Impl3().y());
    }
    
}
