package stdlib.basic.cl;

import static helper.MyAssert.*;

public class ClassLoaderTest implements Runnable {
    
    public static void main(String[] args) {
       new ClassLoaderTest().run();
    }

    @Override
    public void run() {
        // primitive types
        assertNull(int.class.getClassLoader());
        // array types
        assertNull(float[].class.getClassLoader());
        assertNull(new int[0].getClass().getClassLoader());
        assertNull(new Object[0].getClass().getClassLoader());
        assertNull(int[][].class.getClassLoader());
        // bootstrap loader
        assertNull(Object.class.getClassLoader());
        assertNull("".getClass().getClassLoader());
        // system loader
        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
        System.out.println(sysCl.toString());
        System.out.println(ClassLoaderTest.class.getClassLoader());
//        assertSame(sysCl, ClassLoaderTest.class.getClassLoader());
        // platform loader
        // ClassLoader.getPlatformClassLoader();
    }

}
