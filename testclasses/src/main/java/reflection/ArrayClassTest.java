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
        Class<?> c = new int[0].getClass();
        assertEquals("[I", c.getName());
        assertEquals(Object.class, c.getSuperclass());
        assertArrayEquals(new Class<?>[]{Cloneable.class, Serializable.class}, c.getInterfaces());
        assertEquals(0, c.getFields().length);
        assertEquals(0, c.getDeclaredFields().length);
        assertEquals(9, c.getMethods().length);
        assertEquals(0, c.getDeclaredMethods().length);
    }
    
}
