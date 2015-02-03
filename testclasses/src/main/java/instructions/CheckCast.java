package instructions;

import java.io.Serializable;

public class CheckCast {
    
    public static void main(String[] args) {
        Object x = "string";
        String y = (String) x;
        
        Object a = args;
        String[] b = (String[]) a;
        
        CharSequence cs = (CharSequence) x;
        Serializable s = (Serializable) x;
        
        System.out.println("OK!");
    }
    
}
