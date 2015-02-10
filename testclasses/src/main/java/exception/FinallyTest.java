package exception;

public class FinallyTest {
    
    public static void main(String[] args) {
        int x = 1;
        try {
            bad();
            x = 100;
        } catch (Exception e) {
            x += 2;
        } finally {
            x *= 3;
        }
        
        if (x == (1 + 2) * 3) {
            System.out.println("OK!");
        }
    }
    
    private static void bad() {
        throw new RuntimeException("BAD!");
    }
    
}
