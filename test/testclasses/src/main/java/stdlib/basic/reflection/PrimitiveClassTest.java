package stdlib.basic.reflection;

import static helper.MyAssert.*;

public class PrimitiveClassTest implements Runnable {
    
    public static void main(String[] args) {
        new PrimitiveClassTest().run();
    }

    @Override
    public void run() {
        testPrimitiveClass(void.class,    "void");
        testPrimitiveClass(boolean.class, "boolean");
        testPrimitiveClass(byte.class,    "byte");
        testPrimitiveClass(char.class,    "char");
        testPrimitiveClass(short.class,   "short");
        testPrimitiveClass(int.class,     "int");
        testPrimitiveClass(long.class,    "long");
        testPrimitiveClass(float.class,   "float");
        testPrimitiveClass(double.class,  "double");
    }

    private void testPrimitiveClass(Class<?> c, String name) {
        assertEquals(name, c.getName());
        assertNull(c.getSuperclass());
        assertEquals(0, c.getFields().length);
        assertEquals(0, c.getDeclaredFields().length);
        assertEquals(0, c.getMethods().length);
        assertEquals(0, c.getDeclaredMethods().length);
    }

}
