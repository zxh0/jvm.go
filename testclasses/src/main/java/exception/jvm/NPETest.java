package exception.jvm;

public class NPETest {
    
    public static void main(String[] args) {
        test();
    }
    
    private static void test() {
        new NPETest().foo();
    }
    
    private void foo() {
        bar();
    }
    
    private void bar() {
        Object x = null;
        synchronized(x) {
            System.out.println("BAD!");
        }
    }
    
}
