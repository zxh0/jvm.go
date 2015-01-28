public class FieldsTest {
    
    static int x;
    static long y;
    float a;
    double b;
    
    public static void main(String[] args) {
        FieldsTest.x = 100;
        FieldsTest.y = 200L;
        long z = FieldsTest.x / FieldsTest.y;
        FieldsTest t = new FieldsTest();
        t.a = 3.14f;
        t.b = 2.71828;
        double c = t.a / t.b;
        System.out.println("OK!");
    }
    
}
