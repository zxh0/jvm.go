package java6.field;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

// A constant variable is a final variable of primitive type or type String
// that is initialized with a constant expression (§15.28).
public class ConstantStaticFieldsTest {
    
    public static final boolean z = true;
    public static final byte b = 125;
    public static final char c = 'c';
    public static final short s = 300;
    public static final int x = 100;
    public static final int y = x + 18;
    public static final long j = 1L;
    public static final float f = 3.14f;
    public static final double d = 2.71828;
    public static final String str1 = "hello";
    public static final String str2 = str1 + " world!";
    
    @Test
    public void test() {
        assertEquals(true, ConstantStaticFieldsTest.z);
        assertEquals(true, getFieldValue("z"));
        assertEquals((byte)125, ConstantStaticFieldsTest.b);
        assertEquals((byte)125, getFieldValue("b"));
        assertEquals('c', ConstantStaticFieldsTest.c);
        assertEquals('c', getFieldValue("c"));
        assertEquals((short)300, ConstantStaticFieldsTest.s);
        assertEquals((short)300, getFieldValue("s"));
        assertEquals(100, ConstantStaticFieldsTest.x);
        assertEquals(100, getFieldValue("x"));
        assertEquals(118, ConstantStaticFieldsTest.y);
        assertEquals(118, getFieldValue("y"));
        assertEquals(1L, ConstantStaticFieldsTest.j);
        assertEquals(1L, getFieldValue("j"));
        assertEquals(3.14f, ConstantStaticFieldsTest.f, 0.1);
        assertEquals(3.14f, getFieldValue("f"));
        assertEquals(2.71828, ConstantStaticFieldsTest.d, 0.1);
        assertEquals(2.71828, getFieldValue("d"));
        assertEquals("hello", ConstantStaticFieldsTest.str1);
        assertEquals("hello", getFieldValue("str1"));
        assertEquals("hello world!", ConstantStaticFieldsTest.str2);
        assertEquals("hello world!", getFieldValue("str2"));
    }
    
    private static Object getFieldValue(String name) {
        try {
            return ConstantStaticFieldsTest.class.getField(name).get(null);
        } catch (ReflectiveOperationException e) {
            throw new RuntimeException(e);
        }
    }
    
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(ConstantStaticFieldsTest.class);
    }
    
}
