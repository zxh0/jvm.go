package java6.reflection;

import java.lang.reflect.Method;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ClassTest implements Runnable {

    private int a;
    public double b;
    static boolean z;
    
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(ClassTest.class);
    }
    
    @Override
    @Test
    public void run() throws RuntimeException {
        //System.out.println("run!");
    }
    
    @Test
    public void _package() {
        assertEquals("reflection", getClass().getPackage().getName());
    }
    
    @Test
    public void _class() {
        Class<?> c = ClassTest.class;
        assertEquals("reflection.ClassTest", c.getName());
        assertEquals(Object.class, c.getSuperclass());
        assertArrayEquals(new Class<?>[]{Runnable.class}, c.getInterfaces());
        assertEquals(1, c.getFields().length);
        assertEquals(3, c.getDeclaredFields().length);
        assertEquals(14, c.getMethods().length);
        assertEquals(5, c.getDeclaredMethods().length);
    }

    @Test
    public void method() throws Exception {
        Method main = ClassTest.class.getMethod("main", String[].class);
        assertArrayEquals(new Class<?>[]{Exception.class}, main.getExceptionTypes());
        assertArrayEquals(new Class<?>[]{String[].class}, main.getParameterTypes());
        assertEquals(0, main.getDeclaredAnnotations().length);
        
        Method run = ClassTest.class.getMethod("run");
        assertEquals(1, run.getDeclaredAnnotations().length);
    }
    
}
