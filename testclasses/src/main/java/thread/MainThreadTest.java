package thread;

import junit.UnitTestRunner;
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
    }
    
}
