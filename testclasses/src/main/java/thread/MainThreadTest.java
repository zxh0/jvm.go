package thread;

import junit.UnitTestRunner;
import org.junit.Test;

public class MainThreadTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(MainThreadTest.class);
    }
    
    @Test
    public void mainThread() {
        Thread mainThread = Thread.currentThread();
        System.out.println(mainThread);
    }
    
}
