package instructions;

public class ANewArray {
    
    public static void main(String[] args) {
        Object[] arr = new Object[8];
        if (arr.length == 8) {
            System.out.println("OK!");
        } else {
            System.out.println("Fail!");
        }
    }
    
}
