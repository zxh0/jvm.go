package exception;

public class StackTraceTest {
    
    public static void main(String[] args) {
        try {
            foo();
        } catch (Exception e) {
            e.printStackTrace(System.err);
        }
    }
    
    private static void foo() {
        bar();
    }
    
    private static void bar() {
        throw new RuntimeException("BAD!");
    }
    
}
