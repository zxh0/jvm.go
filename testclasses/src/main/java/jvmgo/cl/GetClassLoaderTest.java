package jvmgo.cl;

import jvmgo.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class GetClassLoaderTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(GetClassLoaderTest.class);
    }
    
    @Test
    public void array() {
        assertNull(boolean[].class.getClassLoader());
        assertNull(int[].class.getClassLoader());
    }
    
    @Test
    public void getClassLoader() {
        ClassLoader bootCl = Object.class.getClassLoader();
        assertNull(bootCl);
        
        ClassLoader appCl = GetClassLoaderTest.class.getClassLoader();
        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
        assertSame(sysCl, appCl);
    }
    
}
