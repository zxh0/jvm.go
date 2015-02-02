package instructions;

public class CheckCast {
    
    public static void main(String[] args) {
        Object x = "string";
        String y = (String) x;
        
        Object a = args;
        String[] b = (String[]) a;
        
        System.out.println("OK!");
    }
    
}
