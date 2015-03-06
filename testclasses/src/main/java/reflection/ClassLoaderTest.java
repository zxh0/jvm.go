package reflection;

import junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class ClassLoaderTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(ClassLoaderTest.class);
    }
    
    @Test
    public void test() {
        ClassLoader bootCl = Object.class.getClassLoader();
        ClassLoader appCl = ClassLoaderTest.class.getClassLoader();
        
        assertNull(bootCl);
        assertNotNull(appCl);
    }
    
}
