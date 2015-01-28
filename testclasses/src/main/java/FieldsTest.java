public class FieldsTest {
    
    static int x;
    static int y;
    int a;
    int b;
    
    public static void main(String[] args) {
        FieldsTest.x = 100;
        FieldsTest.y = 200;
        long z = FieldsTest.x / FieldsTest.y;
        FieldsTest t = new FieldsTest();
        t.a = 100;
        t.b = 200;
        int c = t.a / t.b;
        System.out.println("OK!");
    }
    
}
