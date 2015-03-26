package java7.field;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

// A constant variable is a final variable of primitive type or type String
// that is initialized with a constant expression (§15.28).
public class ConstantVariablesTest {
    
    public static final byte b = 125;
    public static final char c = 'c';
    public static final short s = 300;
    public static final int x = 100;
    public static final int y = x + 13;
    public static final long j = 1L;
    public static final float f = 3.14f;
    public static final double d = 2.71828;
    public static final String str1 = "hello";
    public static final String str2 = str1 + "world";
    
    @Test
    public void test() {
        assertEquals((byte)125, ConstantVariablesTest.b);
        assertEquals((byte)125, getFieldValue("b"));
        
        
    }
    
    private static Object getFieldValue(String name) {
        try {
            return ConstantVariablesTest.class.getField(name).get(null);
        } catch (ReflectiveOperationException e) {
            throw new RuntimeException(e);
        }
    }
    
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(ConstantVariablesTest.class);
    }
    
}
