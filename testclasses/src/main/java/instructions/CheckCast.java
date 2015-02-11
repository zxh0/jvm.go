package instructions;

import java.io.Closeable;
import java.io.IOException;
import java.io.Serializable;

public class CheckCast {

    static class Sup implements Closeable {

        @Override
        public void close() throws IOException {
            System.out.println("close!");
        }

    }

    static class Sub extends Sup {

    }
    
    public static void main(String[] args) {
        sClass();
        sInterface();
        sArray();
        System.out.println("OK!");
    }
    
    private static void sClass() {
        Object s = new Sub();
        Sub t1 = (Sub) s;
        Sup t2 = (Sup) s;
        
    }
    
    private static void sInterface() {
        
    }
    
    private static void sArray() {
        
    }
    
}
