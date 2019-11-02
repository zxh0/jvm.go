package stdlib.basic.reflection;

import java.lang.reflect.Method;
import static helper.MyAssert.*;

public class ClassTest implements Runnable {

    private int a;
    public double b;
    static boolean z;

    @SafeVarargs
    public static void main(String ...args) throws Exception {
        new ClassTest().run();
    }

    @Override
    public void run() {
        testGetPackage();
        testGetClass();
        try {
            testGetMethod();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    public void testGetPackage() {
        assertEquals("stdlib.basic.reflection", getClass().getPackage().getName());
    }

    public void testGetClass() {
        Class<?> c = ClassTest.class;
        assertEquals("stdlib.basic.reflection.ClassTest", c.getName());
        assertEquals(Object.class, c.getSuperclass());
        assertArrayEquals(new Class<?>[]{Runnable.class}, c.getInterfaces());
        assertEquals(1, c.getFields().length);
        assertEquals(3, c.getDeclaredFields().length);
        assertEquals(14, c.getMethods().length);
        assertEquals(5, c.getDeclaredMethods().length);
    }

    public void testGetMethod() throws Exception {
        Method main = ClassTest.class.getMethod("main", String[].class);
        assertArrayEquals(new Class<?>[]{Exception.class}, main.getExceptionTypes());
        assertArrayEquals(new Class<?>[]{String[].class}, main.getParameterTypes());
        assertEquals(1, main.getDeclaredAnnotations().length);
        
        Method run = ClassTest.class.getMethod("run");
        assertEquals(0, run.getDeclaredAnnotations().length);
    }

}
