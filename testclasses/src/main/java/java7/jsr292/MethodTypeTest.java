package java7.jsr292;

import java.lang.invoke.MethodHandle;
import java.lang.invoke.MethodType;

public class MethodTypeTest {
    
    public static void main(String[] args) throws Exception {
        //MethodType.fromMethodDescriptorString("()V", null);
        MethodType.methodType(Void.class);
        //System.out.println(MethodHandle.class);
    }
    
}
