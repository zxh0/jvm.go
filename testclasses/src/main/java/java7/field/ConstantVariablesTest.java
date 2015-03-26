package java7.field;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

// A constant variable is a final variable of primitive type or type String
// that is initialized with a constant expression (§15.28).
public class ConstantVariablesTest {
    
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
        fail("todo");
//        assertEquals(true, ConstantVariablesTest.z);
//        assertEquals(true, getFieldValue("z"));
//        assertEquals((byte)125, ConstantVariablesTest.b);
//        assertEquals((byte)125, getFieldValue("b"));
//        assertEquals('c', ConstantVariablesTest.c);
//        assertEquals('c', getFieldValue("c"));
//        assertEquals((short)300, ConstantVariablesTest.s);
//        assertEquals((short)300, getFieldValue("s"));
//        assertEquals(100, ConstantVariablesTest.x);
//        assertEquals(100, getFieldValue("x"));
//        assertEquals(118, ConstantVariablesTest.y);
//        assertEquals(118, getFieldValue("y"));
//        assertEquals(1L, ConstantVariablesTest.j);
//        assertEquals(1L, getFieldValue("j"));
//        assertEquals(3.14f, ConstantVariablesTest.f, 0.1);
//        assertEquals(3.14f, getFieldValue("f"));
//        assertEquals(2.71828, ConstantVariablesTest.d, 0.1);
//        assertEquals(2.71828, getFieldValue("d"));
//        assertEquals("hello", ConstantVariablesTest.str1);
//        assertEquals("hello", getFieldValue("str1"));
//        assertEquals("hello world!", ConstantVariablesTest.str2);
//        assertEquals("hello world!", getFieldValue("str2"));
    }
    
    private static Object getFieldValue(String name) {
        try {
            return ConstantVariablesTest.class.getField(name).get(null);
        } catch (ReflectiveOperationException e) {
            throw new RuntimeException(e);
        }
    }
    
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(ConstantVariablesTest.class);
    }
    
}
