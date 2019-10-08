package jvm;

public class InheritanceTest {

    static class Super {
        int x;
        long y;
    }
    static class Sub extends Super {
        float a;
        double b;
    }

    public static void main(String[] args) {
        Sub sub = new Sub();
        sub.x = 1;
        sub.y = 2L;
        sub.a = 3.14f;
        sub.b = 2.71828;
        
        int x = sub.x;
        long y = sub.y;
        float a = sub.a;
        double b = sub.b;
        
        Super sup = sub;
        long z = sup.x + sup.y;
        
        System.out.println("OK!");
    }
    
}

