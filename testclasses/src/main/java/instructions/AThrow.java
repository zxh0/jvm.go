package instructions;

public class AThrow {

    public static void main(String[] args) {
        try {
            throw new RuntimeException("**RE!");
        } catch (RuntimeException e) {
            System.out.println(e.getMessage());
        }
    }
    
}
