package instructions;

import java.io.Closeable;
import java.io.IOException;
import java.io.Serializable;

public class CheckCast {
    
    public static void main(String[] args) {
        Object x = "string";
        String y = (String) x;
        System.out.println("OK1!");
        
        Object a = args;
        String[] b = (String[]) a;
        System.out.println("OK2!");
        
        CharSequence cs = (CharSequence) x;
        Serializable s = (Serializable) x;
        System.out.println("OK3!");
        
        Object sub = new Sub();
        Sup sup = (Sup) sub;
        Closeable c = (Closeable) sub;
        
        AutoCloseable ac = (AutoCloseable) sub;
        System.out.println("OK4!");
        
    }
    
}

class Sup implements Closeable {

    @Override
    public void close() throws IOException {
        System.out.println("close!");
    }
    
}

class Sub extends Sup {
    
}
