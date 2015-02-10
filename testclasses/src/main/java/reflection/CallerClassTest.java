package reflection;

import sun.reflect.Reflection;

public class CallerClassTest {
    
    public static void main(String[] args) {
        Foo.test();
    }
    
    static class Foo {
        static void test() {
            Bar.test();
        }
    }
    
    static class Bar {
        static void test() {
            System.out.println(Reflection.getCallerClass(0).getName());
            System.out.println(Reflection.getCallerClass(1).getName());
            System.out.println(Reflection.getCallerClass(2).getName());
            System.out.println(Reflection.getCallerClass(3).getName());
        }
    }
}
