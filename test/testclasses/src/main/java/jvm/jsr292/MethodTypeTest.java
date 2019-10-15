package jvm.jsr292;

import java.util.Arrays;

import static java.lang.invoke.MethodType.*;
import static helper.MyAssert.assertEquals;

public class MethodTypeTest {

    public static void main(String[] args) {
        assertEquals("()void",           methodType(void.class).toString());                              // methodType(Class<?> rtype)
        assertEquals("(String)Object",   methodType(Object.class, String.class).toString());              // methodType(Class<?> rtype, Class<?> ptype0)
        assertEquals("(int,long)String", methodType(String.class, int.class, long.class).toString());     // methodType(Class<?> rtype, Class<?> ptype0, Class<?>... ptypes)
        assertEquals("(int)String",      methodType(String.class, new Class<?>[]{int.class}).toString()); // methodType(Class<?> rtype, Class<?>[] ptypes)
        assertEquals("(int)String",      methodType(String.class, Arrays.asList(int.class)).toString());  // methodType(Class<?> rtype, List<Class<?>> ptypes)
        assertEquals("()void", fromMethodDescriptorString("()V", null).toString());
        System.out.println("OK!");
    }

}
