package stdlib.basic.cl;

import java.io.InputStream;
import java.lang.reflect.Method;
import java.net.URL;
import java.net.URLClassLoader;

import static helper.MyAssert.*;

public class ClassLoaderTest2 implements Runnable {
    
    public static void main(String[] args) {
        new ClassLoaderTest2().run();
    }

    @Override
    public void run() {
        // TODO
    }

    public void getResource() {
        ClassLoader appCl = ClassLoaderTest2.class.getClassLoader();
        //URL url = appCl.getResource("org/eclipse/jetty/http/mime.properties");
        URL url = appCl.getResource("LICENSE.txt");
        //System.out.println(url);
        assertNotNull(url);
        
        InputStream is = appCl.getResourceAsStream("LICENSE.txt");
        assertNotNull(is);
    }

    public void findLoadedClass() throws Exception {
        Method m = ClassLoader.class.getDeclaredMethod("findLoadedClass", String.class);
        m.setAccessible(true);
        
        ClassLoader appCl = ClassLoaderTest2.class.getClassLoader();
        assertEquals(Object.class, m.invoke(appCl, "java.lang.Object"));
        assertEquals(ClassLoaderTest2.class, m.invoke(appCl, "stdlib.basic.cl.ClassLoaderTest"));
        assertEquals(null, m.invoke(appCl, "foo.bar.XXX"));
        
        ClassLoader urlCl = new URLClassLoader(new URL[0]);
        assertEquals(appCl, urlCl.getParent());
        assertEquals(null, m.invoke(urlCl, "java.lang.Object"));
        assertEquals(null, m.invoke(urlCl, "jvm.java7.cl.ClassLoaderTest"));
    }

}
