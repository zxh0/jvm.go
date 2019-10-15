package jvm.jsr292;

import java.lang.invoke.MethodHandle;
import java.lang.invoke.MethodHandles;
import java.lang.invoke.MethodType;

import static helper.MyAssert.assertEquals;

public class MethodHandleTest {

    private static int x;

    public static String test() {
        return "hi";
    }

    public static void main(String[] args) throws Throwable {
        MethodHandles.Lookup lookup = MethodHandles.lookup();

        MethodType mt = MethodType.methodType(String.class);
        MethodHandle mh = lookup.findStatic(MethodHandleTest.class, "test", mt);
        assertEquals("hi", mh.invoke());
        assertEquals("hi", (String) mh.invokeExact());

        MethodHandle getter = lookup.findStaticGetter(MethodHandleTest.class, "x", int.class);
        MethodHandle setter = lookup.findStaticSetter(MethodHandleTest.class, "x", int.class);
        setter.invoke(100);
        assertEquals(100, getter.invoke());

        System.out.println("OK!");
    }

}
