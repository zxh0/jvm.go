package reflection;

import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.util.Arrays;
import java.util.Properties;
import sun.reflect.Reflection;

public class ClassTest implements Runnable {
    static {
        //System.out.println("clinit");
    }

    private int a;
    public double b;
    static boolean z;
    
    public ClassTest() {
        System.out.println("init");
    }
    
    public static void main(String[] args) {
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
        
        //Object x = new Reflection();
        //new Properties();
        
//        for (Method m : c.getMethods()) {
//            System.out.println("method: " + m.getName());
//        }
//        for (Method m : c.getDeclaredMethods()) {
//            System.out.println("declaredMethods: " + m.getName());
//        }
    }

    @Override
    public void run() {
        System.out.println("run!");
    }
    
}
