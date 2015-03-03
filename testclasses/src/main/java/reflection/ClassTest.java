package reflection;

import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.util.Arrays;

public class ClassTest implements Runnable {

    private int a;
    public double b;
    static boolean z;
    
    public ClassTest() {
        System.out.println("init");
    }
    
    @Override
    public void run() throws RuntimeException {
        System.out.println("run!");
    }
    
    public static void main(String[] args) throws Exception {
        _class();
        method();
    }
    
    public static void _class() {
        //System.out.println(Object.class.getSuperclass());
        
        Class<?> c = ClassTest.class;
        System.out.println(c.getName());
        System.out.println("superclass: " + c.getSuperclass().getName());
        for (Class<?> ic : c.getInterfaces()) {
            System.out.println("interfaces: " + ic.getName());
        }
        for (Field f : c.getFields()) {
            System.out.println("field: " + f.getName());
            System.out.println("field type: " + f.getType());
        }
        for (Field f : c.getDeclaredFields()) {
            System.out.println("declaredFields:" + f.getName());
        }
        for (Method m : c.getMethods()) {
            System.out.println("method: " + m.getName());
        }
        for (Method m : c.getDeclaredMethods()) {
            System.out.println("declaredMethods: " + m.getName());
        }
    }

    public static void method() throws Exception {
        Method run = ClassTest.class.getMethod("run");
        System.out.println(Arrays.toString(run.getExceptionTypes()));
        System.out.println(run.getParameterTypes().length);
    }
    
}
