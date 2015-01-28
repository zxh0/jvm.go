public class InheritanceTest {
    
    public static void main(String[] args) {
        Sub sub = new Sub();
        sub.x = 1;
        sub.y = 2L;
        sub.a = 3.14f;
        sub.b = 2.71828;
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
