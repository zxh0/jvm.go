package java6.field;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class FieldAccessTest {
    
    private static interface I {
        static int i = val(1);
    }
    private static interface J {
        static int j = val(2);
    }
    private static interface K extends I, J {
        static int k = val(3);
    }
    private static class A implements K {
        static int a = val(4);
    }
    private static class B extends A {
        static int b = val(5);
    }
    
    private static int val(int x) {
        return x;
    }
    
    @Test
    public void test() {
        assertEquals(1, B.i);
        assertEquals(2, B.j);
        assertEquals(3, B.k);
        assertEquals(4, B.a);
        assertEquals(5, B.b);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(FieldAccessTest.class);
    }
    
}
