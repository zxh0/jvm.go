package java6.cl;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class GetClassLoaderTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(GetClassLoaderTest.class);
    }
    
    @Test
    public void primitiveTypes() {
        assertNull(int.class.getClassLoader());
    }
    
    @Test
    public void array() {
        assertNull(float[].class.getClassLoader());
        assertNull(new int[0].getClass().getClassLoader());
        assertNull(new Object[0].getClass().getClassLoader());
        assertNull(int[][].class.getClassLoader());
    }
    
    @Test
    public void bootCl() {
        assertNull(Object.class.getClassLoader());
        assertNull("".getClass().getClassLoader());
    }
    
    //@Test
    public void sysCl() {
        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
        assertSame(sysCl, GetClassLoaderTest.class.getClassLoader());
    }
    
}
