package jvm.cls;

import static helper.MyAssert.*;

// A constant variable is a final variable of primitive type or type String
// that is initialized with a constant expression (§15.28).
public class FieldConstantStaticVarTest implements Runnable {
    
    public static final boolean z = true;
    public static final byte b = 125;
    public static final char c = 'c';
    public static final short s = 300;
    public static final int x = 100;
    public static final int y = x + 18;
    public static final long j = 1L;
    public static final float f = 1.5f;
    public static final double d = 2.5;
    public static final String str1 = "hello";
    public static final String str2 = str1 + " world!";

    public static void main(String[] args) {
        new FieldConstantStaticVarTest().run();
    }

    @Override
    public void run() {
        assertEquals(true, FieldConstantStaticVarTest.z);
        assertEquals(true, getFieldValue("z"));
        assertEquals((byte)125, FieldConstantStaticVarTest.b);
        assertEquals((byte)125, getFieldValue("b"));
        assertEquals('c', FieldConstantStaticVarTest.c);
        assertEquals('c', getFieldValue("c"));
        assertEquals((short)300, FieldConstantStaticVarTest.s);
        assertEquals((short)300, getFieldValue("s"));
        assertEquals(100, FieldConstantStaticVarTest.x);
        assertEquals(100, getFieldValue("x"));
        assertEquals(118, FieldConstantStaticVarTest.y);
        assertEquals(118, getFieldValue("y"));
        assertEquals(1L, FieldConstantStaticVarTest.j);
        assertEquals(1L, getFieldValue("j"));
        assertEquals(1.5f, FieldConstantStaticVarTest.f);
        assertEquals(1.5f, getFieldValue("f"));
        assertEquals(2.5, FieldConstantStaticVarTest.d);
        assertEquals(2.5, getFieldValue("d"));
        assertEquals("hello", FieldConstantStaticVarTest.str1);
        assertEquals("hello", getFieldValue("str1"));
        assertEquals("hello world!", FieldConstantStaticVarTest.str2);
        assertEquals("hello world!", getFieldValue("str2"));
    }

    private static Object getFieldValue(String name) {
        try {
            return FieldConstantStaticVarTest.class.getField(name).get(null);
        } catch (ReflectiveOperationException e) {
            throw new RuntimeException(e);
        }
    }

}
