package java7.exception;

import libs.junit.UnitTestRunner;
import org.junit.Test;

public class ClassLibExTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ClassLibExTest.class);
    }
    
    @Test
    public void threadSleep() throws InterruptedException {
        Thread.sleep(-1);
    }
    
}
