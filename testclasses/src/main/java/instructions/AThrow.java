package instructions;

public class AThrow {

    public static void main(String[] args) {
        try {
            foo();
        } catch (RuntimeException e) {
            System.out.println("main!!");
            System.out.println(e.getMessage());
        }
    }
    
    static void foo() {
        try {
            throw new RuntimeException("foo!");
        } catch (RuntimeException e) {
            System.out.println(e.getMessage());
            throw new RuntimeException("bar!");
        }
    }
    
}
