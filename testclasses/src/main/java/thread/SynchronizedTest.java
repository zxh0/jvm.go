package thread;

public class SynchronizedTest {
    
    public static void main(String[] args) {
        synchronized (SynchronizedTest.class) {
            test();
        }
    }
    
    private static synchronized void test() {
        new SynchronizedTest().foo();
    }
    
    private synchronized void foo() {
        bar();
    }
    
    private void bar() {
        System.out.println("OK!");
    }
    
}
