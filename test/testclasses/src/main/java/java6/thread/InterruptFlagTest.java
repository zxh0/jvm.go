package java6.thread;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class InterruptFlagTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(InterruptFlagTest.class);
    }
    
    @Test
    public void interruptFlag() {
        Thread t = Thread.currentThread();
        assertFalse(t.isInterrupted());
        
        t.interrupt();
        assertTrue(t.isInterrupted());
        assertTrue(t.isInterrupted());
        
        assertTrue(Thread.interrupted());
        assertFalse(Thread.interrupted());
        assertFalse(t.isInterrupted());
    }
    
}
