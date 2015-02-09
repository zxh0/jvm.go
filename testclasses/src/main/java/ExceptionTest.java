public class ExceptionTest {
    
    public static void main(String[] args) {
        foo();
    }
    
    private static void foo() {
        bar();
    }
    
    private static void bar() {
        throw new RuntimeException("BAD!");
    }
    
}
