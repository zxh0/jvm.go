package stdlib.basic.reflection;

import java.io.Serializable;

import static helper.MyAssert.*;

public class ArrayClassTest implements Runnable {
    
    public static void main(String[] args) {
        new ArrayClassTest().run();
    }

    @Override
    public void run() {
        testArrayClasses();
        testGetComponentType();
    }

    private static void testArrayClasses() {
        testArrayClass(boolean[].class, "[Z");
        testArrayClass(byte[].class,    "[B");
        testArrayClass(char[].class,    "[C");
        testArrayClass(short[].class,   "[S");
        testArrayClass(int[].class,     "[I");
        testArrayClass(long[].class,    "[J");
        testArrayClass(float[].class,   "[F");
        testArrayClass(double[].class,  "[D");
        testArrayClass(int[][].class,   "[[I");
        testArrayClass(Object[].class,  "[Ljava.lang.Object;");
        testArrayClass(Object[][].class,"[[Ljava.lang.Object;");
    }

    private static void testArrayClass(Class<?> c, String name) {
        assertEquals(name, c.getName());
        assertEquals(Object.class, c.getSuperclass());
        assertArrayEquals(new Class<?>[]{Cloneable.class, Serializable.class}, c.getInterfaces());
        assertEquals(0, c.getFields().length);
        assertEquals(0, c.getDeclaredFields().length);
        assertEquals(9, c.getMethods().length);
        assertEquals(0, c.getDeclaredMethods().length);
    }

    private static void testGetComponentType() {
        assertSame(int.class, int[].class.getComponentType());
        assertSame(String.class, new String[0].getClass().getComponentType());
    }

}
