package java6.reflection;

import org.junit.Test;
import static org.junit.Assert.*;
import libs.junit.UnitTestRunner;

public class PrimitiveClassTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(PrimitiveClassTest.class);
    }
    
    @Test
    public void test() {
        testPrimitiveClass(void.class,      "void");
        testPrimitiveClass(boolean.class,   "boolean");
        testPrimitiveClass(byte.class,      "byte");
        testPrimitiveClass(char.class,      "char");
        testPrimitiveClass(short.class,     "short");
        testPrimitiveClass(int.class,       "int");
        testPrimitiveClass(long.class,      "long");
        testPrimitiveClass(float.class,     "float");
        testPrimitiveClass(double.class,    "double");
    }
    
    private void testPrimitiveClass(Class<?> c, String name) {
        assertEquals(name, c.getName());
        assertEquals(null, c.getSuperclass());
        assertEquals(0, c.getFields().length);
        assertEquals(0, c.getDeclaredFields().length);
        assertEquals(0, c.getMethods().length);
        assertEquals(0, c.getDeclaredMethods().length);
    }
    
}
