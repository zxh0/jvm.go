package exception;

public class UncaughtTest {
    
    public static void main(String[] args) {
        foo();
    }
    
    private static void foo() {
        bar();
    }
    
    private static void bar() {
        bad();
    }
    
    private static void bad() {
        throw new RuntimeException("BAD!");
    }
    
}
