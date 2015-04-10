package java6;

import libs.junit.UnitTestRunner;
import org.junit.Test;

public class SysPropsTest {
    
    public static void main(String[] args) {
        //UnitTestRunner.run(SysPropsTest.class);
        System.out.println(System.getProperty("os.name"));
        System.out.println(System.getProperty("os.arch"));
        System.out.println(System.getProperty("os.version"));
    }
    
    @Test
    public void test() {
        String osName = System.getProperty("os.name");
        System.out.println(osName);
    }
    
}
