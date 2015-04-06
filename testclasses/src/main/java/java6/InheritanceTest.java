package java6;

public class InheritanceTest {
    
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

class Super {
    int x;
    long y;
}
class Sub extends Super {
    float a;
    double b;
}
