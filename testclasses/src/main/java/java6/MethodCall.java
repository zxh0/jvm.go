package java6;

public class MethodCall {
    
    public static void main(String[] args) {
        f(1, 2L, 3.14f, 2.71828, args);
        new MethodCall().g(1, 2L, 3.14f, 2.71828, args);
        System.out.println("OK");
    }
    
    private static void f(int a, long b, float c, double d, Object e) {
        int x = a;
        long y = b;
        float z = c;
        double u = d;
        Object v = e;
    }
    
    private void g(int a, long b, float c, double d, Object e) {
        int x = a;
        long y = b;
        float z = c;
        double u = d;
        Object v = e;
    }
    
}
