package java6.field;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class FieldsTest {
    
    static class Sup {
        static int x;
        int a;
    }
    
    static class Sub extends Sup {
        static int y;
        int b;
    }
    
    @Test
    public void staticFields() {
        int z = Sub.x + Sub.y;
        z += 100;
        Sub.y = z;
        Sub.x = z;
        assertTrue(Sub.x == 100 && Sub.y == 100);
    }
    
    @Test
    public void instanceFields() {
        Sub sub = new Sub();
        int c = sub.a + sub.b;
        c += 100;
        sub.a = c;
        sub.b = c;
        assertTrue(sub.a == 100 && sub.b == 100);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(FieldsTest.class);
    }
    
}
