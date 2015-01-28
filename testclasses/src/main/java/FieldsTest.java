public class FieldsTest {
    
    static int x;
    static int y;
    int a;
    int b;
    
    public static void main(String[] args) {
        FieldsTest.x = 200;
        FieldsTest.y = 100;
        int z = FieldsTest.x / FieldsTest.y;
        if (z == 2) {
            System.out.println("OKXY!");
        } else {
            System.out.println("FailXY!");
        }
        
        FieldsTest t = new FieldsTest();
        t.a = 200;
        t.b = 100;
        int c = t.a / t.b;
        if (c == 2) {
            System.out.println("OKAB!");
        } else {
            System.out.println("FailAB!");
        }
    }
    
}
