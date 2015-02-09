package exception;

public class ExceptionTest {
    
    public static void main(String[] args) {
        foo();
    }
    
    private static void foo() {
        bar();
    }
    
    private static void bar() {
        try {
            throw new RuntimeException("BAD!");
        } catch (Exception e) {
            System.out.println("111");
        } finally {
            System.out.println("222");
        }
    }
    
}
