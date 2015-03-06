package reflection;

import junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ClassLoaderTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ClassLoaderTest.class);
    }
    
    @Test
    public void sysClassLoader() {
        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
        assertEquals("sun.misc.Launcher$AppClassLoader", sysCl.getClass().getName());
        
        ClassLoader extCl = sysCl.getParent();
        assertEquals("sun.misc.Launcher$ExtClassLoader", extCl.getClass().getName());
        
        ClassLoader bootCl = extCl.getParent();
        assertNull(bootCl);
    }
    
    @Test
    public void getClassLoader() {
        ClassLoader bootCl = Object.class.getClassLoader();
        assertNull(bootCl);
        
        ClassLoader appCl = ClassLoaderTest.class.getClassLoader();
        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
        assertSame(sysCl, appCl);
    }
    
}
