package stdlib.basic.cl;

import static helper.MyAssert.*;

public class ClassLoaderTest implements Runnable {

    public static void main(String[] args) {
       new ClassLoaderTest().run();
    }

    @Override
    public void run() {
        testBootLoader();
        testPlatformLoader();
        testAppLoader();
    }

    private static void testBootLoader() {
        // primitive types
        assertNull(int.class.getClassLoader());
        // array types
        assertNull(float[].class.getClassLoader());
        assertNull(new int[0].getClass().getClassLoader());
        assertNull(new Object[0].getClass().getClassLoader());
        assertNull(int[][].class.getClassLoader());
        // basic types
        assertNull(Object.class.getClassLoader());
        assertNull("".getClass().getClassLoader());
    }

    private static void testPlatformLoader() {
         ClassLoader pl = ClassLoader.getPlatformClassLoader();
    }

    private static void testAppLoader() {
        ClassLoader sysCl = ClassLoader.getSystemClassLoader();
        assertSame(sysCl, ClassLoaderTest.class.getClassLoader());
        assertTrue(sysCl.toString().startsWith(
                "jdk.internal.loader.ClassLoaders$AppClassLoader@"));
    }

}
