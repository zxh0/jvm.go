package jsr292;

import java.lang.invoke.MethodType;

public class MethodTypeTest {
    
    public static void main(String[] args) throws Exception {
        MethodType.fromMethodDescriptorString("()V", null);
    }
    
}
