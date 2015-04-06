package java6.thread;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class MainThreadTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(MainThreadTest.class);
    }
    
    @Test
    public void mainThread() {
        Thread mainThread = Thread.currentThread();
        assertEquals("main", mainThread.getName());
        assertTrue("isAlive", mainThread.isAlive());
        assertFalse("isDaemon", mainThread.isDaemon());
        //System.out.println(mainThread.getThreadGroup());
    }
    
}
