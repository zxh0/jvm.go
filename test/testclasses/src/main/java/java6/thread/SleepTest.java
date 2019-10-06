package java6.thread;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class SleepTest {
    
    public static void main(String[] args)  {
        UnitTestRunner.run(SleepTest.class);
    }
    
    @Test
    public void sleep() throws InterruptedException {
        long beforeSleep = System.currentTimeMillis();
        Thread.sleep(100);
        long afterSleep = System.currentTimeMillis();
        assertTrue(afterSleep - beforeSleep >= 100);
    }
    
}
