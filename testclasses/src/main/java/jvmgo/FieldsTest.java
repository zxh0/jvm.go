package jvmgo;

public class FieldsTest {
    
    static class Sup {
        static int x;
        int a;
    }
    
    static class Sub extends Sup {
        static int y;
        int b;
    }
    
    public static void main(String[] args) {
        staticFields();
        instanceFields();
    }
    
    private static void staticFields() {
        int z = Sub.x + Sub.y;
        z += 100;
        Sub.y = z;
        Sub.x = z;
        if (Sub.x == 100 && Sub.y == 100) {
            System.out.println("OK! staticFields");
        }
    }
    
    private static void instanceFields() {
        Sub sub = new Sub();
        int c = sub.a + sub.b;
        c += 100;
        sub.a = c;
        sub.b = c;
        if (sub.a == 100 && sub.b == 100) {
            System.out.println("OK! instanceFields");
        }
    }
    
}
