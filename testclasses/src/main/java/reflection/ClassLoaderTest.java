package reflection;

import junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ClassLoaderTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ClassLoaderTest.class);
        

//        System.out.println(appCl.getParent());
//        System.out.println(appCl.getParent().getParent());
        
//        ClassLoader bootCl = Object.class.getClassLoader();
//        System.out.println("bootCl:"+bootCl);
//        ClassLoader appCl = ClassLoaderTest.class.getClassLoader();
//        System.out.println("appCl:"+appCl);
//        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
//        System.out.println("sysCl:"+sysCl);
//        
//        System.out.println(Test.class.getClassLoader());
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
    
    //@Test
    public void test() {
        ClassLoader bootCl = Object.class.getClassLoader();
        ClassLoader appCl = ClassLoaderTest.class.getClassLoader();
        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
        ClassLoader extCl = appCl.getParent();
        
        assertNull(bootCl);
        assertNotNull(appCl);
        assertNotNull(extCl);
        assertSame(appCl, sysCl);
    }
    
}
