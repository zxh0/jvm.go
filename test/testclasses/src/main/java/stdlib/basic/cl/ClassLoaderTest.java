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
        testParents();
        testLoadClass();
        //testClassNotFound(); // TODO
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
        ClassLoader plCL = ClassLoader.getPlatformClassLoader();
        assertTrue(plCL.toString().startsWith(
                "jdk.internal.loader.ClassLoaders$PlatformClassLoader@"));
    }

    private static void testAppLoader() {
        ClassLoader sysCL = ClassLoader.getSystemClassLoader();
        assertSame(sysCL, ClassLoaderTest.class.getClassLoader());
        assertTrue(sysCL.toString().startsWith(
                "jdk.internal.loader.ClassLoaders$AppClassLoader@"));
    }

    private static void testParents() {
        ClassLoader plCL = ClassLoader.getPlatformClassLoader();
        ClassLoader sysCL = ClassLoader.getSystemClassLoader();
        assertSame(plCL, sysCL.getParent());
        assertNull(plCL.getParent());
    }

    private static void testLoadClass() {
        try {
            ClassLoader sysCL = ClassLoader.getSystemClassLoader();
            assertSame(Object.class, sysCL.loadClass("java.lang.Object"));
            assertSame(ClassLoaderTest.class, sysCL.loadClass("stdlib.basic.cl.ClassLoaderTest"));
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    public void testClassNotFound() {
        try {
            ClassLoader sysCl = ClassLoader.getSystemClassLoader();
            sysCl.loadClass("foo.bar.XXX");
            fail();
        } catch (ClassNotFoundException e) {
            assertEquals("foo.bar.XXX", e.getMessage());
        }
    }

}
