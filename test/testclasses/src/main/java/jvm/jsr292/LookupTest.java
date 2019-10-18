package jvm.jsr292;

import java.lang.invoke.MethodHandles;
import java.lang.invoke.MethodType;

import static helper.MyAssert.assertEquals;

public class LookupTest {

    public LookupTest(String[] x) {}
    private void foo(String[] x) {}
    public void bar(String[] x) {}
    public int f1;
    public static long f2;

    public static void main(String[] args) throws Exception {
        MethodHandles.Lookup lookup = MethodHandles.lookup();
        MethodType mt = MethodType.methodType(void.class, String[].class);

        assertEquals("MethodHandle(String[])void",            lookup.findStatic(LookupTest.class, "main", mt).toString());
        assertEquals("MethodHandle(String[])LookupTest",      lookup.findConstructor(LookupTest.class, mt).toString());
        assertEquals("MethodHandle(LookupTest,String[])void", lookup.findSpecial(LookupTest.class, "foo", mt, LookupTest.class).toString());
        assertEquals("MethodHandle(LookupTest,String[])void", lookup.findVirtual(LookupTest.class, "bar", mt).toString());
        assertEquals("MethodHandle(LookupTest)int",           lookup.findGetter(LookupTest.class, "f1", int.class).toString());
        assertEquals("MethodHandle(LookupTest,int)void",      lookup.findSetter(LookupTest.class, "f1", int.class).toString());
        assertEquals("MethodHandle()long",                    lookup.findStaticGetter(LookupTest.class, "f2", long.class).toString());
        assertEquals("MethodHandle(long)void",                lookup.findStaticSetter(LookupTest.class, "f2", long.class).toString());

        //System.out.println(lookup.findVirtual(LookupTest.class, "toString", MethodType.methodType(String.class)));
        System.out.println("OK!");
    }

}
