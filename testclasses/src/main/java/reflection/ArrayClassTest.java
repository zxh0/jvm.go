package reflection;

import java.io.Serializable;
import org.junit.Test;
import unit.UnitTestRunner;
import static org.junit.Assert.*;

public class ArrayClassTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ArrayClassTest.class);
    }
    
    @Test
    public void primitiveArray() {
        testArrayClass(int[].class, "[I");
    }
    
    private void testArrayClass(Class<?> c, String name) {
        assertEquals(Object.class, c.getSuperclass());
        assertEquals(name, c.getName());
        assertArrayEquals(new Class<?>[]{Cloneable.class, Serializable.class}, c.getInterfaces());
        assertEquals(0, c.getFields().length);
        assertEquals(0, c.getDeclaredFields().length);
        assertEquals(9, c.getMethods().length);
        assertEquals(0, c.getDeclaredMethods().length);
    }
    
}
