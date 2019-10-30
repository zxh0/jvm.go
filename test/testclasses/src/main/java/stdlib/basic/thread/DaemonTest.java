package stdlib.basic.thread;

import helper.UnitTestRunner;
import static helper.MyAssert.*;

public class DaemonTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(DaemonTest.class);
    }
    
//    @Test
    public void isDaemon() {
        Thread mainThread = Thread.currentThread();
        assertFalse(mainThread.isDaemon());
        
        Thread newThread = new Thread();
        assertFalse(newThread.isDaemon());
        
        newThread.setDaemon(true);
        assertTrue(newThread.isDaemon());
    }
    
}
