package jls8;

import java.io.PrintWriter;
import java.io.StringWriter;

public class StringOut extends PrintWriter {
    
    public StringOut () {
        super(new StringWriter());
    }
    
    @Override
    public String toString() {
        String str = out.toString();
        return str.replaceAll("\r\n", "\n");
    }
    
}
