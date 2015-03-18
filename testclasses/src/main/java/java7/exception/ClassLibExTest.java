package java7.exception;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ClassLibExTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ClassLibExTest.class);
    }
    
    @Test
    public void threadSleep() throws InterruptedException {
        try {
            Thread.sleep(-1);
            fail();
        } catch (IllegalArgumentException e) {
            assertEquals("timeout value is negative", e.getMessage());
        }
    }
    
}
